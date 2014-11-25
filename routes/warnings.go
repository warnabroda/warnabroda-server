package routes

import (
    "warnabroda/models"
    "fmt"
    "net/http"
    "strconv"
    "github.com/go-martini/martini"
    "github.com/coopernurse/gorp"
)

func GetWarnings(enc Encoder, db gorp.SqlExecutor) (int, string) {
    var warnings []models.Warning
    _, err := db.Select(&warnings, "select * from warnings order by id")
    if err != nil {
        checkErr(err, "select failed")
        return http.StatusInternalServerError, ""
    }
    return http.StatusOK, Must(enc.Encode(warningsToIface(warnings)...))
}

func GetWarning(enc Encoder, db gorp.SqlExecutor, parms martini.Params) (int, string) {
    id, err := strconv.Atoi(parms["id"])
    obj, _ := db.Get(models.Warning{}, id)
    if err != nil || obj == nil {
        checkErr(err, "get failed")
        // Invalid id, or does not exist
        return http.StatusNotFound, ""
    }
    entity := obj.(*models.Warning)
    return http.StatusOK, Must(enc.EncodeOne(entity))
}

func AddWarning(entity models.Warning, w http.ResponseWriter, enc Encoder, db gorp.SqlExecutor) (int, string) {
    err := db.Insert(&entity)
    if err != nil {
        checkErr(err, "insert failed")
        return http.StatusConflict, ""
    }
    w.Header().Set("Location", fmt.Sprintf("/warnabroda/warnings/%d", entity.Id))
    return http.StatusCreated, Must(enc.EncodeOne(entity))
}

func UpdateWarning(entity models.Warning, enc Encoder, db gorp.SqlExecutor, parms martini.Params) (int, string) {
    id, err := strconv.Atoi(parms["id"])
    obj, _ := db.Get(models.Warning{}, id)
    if err != nil || obj == nil {
        checkErr(err, "get failed")
        // Invalid id, or does not exist
        return http.StatusNotFound, ""
    }
    oldEntity := obj.(*models.Warning)

    entity.Id = oldEntity.Id
    _, err = db.Update(&entity)
    if err != nil {
        checkErr(err, "update failed")
        return http.StatusConflict, ""
    }
    return http.StatusOK, Must(enc.EncodeOne(entity))
}

func DeleteWarning(db gorp.SqlExecutor, parms martini.Params) (int, string) {
    id, err := strconv.Atoi(parms["id"])
    obj, _ := db.Get(models.Warning{}, id)
    if err != nil || obj == nil {
        checkErr(err, "get failed")
        // Invalid id, or does not exist
        return http.StatusNotFound, ""
    }
    entity := obj.(*models.Warning)
    _, err = db.Delete(entity)
    if err != nil {
        checkErr(err, "delete failed")
        return http.StatusConflict, ""
    }
    return http.StatusNoContent, ""
}

func warningsToIface(v []models.Warning) []interface{} {
    if len(v) == 0 {
        return nil
    }
    ifs := make([]interface{}, len(v))
    for i, v := range v {
        ifs[i] = v
    }
    return ifs
}
