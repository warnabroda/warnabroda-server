package routes

import (
    "warnabroda/models"
    "fmt"
    "net/http"
    "strconv"
    "github.com/go-martini/martini"
    "github.com/coopernurse/gorp"
)

func GetMessages(enc Encoder, db gorp.SqlExecutor) (int, string) {
    var messages []models.Message
    _, err := db.Select(&messages, "select * from messages order by id")
    if err != nil {
        checkErr(err, "select failed")
        return http.StatusInternalServerError, ""
    }
    return http.StatusOK, Must(enc.Encode(messagesToIface(messages)...))
}

func GetMessage(enc Encoder, db gorp.SqlExecutor, parms martini.Params) (int, string) {
    id, err := strconv.Atoi(parms["id"])
    obj, _ := db.Get(models.Message{}, id)
    if err != nil || obj == nil {
        checkErr(err, "get failed")
        // Invalid id, or does not exist
        return http.StatusNotFound, ""
    }
    entity := obj.(*models.Message)
    return http.StatusOK, Must(enc.EncodeOne(entity))
}

func SelectMessage(db gorp.SqlExecutor, id int64) models.Message {
    
    entity := models.Message{}

    err := db.SelectOne(&entity,"SELECT * FROM messages WHERE id=?", id)   
    
    if err != nil {
        checkErr(err, "select failed")
    }

    return entity   
}

func AddMessage(entity models.Message, w http.ResponseWriter, enc Encoder, db gorp.SqlExecutor) (int, string) {
    err := db.Insert(&entity)
    if err != nil {
        checkErr(err, "insert failed")
        return http.StatusConflict, ""
    }
    w.Header().Set("Location", fmt.Sprintf("/warnabroda/messages/%d", entity.Id))
    return http.StatusCreated, Must(enc.EncodeOne(entity))
}

func UpdateMessage(entity models.Message, enc Encoder, db gorp.SqlExecutor, parms martini.Params) (int, string) {
    id, err := strconv.Atoi(parms["id"])
    obj, _ := db.Get(models.Message{}, id)
    if err != nil || obj == nil {
        checkErr(err, "get failed")
        // Invalid id, or does not exist
        return http.StatusNotFound, ""
    }
    oldEntity := obj.(*models.Message)

    entity.Id = oldEntity.Id
    _, err = db.Update(&entity)
    if err != nil {
        checkErr(err, "update failed")
        return http.StatusConflict, ""
    }
    return http.StatusOK, Must(enc.EncodeOne(entity))
}

func DeleteMessage(db gorp.SqlExecutor, parms martini.Params) (int, string) {
    id, err := strconv.Atoi(parms["id"])
    obj, _ := db.Get(models.Message{}, id)
    if err != nil || obj == nil {
        checkErr(err, "get failed")
        // Invalid id, or does not exist
        return http.StatusNotFound, ""
    }
    entity := obj.(*models.Message)
    _, err = db.Delete(entity)
    if err != nil {
        checkErr(err, "delete failed")
        return http.StatusConflict, ""
    }
    return http.StatusNoContent, ""
}

func messagesToIface(v []models.Message) []interface{} {
    if len(v) == 0 {
        return nil
    }
    ifs := make([]interface{}, len(v))
    for i, v := range v {
        ifs[i] = v
    }
    return ifs
}
