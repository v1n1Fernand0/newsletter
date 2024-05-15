package main

import (
	"log"
	"net/http"
	"newsletter/internal/handlers"

	_ "newsletter/docs" // swagger docs

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger" // http-swagger middleware
)

// @title Newsletter API
// @version 1.0
// @description API para gerenciar a configuração e envio de newsletters.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/config", handlers.GetConfig).Methods("GET")
	r.HandleFunc("/api/config", handlers.UpdateConfig).Methods("POST")
	r.HandleFunc("/api/newsletter", handlers.SendNewsletter).Methods("POST")

	r.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", http.FileServer(http.Dir("./docs"))))

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Printf("Servidor iniciando na porta :8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
