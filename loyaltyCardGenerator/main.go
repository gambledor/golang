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

func main() {
	fd, err := os.Create("loyaltyCards.csv")
	checkError(err)
	writeBuffer := bufio.NewWriter(fd)
	defer fd.Close()

	rand.Seed(31)
	for i := 0; i < BATCH_CARDINALITY; i++ {
		fmt.Fprintf(writeBuffer, "%016s\n", strconv.Itoa(rand.Intn(MAX_CARD_NUMBER)))
	}
	writeBuffer.Flush()
}
