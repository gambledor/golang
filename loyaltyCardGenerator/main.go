package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

const MAX_CARD_NUMBER = 9999999999999999
const BATCH_CARDINALITY = 50000

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

// card number producer which writes them to the channel out
func worker(out chan<- string) {
	rand.Seed(111)
	for i := 0; i < BATCH_CARDINALITY; i++ {
		out <- strconv.Itoa(rand.Intn(MAX_CARD_NUMBER))
	}
	close(out)
}

func main() {
	const CHANNEL_LENGTH = 1000
	var channel = make(chan string, CHANNEL_LENGTH)

	go worker(channel)
	fd, err := os.Create("loyaltyCards.csv")
	checkError(err)
	writeBuffer := bufio.NewWriter(fd)
	defer fd.Close()

	for v := range channel {
		fmt.Fprintf(writeBuffer, "%016s\n", v)
	}
	writeBuffer.Flush()
}
