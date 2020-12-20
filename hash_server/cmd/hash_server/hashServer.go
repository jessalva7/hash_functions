package main

import (
	"github.com/jessalva7/hash_functions/hash_server/pkg/hashing"
	"log"
	"net/http"
	"os"
)

func main(){

	addr := os.Getenv("SERVER_PORT")

	hashingServerMux := http.NewServeMux()
	hashingStrategy := hashing.NewHashFunctionStrategy()
	hashingServerMux.Handle( "/", hashing.NewHandler(hashingStrategy) )

	hashServer := http.Server{
		Handler: hashingServerMux,
		Addr:addr,
	}

	if err := hashServer.ListenAndServe(); err != nil {

		log.Fatal("server stopped", err)

	}

}
