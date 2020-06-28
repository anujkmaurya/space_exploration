package delivery

import (
	"github.com/personal-work/space_exploration/internal/usecase"

	"github.com/gorilla/mux"
)

//InitHandlers inits all the routes
func InitHandlers() *mux.Router {

	r := mux.NewRouter()

	//healthcheck api
	r.HandleFunc("/healthcheck", usecase.HealthCheck).Methods("GET")

	//
	r.HandleFunc("/sectors/{", controllers.CreateUser).Methods("GET")
	return r
}
