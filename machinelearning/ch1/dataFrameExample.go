// Package main provides ...
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("I need one csv file")
	}

	irisFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	irisDataFrame := dataframe.ReadCSV(irisFile)

	fmt.Println(irisDataFrame)

	filter := dataframe.F{"species", "==", "Iris-versicolor"}
	verticolorDF := irisDataFrame.Filter(filter).Select([]string{"sepal_width", "species"})
	fmt.Println(verticolorDF)
}
