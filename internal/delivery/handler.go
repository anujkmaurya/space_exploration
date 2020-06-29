package delivery

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/personal-work/space_exploration/internal/models"
)

//HandlerFunc : handler return the data and error
type HandlerFunc func(rw http.ResponseWriter, r *http.Request) (interface{}, error)

func (fn HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	response := models.Response{}

	//start time
	start := time.Now()

	var data interface{}
	var err error

	errStatus := http.StatusOK

	data, err = fn(w, r)

	response.Header.ProcessTime = time.Since(start).String()

	var buf []byte

	w.Header().Set("Content-Type", "application/json")

	//if err is nil and data is present
	if data != nil && err == nil {
		response.Data = data
		if buf, err = json.Marshal(response); err == nil {
			_, err := w.Write(buf)
			if err != nil {
				log.Printf("error in writing http response, err : %+v\n", err)
				return
			}
		}
	}

	//check for error, if yes, fill error message and status code
	if err != nil {

		switch t := err.(type) {
		case *models.AppError:
			//fill error response here
			response.Header.ErrorMessage = t.ErrorMessage
			errStatus = t.HTTPResposeCode
			log.Printf("Error http Code: %+v, Message: %+v\n", t.HTTPResposeCode, t.ErrorMessage)
		}
		w.WriteHeader(errStatus)
	} else {
		return
	}

	buf, err = json.Marshal(response)
	if err != nil {
		log.Printf("error in marshing reponse:  %+v\n", err)
	}

	_, err = w.Write(buf)
	if err != nil {
		log.Printf("[ServeHTTP] [ERROR] in writing http response\n")
	}
}
