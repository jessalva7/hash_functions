package hashing

import (
	"crypto/sha256"
	"hash"
	"log"
)

type sha256HashFunction struct {
	hash.Hash
}

func NewSha256HashFunction() HashFunction{

	return &sha256HashFunction{sha256.New()}

}

func (s *sha256HashFunction) hash( input []byte) []byte {

	_, err := s.Write( input )
	if err != nil {

		log.Print( err )
		return []byte( "Error while hashing" )

	}

	return s.Sum( nil )

}
