package worker

import (
	"log"

	"github.com/Unknow101/GoJWTBrute2/jwtutils"
)

type Candidate struct {
	Jwt jwtutils.JWT
	Key string
}

func Worker(id int, input <-chan Candidate, result chan<- string) {
	for candidat := range input {
		s1, err := jwtutils.CalculateSignature(candidat.Jwt, candidat.Key)
		if err != nil {
			log.Fatal(err)
		}
		if s1 == candidat.Jwt.Signature {
			result <- candidat.Key
		}
	}
}
