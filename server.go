package main

import (
	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/sessionauth"
	"github.com/martini-contrib/sessions"
	"gitlab.com/warnabroda/warnabrodagomartini/models"
	"gitlab.com/warnabroda/warnabrodagomartini/routes"
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
		MaxAge: 3600,
	})

	m.Use(sessions.Sessions("admin_session", store))
	m.Use(sessionauth.SessionUser(models.GenerateAnonymousUser))
	sessionauth.RedirectUrl = "/hq"

	r.Group("/warnabroda", func(r martini.Router) {

		r.Get(`/messages/:lang_key`, routes.GetMessages)
		r.Get(`/contact_types`, routes.GetContact_types)
		r.Get(`/subjects`, routes.GetSubjects)
		r.Post(`/captcha`, binding.Json(models.Captcha{}), routes.CaptchaResponse)

		r.Post(`/warnings`, binding.Json(models.Warning{}), routes.AddWarning)
		r.Get(`/warnings/counter`, routes.CountSentWarnings)
		r.Post(`/warnings/delivery`, binding.Json(models.DefaultStruct{}), routes.SendConfirmation)

		r.Post(`/ignore-list`, binding.Json(models.Ignore_List{}), routes.AddIgnoreList)
		r.Put(`/ignore-list`, binding.Json(models.Ignore_List{}), routes.ConfirmIgnoreList)

		r.Get(`/reply/:hash`, routes.GetReplyByHash)
		r.Post(`/reply`, binding.Json(models.WarningResp{}), routes.SetReply)
		r.Put(`/reply`, binding.Json(models.WarningResp{}), routes.ReadReply)

		r.Group("/hq", func(r martini.Router) {

			r.Get(`/auth-on`, routes.GetAuthenticatedUser)
			r.Post(`/login`, binding.Json(models.UserLogin{}), routes.DoLogin)
			r.Get(`/logout`, routes.DoLogout)
			r.Get(`/user/private`, routes.IsAuthenticated)

			r.Get(`/account/:id`, routes.GetUserById)

			r.Get(`/totals`, routes.WarnaCounter)

			r.Get(`/warnings`, binding.Json(models.Warn{}), routes.ListWarnings)
			r.Get(`/warnings/:id`, routes.GetWarningDetail)

			r.Post(`/messages`, binding.Json(models.MessageStruct{}), routes.SaveOrUpdateMessage)
			r.Get(`/messages/:id`, routes.GetMessage)
			r.Get(`/stats`, routes.GetMessagesStats)

		})

	})

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
