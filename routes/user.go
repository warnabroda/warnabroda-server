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
	// "encoding/json"
	// "strings"
)

func GetUserById(enc Encoder, db gorp.SqlExecutor, parms martini.Params) (int, string) {
	
	id, err := strconv.Atoi(parms["id"])	

	if err != nil {
		checkErr(err, "get failed")
	}
	
	entity := UserById(id, db)

	return http.StatusOK, Must(enc.EncodeOne(entity))
}

func UserById(id int, db gorp.SqlExecutor) *models.User {
	
	obj, err := db.Get(models.User{}, id)
	if err != nil {
		return nil
	}
	entity := obj.(*models.User)

	return entity
}

func GetUserByLogin(postedUser models.UserLogin) *models.User {	

	user := models.User{}

	models.Dbm.SelectOne(&user, "SELECT * FROM users WHERE (username = :user OR email = :user) AND password = :pass ",  
		map[string]interface{}{
	  		"user": postedUser.Username,
	  		"pass": postedUser.Password,
		})
	//checkErr(err, "LOGIN FAILED MISERABLY")

	return &user
}

func DoLogin(entity models.UserLogin, r *http.Request, session sessions.Session, enc Encoder, db gorp.SqlExecutor) (int, string){
	
	status := &models.Message{
			Id:       http.StatusUnauthorized,
			Name:     "Usuário ou Senha errado.",
			Lang_key: "br",
		}
		
	user := GetUserByLogin(entity)
		
		
	if user.Name == "" {
		sessionauth.Logout(session, user)
		session.Clear()
		return http.StatusUnauthorized, Must(enc.EncodeOne(status))
	} else {		
		err := sessionauth.AuthenticateSession(session, user)
		if err != nil {
			status.Name = "Erro ao iniciar Sessão."	
			return http.StatusUnauthorized, Must(enc.EncodeOne(status))
		}
		user.Authenticated = true	
		user.UpdateLastLogin()
		status.Name = "Login realizado com sucesso!"
		return http.StatusOK, Must(enc.EncodeOne(user))
		
	}

	return http.StatusUnauthorized, Must(enc.EncodeOne(status))
}

func IsAuthenticated(enc Encoder, user sessionauth.User) (int, string) {

	fmt.Println(user)

	if user.IsAuthenticated(){
		return http.StatusOK,  ""
	}


	return http.StatusUnauthorized,  Must(enc.EncodeOne(user))
}

func DoLogout(enc Encoder, session sessions.Session, user sessionauth.User, db gorp.SqlExecutor) (int, string) {

	status := &models.Message{
			Id:       http.StatusOK,
			Name:     "Usuario não está logado.",
			Lang_key: "br",
		}

	if user.IsAuthenticated() {

		sessionauth.Logout(session, user)
		session.Clear()
		status.Name = "Logout Realizado com sucesso."

	}
	

	updateUser := UserById(user.UniqueId().(int), db)

	updateUser.Authenticated = false
	db.Update(updateUser)

	return http.StatusOK,  Must(enc.EncodeOne(status))
}