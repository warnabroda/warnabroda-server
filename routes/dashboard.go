package routes

import (
	"net/http"
	"strconv"

	//"fmt"
	"github.com/coopernurse/gorp"
	"github.com/martini-contrib/sessionauth"
	"gitlab.com/warnabroda/warnabrodagomartini/models"
)

const (
	SQL_WARNING_COUNT = "SELECT COUNT(*) AS total FROM warnings WHERE Sent= :sent"
	SQL_WARN_COUNT    = "SELECT " +
		"(SELECT COUNT(*) FROM warnabroda.warnings) AS Allw, " +
		"(SELECT COUNT(*) FROM warnabroda.warnings WHERE sent = true) AS Sent, " +
		"(SELECT COUNT(*) FROM warnabroda.warnings WHERE sent = true AND Id_contact_type = 1) AS EmailSent, " +
		"(SELECT COUNT(*) FROM warnabroda.warnings WHERE sent = true AND Id_contact_type = 2) AS SmsSent,  " +
		"(SELECT COUNT(*) FROM warnabroda.warnings WHERE sent = true AND Id_contact_type = 3) AS WhatsappSent,  " +
		"(SELECT COUNT(*) FROM warnabroda.warnings WHERE sent = false) AS NotSent, " +
		"(SELECT COUNT(*) FROM warnabroda.warnings WHERE sent = false AND Id_contact_type = 1) AS EmailNotSent, " +
		"(SELECT COUNT(*) FROM warnabroda.warnings WHERE sent = false AND Id_contact_type = 2) AS SmsNotSent, " +
		"(SELECT COUNT(*) FROM warnabroda.warnings WHERE sent = false AND Id_contact_type = 3) AS WhatsappNotSent, " +
		"(SELECT COUNT(*) FROM warnabroda.ignore_list) AS IgnoreList, " +
		"(SELECT COUNT(*) FROM warnabroda.ignore_list WHERE confirmed = true) AS Confirmed, " +
		"(SELECT COUNT(*) FROM warnabroda.ignore_list WHERE confirmed = false) AS Unconfirmed, " +
		"(SELECT COUNT(*) FROM warnabroda.ignore_list WHERE confirmed = true  AND contact LIKE '%@%') AS ConfirmedByEmail, " +
		"(SELECT COUNT(*) FROM warnabroda.ignore_list WHERE confirmed = false AND contact LIKE '%@%') AS UnconfirmedByEmail, " +
		"(SELECT COUNT(*) FROM warnabroda.ignore_list WHERE confirmed = true  AND contact NOT LIKE '%@%') AS ConfirmedBySms, " +
		"(SELECT COUNT(*) FROM warnabroda.ignore_list WHERE confirmed = false AND contact NOT LIKE '%@%') AS UnconfirmedBySms, " +
		"(SELECT COUNT(DISTINCT(contact)) FROM warnabroda.warnings) AS UniqueContacts "
)

func ListWarnings(entity models.Warn, enc Encoder, user sessionauth.User, db gorp.SqlExecutor) (int, string) {

	u := UserById(user.UniqueId().(int), db)

	if user.IsAuthenticated() && u.UserRole == models.ROLE_ADMIN {
		sql := "SELECT w.id, msg.name AS message, ct.name AS contact_type, w.contact, w.sent, w.created_date FROM warnings AS w "
		sql += "INNER JOIN messages AS msg ON (msg.id = w.id_message) "
		sql += "INNER JOIN contact_types AS ct ON (ct.id = w.id_contact_type) "
		sql += "ORDER BY w.created_date DESC "

		var warns []models.Warn
		_, err := db.Select(&warns, sql)
		checkErr(err, "SELECT ALL WARNINGS ERROR")

		if err != nil {
			return http.StatusBadRequest, ""
		}
		return http.StatusOK, Must(enc.Encode(warnsToIface(warns)...))

	}

	return http.StatusUnauthorized, ""

}

func warnsToIface(v []models.Warn) []interface{} {
	if len(v) == 0 {
		return nil
	}
	ifs := make([]interface{}, len(v))
	for i, v := range v {
		ifs[i] = v
	}
	return ifs
}

// Count sent warnings
func CountSentWarnings(enc Encoder, db gorp.SqlExecutor) (int, string) {

	total, err := countWarnings(true, db)

	if err == nil {
		return http.StatusOK, total
	} else {
		return http.StatusBadRequest, ""
	}

	return http.StatusBadRequest, ""
}

// count all warnings registered
func WarnaCounter(enc Encoder, db gorp.SqlExecutor, user sessionauth.User) (int, string) {

	counts := models.CountWarning{}

	u := UserById(user.UniqueId().(int), db)

	if user.IsAuthenticated() && u.UserRole == models.ROLE_ADMIN {

		err := db.SelectOne(&counts, SQL_WARN_COUNT)
		checkErr(err, "COUNT SENT WARNINGS ERROR")

		if err == nil {
			return http.StatusOK, Must(enc.EncodeOne(counts))
		} else {
			return http.StatusBadRequest, ""
		}

	}

	return http.StatusUnauthorized, ""
}

// count warnings according to the param sent(true or false) and the specific type of contact
func countWarnings(sent bool, db gorp.SqlExecutor) (string, error) {

	total, err := db.SelectInt(SQL_WARNING_COUNT, map[string]interface{}{
		"sent": sent,
	})
	checkErr(err, "COUNT SENT WARNINGS ERROR")

	if err != nil {
		return "", err
	}

	return strconv.FormatInt(total, 10), nil
}
