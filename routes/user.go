package routes

import (
	"bitbucket.org/hbtsmith/warnabrodagomartini/models"	
	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"	
	"github.com/martini-contrib/sessionauth"
	"github.com/martini-contrib/sessions"
	"net/http"
	"strconv"
	"fmt"

)

func GetUserById(enc Encoder, db gorp.SqlExecutor, parms martini.Params) (int, string) {
	
	id, err := strconv.Atoi(parms["id"])
	obj, _ := db.Get(models.User{}, id)
	if err != nil || obj == nil {
		checkErr(err, "get failed")
		// Invalid id, or does not exist
		return http.StatusNotFound, ""
	}
	entity := obj.(*models.User)
	fmt.Println("ta sendo chamado")
	fmt.Println(entity)
	return http.StatusOK, Must(enc.EncodeOne(entity))
}

func GetUserByLogin(postedUser models.User) *models.User {	

	user := models.User{}

	err := models.Dbm.SelectOne(&user, "SELECT * FROM users WHERE username = :user and password = :pass ",  
		map[string]interface{}{
	  		"user": postedUser.Username,
	  		"pass": postedUser.Password,
		})
	checkErr(err, "LOGIN FAILED MISERABLY")

	return &user
}

func DoLogin(entity models.User, session sessions.Session, enc Encoder, db gorp.SqlExecutor) (int, string){

	fmt.Println(entity)
	fmt.Println("-------")
	fmt.Println(session)
	fmt.Println("-------")
	fmt.Println(enc)
	fmt.Println("-------")
	fmt.Println(db)
	fmt.Println("-------")
	//fmt.Println(sessionauth)
	fmt.Println("-------")

	
	status := &models.Message{
			Id:       http.StatusUnauthorized,
			Name:     "Usuário ou Senha errado.",
			Lang_key: "br",
		}
		
	user := GetUserByLogin(entity)
		
		
	if user.Name == "" {
		return http.StatusUnauthorized, Must(enc.EncodeOne(status))
	} else {
		err := sessionauth.AuthenticateSession(session, user)
		if err != nil {
			status.Name = "Erro ao iniciar Sessão."	
			return http.StatusUnauthorized, Must(enc.EncodeOne(status))
		}			
		status.Name = "Login realizado com sucesso!"
		return http.StatusOK, Must(enc.EncodeOne(status))
		
	}

	return http.StatusUnauthorized, Must(enc.EncodeOne(status))
}