
package server

	import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/", s.handler) 
	mux.HandleFunc("/about", s.aboutHandler)
mux.HandleFunc("/contact", s.contactHandler)

	return mux
}
func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, " <h1> Welcome!</h1> ") 
}

func (s *Server) aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/about.html")
	if err != nil {
		log.Fatal("Error while parsing file")
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func (s *Server) contactHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/contact.html")
	if err != nil {
		log.Fatal("Error while parsing file")
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


