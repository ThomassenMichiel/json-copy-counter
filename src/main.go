package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
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
	jsonFile, err := os.Open("Kopies.json")

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
			toAdd, _ := strconv.Atoi(result.Entries[i].Amount)
			org.Totals[result.Entries[i].Type] = org.Totals[result.Entries[i].Type] + toAdd
		} else {
			org.Name = result.Entries[i].Organisation
			org.Totals = make(map[string]int)
			toAdd, _ := strconv.Atoi(result.Entries[i].Amount)
			org.Totals[result.Entries[i].Type] = org.Totals[result.Entries[i].Type] + toAdd
			organisations[org.Name] = org
		}
	}

	longestKey := 0
	keys := make([]string, 0)
	for org := range organisations {
		keys = append(keys, org)
		for s := range organisations[org].Totals {
			if len(s) > longestKey {
				longestKey = len(s)
			}
		}
	}
	sort.Strings(keys)

	for org := range keys {
		totals := organisations[keys[org]].Totals
		totalsKeys := make([]string, 0)
		for s := range totals {
			totalsKeys = append(totalsKeys, s)
		}
		sort.Strings(totalsKeys)
		fmt.Println(keys[org], ":")
		for s := range totalsKeys {
			fmt.Printf("%*v\t\t%v\n", -longestKey, totalsKeys[s], totals[totalsKeys[s]])
		}
		fmt.Println("---------------------------------")
	}
}
