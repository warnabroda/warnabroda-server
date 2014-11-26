package main

import (
    "warnabroda/models"
    "warnabroda/routes"
    //"log"
    "net/http"
    "regexp"
    "strings"
    "github.com/go-martini/martini"
    "github.com/martini-contrib/binding"
    "github.com/coopernurse/gorp"
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
    // Setup routes
    r := martini.NewRouter()
    
    r.Get(`/warnabroda/messages`, routes.GetMessages)
    // r.Get(`/warnabroda/messages/:id`, routes.GetMessage)
    // r.Post(`/warnabroda/messages`, binding.Json(models.Message{}), routes.AddMessage)
    // r.Put(`/warnabroda/messages/:id`, binding.Json(models.Message{}), routes.UpdateMessage)
    // r.Delete(`/warnabroda/messages/:id`, routes.DeleteMessage)
    
    r.Get(`/warnabroda/contact_types`, routes.GetContact_types)
    // r.Get(`/warnabroda/contact_types/:id`, routes.GetContact_type)
    // r.Post(`/warnabroda/contact_types`, binding.Json(models.Contact_type{}), routes.AddContact_type)
    // r.Put(`/warnabroda/contact_types/:id`, binding.Json(models.Contact_type{}), routes.UpdateContact_type)
    // r.Delete(`/warnabroda/contact_types/:id`, routes.DeleteContact_type)
    
    r.Get(`/warnabroda/subjects`, routes.GetSubjects)
    // r.Get(`/warnabroda/subjects/:id`, routes.GetSubject)
    // r.Post(`/warnabroda/subjects`, binding.Json(models.Subject{}), routes.AddSubject)
    // r.Put(`/warnabroda/subjects/:id`, binding.Json(models.Subject{}), routes.UpdateSubject)
    // r.Delete(`/warnabroda/subjects/:id`, routes.DeleteSubject)
    
    // r.Get(`/warnabroda/warnings`, routes.GetWarnings)
    // r.Get(`/warnabroda/warnings/:id`, routes.GetWarning)
    r.Post(`/warnabroda/warnings`, binding.Json(models.Warning{}), routes.AddWarning)
    // r.Put(`/warnabroda/warnings/:id`, binding.Json(models.Warning{}), routes.UpdateWarning)
    // r.Delete(`/warnabroda/warnings/:id`, routes.DeleteWarning)
    
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
