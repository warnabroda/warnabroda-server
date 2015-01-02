package main

import (
	"bitbucket.org/hbtsmith/warnabrodagomartini/models"
	"bitbucket.org/hbtsmith/warnabrodagomartini/routes"
	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"	
	"github.com/martini-contrib/sessionauth"
	"github.com/martini-contrib/sessions"
	// "github.com/martini-contrib/strict"
	"net/http"
	"regexp"
	"strings"
	// "fmt"
)

// The one and only martini instance.
var m *martini.Martini

func init() {
	m = martini.New()
	// Setup middleware
	m.Use(martini.Recovery())
	m.Use(martini.Logger())
	m.Use(martini.Static("public"))
	m.Use(MapEncoder)
	// m.Use(strict.Strict)
	// Setup routes
	r := martini.NewRouter()

	//warnasecretkey
	store := sessions.NewCookieStore([]byte("799a41cbe4de9a67eaa42acc83c76be7aa57e684"))
	store.Options(sessions.Options{
		MaxAge: 6000,
	})

	m.Use(sessions.Sessions("admin_session", store))
	m.Use(sessionauth.SessionUser(models.GenerateAnonymousUser))
	sessionauth.RedirectUrl = "/hq"	

	r.Group("/warnabroda", func (r martini.Router){
		r.Get(`/messages`, routes.GetMessages)
		r.Get(`/contact_types`, routes.GetContact_types)
		r.Get(`/subjects`, routes.GetSubjects)
		r.Get(`/count-sent-warnings`, routes.CountSentWarnings)
		r.Post(`/warnings`, binding.Json(models.Warning{}), routes.AddWarning)
		r.Post(`/ignore-list`, binding.Json(models.Ignore_List{}), routes.AddIgnoreList)
		r.Post(`/ignore-list-confirm`, binding.Json(models.Ignore_List{}), routes.ConfirmIgnoreList)
		r.Post(`/captcha-validate`, binding.Json(models.Captcha{}), routes.CaptchaResponse)

		r.Group("/hq", func (r martini.Router){

			r.Get(`/account/:id`, routes.GetUserById)	
			r.Get(`/private`, routes.IsAuthenticated)

			r.Get(`/logout`, routes.DoLogout)

			r.Get(`/authenticated-user`, routes.GetAuthenticatedUser)

			// r.Post(`/authentication`, strict.ContentType("application/x-www-form-urlencoded"), binding.Form(models.UserLogin{}), routes.DoLogin)	
			r.Post(`/authentication`, binding.Json(models.UserLogin{}), routes.DoLogin)				
			r.Get(`/count-warnings`, routes.CountWarns)
			r.Get(`/list-warnings`, binding.Json(models.Warn{}), routes.ListWarnings)
			
		})

	})

	// r.Get(`/warnabroda/messages/:id`, routes.GetMessage)
	// r.Post(`/warnabroda/messages`, binding.Json(models.Message{}), routes.AddMessage)
	// r.Put(`/warnabroda/messages/:id`, binding.Json(models.Message{}), routes.UpdateMessage)
	// r.Delete(`/warnabroda/messages/:id`, routes.DeleteMessage)

	// r.Get(`/warnabroda/contact_types/:id`, routes.GetContact_type)
	// r.Post(`/warnabroda/contact_types`, binding.Json(models.Contact_type{}), routes.AddContact_type)
	// r.Put(`/warnabroda/contact_types/:id`, binding.Json(models.Contact_type{}), routes.UpdateContact_type)
	// r.Delete(`/warnabroda/contact_types/:id`, routes.DeleteContact_type)

	// r.Get(`/warnabroda/subjects/:id`, routes.GetSubject)
	// r.Post(`/warnabroda/subjects`, binding.Json(models.Subject{}), routes.AddSubject)
	// r.Put(`/warnabroda/subjects/:id`, binding.Json(models.Subject{}), routes.UpdateSubject)
	// r.Delete(`/warnabroda/subjects/:id`, routes.DeleteSubject)

	// r.Get(`/warnabroda/warnings`, routes.GetWarnings)
	// r.Get(`/warnabroda/warnings/:id`, routes.GetWarning)
	// r.Put(`/warnabroda/warnings/:id`, binding.Json(models.Warning{}), routes.UpdateWarning)
	// r.Delete(`/warnabroda/warnings/:id`, routes.DeleteWarning)


	// r.Get(`/warnabroda/ignore-list/:id`, routes.GetIgnoreContact)
	
	// r.Put(`/warnabroda/ignore_list/:id`, binding.Json(models.Subject{}), routes.UpdateIgnoreList)
	// r.Delete(`/warnabroda/ignore_list/:id`, routes.DeleteIgnoreList)	

	// Inject database
	m.MapTo(models.Dbm, (*gorp.SqlExecutor)(nil))
	// Add the router action
	m.Action(r.Handle)	
}

// The regex to check for the requested format (allows an optional trailing
// slash).
var rxExt = regexp.MustCompile(`(\.(?:xml|text|json))\/?$`)

// MapEncoder intercepts the request's URL, detects the requested format,
// and injects the correct encoder dependency for this request. It rewrites
// the URL to remove the format extension, so that routes can be defined
// without it.
func MapEncoder(c martini.Context, w http.ResponseWriter, r *http.Request) {
	// Get the format extension
	matches := rxExt.FindStringSubmatch(r.URL.Path)
	ft := ".json"
	if len(matches) > 1 {
		// Rewrite the URL without the format extension
		l := len(r.URL.Path) - len(matches[1])
		if strings.HasSuffix(r.URL.Path, "/") {
			l--
		}
		r.URL.Path = r.URL.Path[:l]
		ft = matches[1]
	}
	// Inject the requested encoder
	switch ft {
	case ".xml":
		c.MapTo(routes.XmlEncoder{}, (*routes.Encoder)(nil))
		w.Header().Set("Content-Type", "application/xml")
	case ".text":
		c.MapTo(routes.TextEncoder{}, (*routes.Encoder)(nil))
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	default:
		c.MapTo(routes.JsonEncoder{}, (*routes.Encoder)(nil))
		w.Header().Set("Content-Type", "application/json")
	}
}

func main() {

	m.Run()

}