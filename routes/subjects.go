package routes

import (
	"bitbucket.org/hbtsmith/warnabrodagomartini/models"	
	"bitbucket.org/hbtsmith/warnabrodagomartini/i18n"
	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var global_subjects []models.DefaultStruct

const (
	SQL_SELECT_SUBJECTS_BY_ID		= "SELECT * FROM subjects ORDER BY Id"
	MSG_DEFAULT_SUBJECT				= "Um amigo(a) acaba de lhe dar um toque"
)

func init(){
	_, err := models.Dbm.Select(&global_subjects, SQL_SELECT_SUBJECTS_BY_ID)
	if err != nil {
		checkErr(err, "SELECT ERROR")
	} 
}

// Get a random subject from the previews loaded upon containers startup
func GetRandomSubject() models.DefaultStruct {		
	
	var subject models.DefaultStruct

	// r := rand.New(rand.NewSource(99))
	rand.Seed(time.Now().UTC().UnixNano())	
	
	 if len(global_subjects) > 0 {
				
	 	var index = rand.Intn(len(global_subjects))
		
	 	if index < len(global_subjects) {
	 		subject = global_subjects[index]
	 	} else {
	 		subject = global_subjects[0]
	 	}

	 } else {
		
		subject = models.DefaultStruct{0, MSG_DEFAULT_SUBJECT, i18n.BR_LANG_KEY}
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
