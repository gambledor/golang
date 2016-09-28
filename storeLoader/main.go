// Package main provides ...
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

// File parser
func fileParser(file string) string {
	fd, err := os.Open(file)
	checkError(err)
	defer fd.Close()

	procedure := "DECLARE\nloc_nextval NUMBER;\nmerchant_id NUMBER;\nBEGIN\n"
	procedure = procedure + "SELECT ID_MERCHANT into merchant_id FROM PPT_MERCHANT WHERE CUSTOMER_ID = 2;\n\n"
	storeInsert := []string{"INSERT INTO PPT_STORE (STORE_ID, COMPANY_NAME, FK_LOCATION, MERCHANT_ID, FK_CUSTOMER, TYPE, EXTERNAL_ID, NAME) VALUES\n"}
	locationInsert := []string{"loc_nextval := SEQ_LOCATION.nextval;\nINSERT INTO SCQK_LOCATION (ID_LOCATION, ADDRESS, STATE_PROVINCE, CITY, ZIP_CODE, COUNTRY, LATITUDE, LONGITUDE) VALUES\n"}

	var inserts string
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, "[", "", 1)
		line = strings.Replace(line, "]", "", 1)

		var locs []string
		var stores []string
		var words []string

		words = strings.FieldsFunc(line, f)
		for i := range words {
			tokens := strings.Split(words[i], "=")
			key := strings.Trim(tokens[0], " ")
			if key == "id" {
				stores = append(storeInsert, "(SEQ_PPT_STORE.nextval,")
				stores = append(stores, "'RoadhHouse Grill',")
				stores = append(stores, "loc_nextval,")
				stores = append(stores, "merchant_id,")
				stores = append(stores, "2,")
				stores = append(stores, "0,")
				s := fmt.Sprintf("'%s'", tokens[1]) + ","
				// external_id
				stores = append(stores, s)
			}
			// There is a useless description at the end of the file
			if key == "description" && i < 2 {
				s := fmt.Sprintf("'%s'", tokens[1]) + ");\n"
				stores = append(stores, s)
			}
			if key == "address" { // The first field for laction tuple
				locs = append(locationInsert, "(loc_nextval,")
				s := fmt.Sprintf("'%s',", tokens[1])
				locs = append(locs, s)
			}
			if key == "province" { // province -> state_province
				s := fmt.Sprintf("'%s',", tokens[1])
				locs = append(locs, s)
			}
			if key == "city" {
				s := fmt.Sprintf("'%s',", tokens[1])
				locs = append(locs, s)
			}
			if key == "zipcode" {
				s := fmt.Sprintf("'%s',", tokens[1])
				locs = append(locs, s)
			}
			if key == "country" {
				s := fmt.Sprintf("'%s',", strings.Replace(tokens[1], "A", "", 1))
				locs = append(locs, s)
			}
			if key == "lat" {
				value := tokens[1]
				// removes last 2 ditits
				s := fmt.Sprintf("%s,", value[:len(value)-2])
				locs = append(locs, s)
			}
			if key == "lng" {
				value := tokens[1]
				// removes last 2 ditits
				s := fmt.Sprintf("%s);\n", value[:len(value)-2])
				locs = append(locs, s)
			}
		}
		inserts = inserts + strings.Join(locs, "") + strings.Join(stores, "") + "\n\n"
	}
	procedure = procedure + inserts + "\nrollback;\nend;"

	return procedure
}

// The split rune is comma
func f(c rune) bool {
	return c == ','
}

func main() {

	fileFlag := flag.Bool("file", false, "The store file to load")
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("\t%s --file <STORES_FILE>\n", os.Args[0])
		fmt.Println("record format:")
		fmt.Println("[id=10078, description=ARESE - RISTORANTE, address=Via Giuseppe Eugenio Luraghi, province=MI, city=Arese, zipcode=20020, country=ITA, lat=45.56170833, lng=9.05189444],")
	}

	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("file missing")
		os.Exit(2)
	}
	if *fileFlag {
		fmt.Printf("%s", fileParser(os.Args[2]))
	}
}
