package hashing

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type response struct {
	Hash string 	`json:"hash,omitempty"`
	Err  string  `json:"error,omitempty"`
}

type request struct {
	Key string `json:"key"`
	HashFunc string `json:"hash_function"`
}

func (r *request) FromJSON( reader io.Reader ) (err error)  {

	content , err := ioutil.ReadAll( reader )
	if err != nil {
		return
	}

	err = json.Unmarshal( content, r )
	return

}

func (r *response) ToJSON( writer io.Writer ) (err error)  {

	responseByte, err := json.Marshal( r )
	if  err != nil {
		return
	}

	_, err = writer.Write( responseByte )

	return

}