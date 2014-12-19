package models

import (	
	"github.com/martini-contrib/sessionauth"	
)

// User can be any struct that represents a user in my system
type User struct {
	Id            	int64  `json:"id"`
	Username      	string `json:"username"`
	Password      	string `json:"password"`
	Name			string `json:"name"`
	Email			string `json:"email"`
	LastLogin		string `json:"last_login"`
	UserHole 		string `json:"user_hole"`
	Authenticated 	bool   `json:"authenticated"`
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