package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gustavoaniel/go-api-rest/controllers"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidade", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidade/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidade", controllers.CriaUmaNovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidade/{id}", controllers.DeletaUmaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidade/{id}", controllers.EditaUmaPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}
