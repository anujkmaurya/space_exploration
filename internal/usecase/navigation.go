package usecase

import (
	"net/http"
)

//HealthCheck to check health of service
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("tested OK\n"))
}
