package routes

import (
	"bitbucket.org/hbtsmith/warnabrodagomartini/models"
	"bitbucket.org/hbtsmith/warnabrodagomartini/i18n"	
	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"	
	"github.com/martini-contrib/sessionauth"
	"github.com/martini-contrib/sessions"
	"net/http"
	"strconv"
//	"fmt"
	// "encoding/json"
	// "strings"
)

const (
	SQL_LOGIN				= "SELECT * FROM users WHERE (username = :user OR email = :user) AND password = :pass "
	MSG_LOGIN_INVALID		= "Usuário ou Senha inválidos."
	MSG_LOGIN_REQUIRED		= "Usuário não está logado."
	MSG_SESSION_INIT_ERROR	= "Erro ao iniciar Sessão."
	MSG_SUCCESSFUL_LOGIN	= "Login realizado com sucesso!"
	MSG_SUCCESSFUL_LOGOUT	= "Logout realizado com sucesso."
)

func GetUserById(enc Encoder, db gorp.SqlExecutor, parms martini.Params) (int, string) {
	
	id, err := strconv.Atoi(parms["id"])	

	if err != nil {
		checkErr(err, "GET USER ERROR")
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

func GetUserByLogin(postedUser models.UserLogin, db gorp.SqlExecutor) *models.User {	

	user := models.User{}

	db.SelectOne(&user, SQL_LOGIN,  
		map[string]interface{}{
	  		"user": postedUser.Username,
	  		"pass": postedUser.Password,
		})


	return &user
}

func DoLogin(entity models.UserLogin, session sessions.Session, enc Encoder, db gorp.SqlExecutor) (int, string){
	
	status := &models.DefaultStruct{
			Id:       http.StatusUnauthorized,
			Name:     MSG_LOGIN_INVALID,
			Lang_key: i18n.BR_LANG_KEY,
		}
		
	user := GetUserByLogin(entity, db)

	if user.Name != "" {

		err := sessionauth.AuthenticateSession(session, user)
		if err != nil {
			status.Name = MSG_SESSION_INIT_ERROR	
			return http.StatusForbidden, Must(enc.EncodeOne(status))
		}
		user.Authenticated = true	
		user.UpdateLastLogin()
		status.Name = MSG_SUCCESSFUL_LOGIN
		return http.StatusOK, Must(enc.EncodeOne(user))		
	
	} else {		
	
		sessionauth.Logout(session, user)
		session.Clear()
		return http.StatusForbidden, Must(enc.EncodeOne(status))
	
	}

	return http.StatusForbidden, Must(enc.EncodeOne(status))
}

func IsAuthenticated(enc Encoder, user sessionauth.User) (int, string) {	

	if user.IsAuthenticated(){
		return http.StatusOK,  ""
	}


	return http.StatusUnauthorized,  Must(enc.EncodeOne(user))
}

func DoLogout(enc Encoder, session sessions.Session, user sessionauth.User, db gorp.SqlExecutor) (int, string) {

	status := &models.DefaultStruct{
			Id:       http.StatusOK,
			Name:     MSG_LOGIN_REQUIRED,
			Lang_key: i18n.BR_LANG_KEY,
		}

	if user.IsAuthenticated() {

		sessionauth.Logout(session, user)
		session.Clear()
		status.Name = MSG_SUCCESSFUL_LOGOUT
	}	

	updateUser := UserById(user.UniqueId().(int), db)

	updateUser.Authenticated = false
	db.Update(updateUser)

	return http.StatusOK,  Must(enc.EncodeOne(status))
}


func GetAuthenticatedUser(enc Encoder, user sessionauth.User, db gorp.SqlExecutor) (int, string) {
	
	if user.IsAuthenticated() {

		authUser := UserById(user.UniqueId().(int), db)

		return http.StatusOK,  Must(enc.EncodeOne(authUser))
	} 

	return http.StatusUnauthorized,  Must(enc.EncodeOne(user))
}