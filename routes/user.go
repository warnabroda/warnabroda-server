package routes

import (
	//"fmt"
	"net/http"
	"strconv"

	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessionauth"
	"github.com/martini-contrib/sessions"
	"gitlab.com/warnabroda/warnabrodagomartini/messages"
	"gitlab.com/warnabroda/warnabrodagomartini/models"
	// "encoding/json"
	// "strings"
)

const (
	SQL_LOGIN = "SELECT * FROM users WHERE (username = :user OR email = :user) AND password = :pass "
)

func GetUserById(enc Encoder, db gorp.SqlExecutor, user sessionauth.User, parms martini.Params) (int, string) {

	u := models.GetAuthenticatedUser(user)

	if user.IsAuthenticated() && u.UserRole == models.ROLE_ADMIN {
		id, err := strconv.Atoi(parms["id"])

		if err != nil {
			return http.StatusNotFound, ""
		}

		entity := UserById(id, db)

		return http.StatusOK, Must(enc.EncodeOne(entity))
	}

	return http.StatusUnauthorized, ""
}

func UserById(id int, db gorp.SqlExecutor) *models.User {

	obj, err := db.Get(models.User{}, id)
	if err != nil {
		return nil
	}
	entity := obj.(*models.User)

	return entity
}

func GetUserByLogin(postedUser models.UserLogin, db gorp.SqlExecutor) *models.User {

	user := models.User{}

	db.SelectOne(&user, SQL_LOGIN,
		map[string]interface{}{
			"user": postedUser.Username,
			"pass": postedUser.Password,
		})

	return &user
}

func DoLogin(entity models.UserLogin, session sessions.Session, enc Encoder, db gorp.SqlExecutor) (int, string) {

	status := &models.DefaultStruct{
		Id:       http.StatusForbidden,
		Name:     messages.GetLocaleMessage("en", "MSG_LOGIN_INVALID"),
		Lang_key: "en",
	}

	user := GetUserByLogin(entity, db)

	if user.Name != "" {

		err := sessionauth.AuthenticateSession(session, user)
		if err != nil {
			status.Name = messages.GetLocaleMessage("en", "MSG_SESSION_INIT_ERROR")
			return http.StatusForbidden, Must(enc.EncodeOne(status))
		}
		user.Authenticated = true
		user.UpdateLastLogin()
		status.Name = messages.GetLocaleMessage("en", "MSG_SUCCESSFUL_LOGIN")
		return http.StatusOK, Must(enc.EncodeOne(user))

	} else {

		sessionauth.Logout(session, user)
		session.Clear()
		return http.StatusForbidden, Must(enc.EncodeOne(status))

	}

	return http.StatusForbidden, Must(enc.EncodeOne(status))
}

func IsAuthenticated(enc Encoder, user sessionauth.User) (int, string) {

	u := models.GetAuthenticatedUser(user)

	if user.IsAuthenticated() && u.UserRole == models.ROLE_ADMIN {
		return http.StatusOK, ""
	}

	return http.StatusUnauthorized, Must(enc.EncodeOne(user))
}

func DoLogout(enc Encoder, session sessions.Session, user sessionauth.User, db gorp.SqlExecutor) (int, string) {

	status := &models.DefaultStruct{
		Id:       http.StatusOK,
		Name:     messages.GetLocaleMessage("en", "MSG_LOGIN_REQUIRED"),
		Lang_key: "en",
	}

	if user.IsAuthenticated() {

		sessionauth.Logout(session, user)
		session.Clear()
		status.Name = messages.GetLocaleMessage("en", "MSG_SUCCESSFUL_LOGOUT")
	}

	updateUser := UserById(user.UniqueId().(int), db)

	updateUser.Authenticated = false
	db.Update(updateUser)

	return http.StatusOK, Must(enc.EncodeOne(status))
}

func GetAuthenticatedUser(enc Encoder, user sessionauth.User, db gorp.SqlExecutor) (int, string) {

	if user.IsAuthenticated() {

		authUser := UserById(user.UniqueId().(int), db)

		return http.StatusOK, Must(enc.EncodeOne(authUser))
	}

	return http.StatusUnauthorized, Must(enc.EncodeOne(user))
}
