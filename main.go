package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Unknow101/GoJWTBrute2/jwtutils"
	"github.com/Unknow101/GoJWTBrute2/utils"
	"github.com/Unknow101/GoJWTBrute2/worker"
)

func PrepareWorkerPool(poolsize int, jwt, wordlist string) error {
	input := make(chan worker.Candidate, 100000)
	results := make(chan string, 100000)
	for w := 1; w <= poolsize; w++ {
		go worker.Worker(w, input, results)
	}
	jwtStruct, err := jwtutils.ParseToken(jwt)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(wordlist)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		candidate := worker.Candidate{jwtStruct, scanner.Text()}
		input <- candidate
	}
	close(input)
	secretKey := <-results
	close(results)
	utils.FoundKey(secretKey)
	return nil
}

func main() {
	start := time.Now()
	fmt.Println("GoSMBBrute v0.2")
	jwt := flag.String("jwt", "", "JWT to crack")
	wordlist := flag.String("w", "", "Password wordlist to test")
	flag.Parse()
	if *jwt == "" || *wordlist == "" {
		fmt.Println("[-] You have to specify a jwt and a wordlist")
		os.Exit(1)
	}
	PrepareWorkerPool(10, *jwt, *wordlist)
	elapsed := time.Since(start)
	log.Printf("Program took %s", elapsed)

}
