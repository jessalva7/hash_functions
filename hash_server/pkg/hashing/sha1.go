package hashing

import (
	"crypto/sha1"
	"hash"
	"log"
)

type sha1HashFunction struct {
	hash.Hash
}

func NewSha1HashFunction() HashFunction{

	return &sha1HashFunction{sha1.New()}

}

func (s *sha1HashFunction) hash( input []byte) []byte {

	_, err := s.Write( input )
	if err != nil {

		log.Print( err )
		return []byte( "Error while hashing" )

	}

	return s.Sum( nil )

}
