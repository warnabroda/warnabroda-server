package routes

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"
	
	"bitbucket.org/hbtsmith/warnabrodagomartini/models"	
	"bitbucket.org/hbtsmith/warnabrodagomartini/messages"
	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"
)

const (
	SQL_SELECT_SUBJECTS_BY_ID 		= "SELECT * FROM subjects ORDER BY Id"	
	SQL_SELECT_SUBJECTS_BY_LANG_KEY	= "SELECT * FROM subjects WHERE lang_key = ? ORDER BY Id"	
)

func GetSubjectsByLangKey(lang_key string) []models.DefaultStruct{

	var subjects []models.DefaultStruct
	_, err := models.Dbm.Select(&subjects, SQL_SELECT_SUBJECTS_BY_LANG_KEY, lang_key)
	if err != nil {
		checkErr(err, "SELECT ERROR")
	} 

	return subjects
}

// Get a random subject from the previews loaded upon containers startup
func GetRandomSubject(lang_key string) models.DefaultStruct {		
	
	 var subject models.DefaultStruct

	subjects := GetSubjectsByLangKey(lang_key)

	// r := rand.New(rand.NewSource(99))
	rand.Seed(time.Now().UTC().UnixNano())	
	
	 if len(subjects) > 0 {
				
	 	var index = rand.Intn(len(subjects))
		
	 	if index < len(subjects) {
	 		subject = subjects[index]
	 	} else {
	 		subject = subjects[0]
	 	}

	 } else {
		
		subject = models.DefaultStruct{0, messages.GetLocaleMessage(lang_key, "MSG_DEFAULT_SUBJECT"), lang_key, true, ""}
	 }

	return subject
}

func GetSubjects(enc Encoder, db gorp.SqlExecutor) (int, string) {
	var subjects []models.DefaultStruct
	_, err := db.Select(&subjects, SQL_SELECT_SUBJECTS_BY_ID)
	if err != nil {
		checkErr(err, "SELECT SUBJECTS FAILED")
		return http.StatusInternalServerError, ""
	}
	return http.StatusOK, Must(enc.Encode(subjectsToIface(subjects)...))
}

// Get a subject by ID
func GetSubject(enc Encoder, db gorp.SqlExecutor, parms martini.Params) (int, string) {
	id, err := strconv.Atoi(parms["id"])
	obj, _ := db.Get(models.DefaultStruct{}, id)
	if err != nil || obj == nil {
		checkErr(err, "GET SUBJECT FAILED")
		// Invalid id, or does not exist
		return http.StatusNotFound, ""
	}
	entity := obj.(*models.DefaultStruct)
	return http.StatusOK, Must(enc.EncodeOne(entity))
}

func subjectsToIface(v []models.DefaultStruct) []interface{} {
	if len(v) == 0 {
		return nil
	}
	ifs := make([]interface{}, len(v))
	for i, v := range v {
		ifs[i] = v
	}
	return ifs
}
