package routes

import (
	"bitbucket.org/hbtsmith/warnabrodagomartini/models"	
	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var global_subjects []models.DefaultStruct

func init(){
	_, err := models.Dbm.Select(&global_subjects, "SELECT * FROM subjects ORDER BY Id")
	if err != nil {
		checkErr(err, "SELECT ERROR")
	} 
}


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
		
		subject = models.DefaultStruct{0, "Um amigo(a) acaba de lhe dar um toque", "br"}
	 }

	return subject
}

func GetSubjects(enc Encoder, db gorp.SqlExecutor) (int, string) {
	var subjects []models.DefaultStruct
	_, err := db.Select(&subjects, "select * from subjects order by id")
	if err != nil {
		checkErr(err, "select failed")
		return http.StatusInternalServerError, ""
	}
	return http.StatusOK, Must(enc.Encode(subjectsToIface(subjects)...))
}

func GetSubject(enc Encoder, db gorp.SqlExecutor, parms martini.Params) (int, string) {
	id, err := strconv.Atoi(parms["id"])
	obj, _ := db.Get(models.DefaultStruct{}, id)
	if err != nil || obj == nil {
		checkErr(err, "get failed")
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
