package hashing

import (
	"fmt"
	"log"
	"net/http"
)

type handler struct {
	hashStrategy HashFunctionStrategy
}

func NewHandler( hashStrategy HashFunctionStrategy ) http.Handler{
	return &handler{hashStrategy: hashStrategy}
}

func (h *handler)  ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	resp := response{}
	hashingRequest := request{}
	const ApplicationJson = "application/json"
	if r.Method == http.MethodPost {

		requestBody := r.Body
		err := hashingRequest.FromJSON(requestBody)

		if err != nil {
			log.Print("Error while parsing Input")
			resp.Err = err.Error()
			ServerRespond(resp, rw, http.StatusBadRequest, ApplicationJson)
			return
		}


		hashFunc, err := h.hashStrategy.GetHashFunction( hashingRequest.HashFunc )
		if err != nil {
			log.Print("Error while finding hash function")
			resp.Err = err.Error()
			ServerRespond(resp, rw, http.StatusBadRequest, ApplicationJson)
			return
		}

		hashValue := hashFunc.hash( []byte(hashingRequest.Key) )
		resp.Hash= fmt.Sprintf("%x",hashValue)
		ServerRespond(resp, rw, http.StatusOK, ApplicationJson)

		return
	}

	resp.Err = "We don't support this yet :)"
	ServerRespond(resp, rw, http.StatusMethodNotAllowed, ApplicationJson)

}

func ServerRespond(resp response, rw http.ResponseWriter, status int, contentType string)  {


	rw.WriteHeader(status)
	rw.Header().Set("Content-type", contentType)

	err := resp.ToJSON( rw )
	if err != nil {

		_, _ = rw.Write([]byte("Something went wrong"))
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Header().Set("Content-type", contentType)
	}



}
