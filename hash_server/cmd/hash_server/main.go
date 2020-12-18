package hash_server

import (
	"log"
	"net/http"
	"os"
)

func main(){

	addr := os.Getenv("SERVER_PORT")

	if err := http.ListenAndServe(addr, nil); err != nil {

		log.Fatal("server stopped")

	}

}
