package routes

import (
	"time"	

	"bitbucket.org/hbtsmith/warnabrodagomartini/models"
//	"fmt"	
)


const (
	SQL_REMOVE_OLD_IGNOREME_REQUESTS		= "DELETE FROM ignore_list WHERE confirmed = false AND (created_date + INTERVAL 24 HOUR) < NOW()"
	SQL_LOGOFF_EXPIRED_SESSION_USERS		= "UPDATE users SET authenticated = FALSE WHERE (last_login + INTERVAL 1 HOUR) < NOW()"
)

//Initialize all required functions when container is up.
func init(){
	HourlyTimer()	
}

// Run hourly tasks:
// - Clean unconfirmed ignore-me requests
// - Logoff users with expired sessions
func HourlyTimer(){
	
	models.Dbm.Exec(SQL_REMOVE_OLD_IGNOREME_REQUESTS)
	models.Dbm.Exec(SQL_LOGOFF_EXPIRED_SESSION_USERS)
	
	ticker := time.NewTicker(time.Hour)
	quit := make(chan struct{})
	go func() {
	    for {
	       select {
	        case <- ticker.C:	        	
	            models.Dbm.Exec(SQL_REMOVE_OLD_IGNOREME_REQUESTS)
	            models.Dbm.Exec(SQL_LOGOFF_EXPIRED_SESSION_USERS)	        	
	        case <- quit:
	            ticker.Stop()
	            return
	        }
	    }
	 }()
}