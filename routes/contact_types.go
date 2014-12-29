package routes

import (
	"bitbucket.org/hbtsmith/warnabrodagomartini/models"
	"fmt"
	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"
	"net/http"
	"strconv"
)

func GetContact_types(enc Encoder, db gorp.SqlExecutor) (int, string) {
	var contact_types []models.DefaultStruct
	_, err := db.Select(&contact_types, "select * from contact_types order by id")
	if err != nil {
		checkErr(err, "select failed")
		return http.StatusInternalServerError, ""
	}
	return http.StatusOK, Must(enc.Encode(contact_typesToIface(contact_types)...))
}

func GetContact_type(enc Encoder, db gorp.SqlExecutor, parms martini.Params) (int, string) {
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

func AddContact_type(entity models.DefaultStruct, w http.ResponseWriter, enc Encoder, db gorp.SqlExecutor) (int, string) {
	err := db.Insert(&entity)
	if err != nil {
		checkErr(err, "insert failed")
		return http.StatusConflict, ""
	}
	w.Header().Set("Location", fmt.Sprintf("/warnabroda/contact_types/%d", entity.Id))
	return http.StatusCreated, Must(enc.EncodeOne(entity))
}

func UpdateContact_type(entity models.DefaultStruct, enc Encoder, db gorp.SqlExecutor, parms martini.Params) (int, string) {
	id, err := strconv.Atoi(parms["id"])
	obj, _ := db.Get(models.DefaultStruct{}, id)
	if err != nil || obj == nil {
		checkErr(err, "get failed")
		// Invalid id, or does not exist
		return http.StatusNotFound, ""
	}
	oldEntity := obj.(*models.DefaultStruct)

	entity.Id = oldEntity.Id
	_, err = db.Update(&entity)
	if err != nil {
		checkErr(err, "update failed")
		return http.StatusConflict, ""
	}
	return http.StatusOK, Must(enc.EncodeOne(entity))
}

func DeleteContact_type(db gorp.SqlExecutor, parms martini.Params) (int, string) {
	id, err := strconv.Atoi(parms["id"])
	obj, _ := db.Get(models.DefaultStruct{}, id)
	if err != nil || obj == nil {
		checkErr(err, "get failed")
		// Invalid id, or does not exist
		return http.StatusNotFound, ""
	}
	entity := obj.(*models.DefaultStruct)
	_, err = db.Delete(entity)
	if err != nil {
		checkErr(err, "delete failed")
		return http.StatusConflict, ""
	}
	return http.StatusNoContent, ""
}

func contact_typesToIface(v []models.DefaultStruct) []interface{} {
	if len(v) == 0 {
		return nil
	}
	ifs := make([]interface{}, len(v))
	for i, v := range v {
		ifs[i] = v
	}
	return ifs
}
