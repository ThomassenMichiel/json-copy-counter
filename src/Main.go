package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Copies struct {
	Entries []Entry `json:"Kopies_View"`
}

type Entry struct {
	Amount       string `json:"Hoeveelheid"`
	Type         string `json:"Kopie"`
	Organisation string `json:"Organisatie"`
}

type Organisation struct {
	Name   string
	Totals map[string]int
}

func main() {
	jsonFile, err := os.Open("copies.json")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			os.Exit(1)
		}
	}(jsonFile)

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result Copies
	organisations := make(map[string]Organisation)

	err = json.Unmarshal(byteValue, &result)

	if err != nil {
		os.Exit(1)
	}

	for i := 0; i < len(result.Entries); i++ {
		if org, ok := organisations[result.Entries[i].Organisation]; ok {
			if entry, ok := org.Totals[result.Entries[i].Type]; ok {
				toAdd, _ := strconv.Atoi(result.Entries[i].Amount)
				org.Totals[result.Entries[i].Type] = entry + toAdd
			} else {
				org.Totals[result.Entries[i].Type] = entry
			}
		} else {
			org.Name = result.Entries[i].Organisation
			org.Totals = make(map[string]int)
			toAdd, _ := strconv.Atoi(result.Entries[i].Amount)
			org.Totals[result.Entries[i].Type] = org.Totals[result.Entries[i].Type] + toAdd
			organisations[org.Name] = org
		}
	}

	for org := range organisations {
		fmt.Println(org, ":")
		for s := range organisations[org].Totals {
			fmt.Printf("%v\t\t%v\n", s, organisations[org].Totals[s])
			fmt.Println("---------------------------------")
		}
	}
}
