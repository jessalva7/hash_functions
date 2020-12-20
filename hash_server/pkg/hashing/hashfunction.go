package hashing

import (
	"errors"
	"hash"
)

type HashFunction interface {

	hash.Hash
	hash([]byte) []byte

}

type HashFunctionStrategy struct {
	hashFunctionMap map[string]HashFunction
}

func NewHashFunctionStrategy() HashFunctionStrategy {

	hashMappings := make( map[string]HashFunction )
	hashMappings[ "SHA256" ] = NewSha256HashFunction()
	hashMappings[ "SHA1" ] = NewSha1HashFunction()

	return HashFunctionStrategy{hashFunctionMap:  hashMappings}


}

func ( hfs *HashFunctionStrategy) GetHashFunction( hashFunc string ) (hashFun HashFunction,err error){

	hashFun = hfs.hashFunctionMap[hashFunc]
	if hashFun == nil {
		err = errors.New( "no hashing method found" )
	}
	return

}