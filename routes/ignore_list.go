package routes

import (
	"bitbucket.org/hbtsmith/warnabrodagomartini/models"
	"fmt"
	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"
	"net/http"
	"strconv"
)

func AddIgnoreList(entity models.Ignore_List, w http.ResponseWriter, enc Encoder, db gorp.SqlExecutor) (int, string) {

	
	status := &models.Message{
		Id:       200,
		Name:     "Contato Adicionado à Lista de Ignorados!",
		Lang_key: "br",
	}

	entity.Created_by = "user"
	entity.Created_date = time.Now().String()	

	err := db.Insert(&entity)
	if err != nil {
		//checkErr(err, "insert failed")
		status := &models.Message{
			Id:       403,
			Name:     "Contato Já estava na Lista de Ignorados!",
			Lang_key: "br",
		}
	}
	w.Header().Set("Location", fmt.Sprintf("/warnabroda/ignore_list/%d", entity.Id))
	return http.StatusCreated, Must(enc.EncodeOne(status))
}

func DeleteIgnoreList(db gorp.SqlExecutor, parms martini.Params) (int, string) {
	id, err := strconv.Atoi(parms["id"])
	obj, _ := db.Get(models.Ignore_List{}, id)
	if err != nil || obj == nil {
		checkErr(err, "get failed")
		// Invalid id, or does not exist
		return http.StatusNotFound, ""
	}
	entity := obj.(*models.Ignore_List)
	_, err = db.Delete(entity)
	if err != nil {
		checkErr(err, "delete failed")
		return http.StatusConflict, ""
	}
	return http.StatusNoContent, ""
}
