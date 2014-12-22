package models

import (	
	"github.com/martini-contrib/sessionauth"
	"time"
)

// User can be any struct that represents a user in my system
type User struct {
	Id            	int64  `json:"id" db:"id"`
	Username      	string `json:"username" db:"username"`
	Password      	string `json:"-" db:"password"`
	Name			string `json:"name" db:"name"`
	Email			string `json:"email" db:"email"`
	Last_login		string `json:"last_login" db:"last_login"`
	User_hole 		string `json:"user_hole" db:"user_hole"`
	Authenticated 	bool   `json:"authenticated" db:"authenticated"`
}

type UserLogin struct {
	Username      	 string `json:"username" form:"username"`
	Password      	 string `json:"password" form:"password"`
	Ip 				 string `json:"ip" form:"ip"`  
    Browser 		 string `json:"browser" form:"browser"`  
    Operating_system string `json:"operating_system" form:"operating_system"`  
    Device 			 string `json:"device" form:"device"`  
    Raw 			 string `json:"raw" form:"raw"`
}

// GetAnonymousUser should generate an anonymous user model
// for all sessions. This should be an unauthenticated 0 value struct.
func GenerateAnonymousUser() sessionauth.User {
	return &User{}
}

// Login will preform any actions that are required to make a user model
// officially authenticated.
func (u *User) Login() {
	// Update last login time
	// Add to logged-in user's list
	// etc ...
	u.Authenticated = true
}

// Logout will preform any actions that are required to completely
// logout a user.
func (u *User) Logout() {
	// Remove from logged-in user's list
	// etc ...
	u.Authenticated = false
}

func (u *User) IsAuthenticated() bool {
	return u.Authenticated
}

func (u *User) UniqueId() interface{} {
	return u.Id
}

// GetById will populate a user object from a database model with
// a matching id.
func (u *User) GetById(id interface{}) error {
	err := Dbm.SelectOne(u, "SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) UpdateLastLogin(){
	u.Last_login = time.Now().String()
	Dbm.Update(u)
}