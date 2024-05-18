package constants

const (
	PAGESPATH       = "pages/"
	ROUTESFILEPATH  = "internal/server/routes.go"
	ROUTETEMPLATE   = "mux.HandleFunc(\"/%[1]s\", s.%[2]s)"
	HANDLERTEMPLATE = `
func (s *Server) %[1]sHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/%[2]s.html")
	if err != nil {
		log.Fatal("Error while parsing file")
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
`
	ROUTESFILETEMPLATE = `
package server

	import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	%[1]s
	return mux
}
%[2]s
`
)

const (
	NOFILESROUTETEMPLATE = `
	mux.HandleFunc("/", s.handler) 
	`
	NOFILESHANDLERTEMPLATE = `func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, " <h1> Welcome!</h1> ") 
}
`
)
