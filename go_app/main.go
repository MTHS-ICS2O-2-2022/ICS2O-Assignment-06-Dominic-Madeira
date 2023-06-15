// Copyright (c) 2023 Dominic M. All rights reserved
//
// Created by: Dominic M
// Created on: June 2023
// This program is a shows you a certain stat of pikachu.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type StatData struct {
	Data struct {
		Hp string `json:"hp"`
		Subtypes []string `json:"subtypes"`
		Types []string `json:"types"`
		Attacks []struct {
			Name string `json:"name"`
		} `json:"attacks"`
		Weaknesses []struct {
			Type string `json:"type"`
		} `json:"weaknesses"`
		EvolvesTo []string `json:"evolvesTo"`
	} `json:"data"`
}

func main() {

	var jsonData StatData

	var stat int

	fmt.Println("This program is a shows you a certain stat of pikachu.")
	fmt.Println()
	fmt.Print("Please enter a number between 1 and 6 corresponding to the stat: ")
	fmt.Println()
	fmt.Println("1. Attacks")
	fmt.Println("2. Type")
	fmt.Println("3. Weakness")
	fmt.Println("4. Evolves To")
	fmt.Println("5. Health")
	fmt.Println("6. Stage")
	fmt.Scanln(&stat)
	fmt.Println()

	urlAddress := "https://api.pokemontcg.io/v2/cards/basep-1"

	response, err := http.Get(urlAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		fmt.Println(err)
		return
	}

	attack1 := jsonData.Data.Attacks[0].Name
	attack2 := jsonData.Data.Attacks[1].Name
	types := jsonData.Data.Types[0]
	evolvesTo := jsonData.Data.EvolvesTo[0]
	hp := jsonData.Data.Hp
	weakness := jsonData.Data.Weaknesses[0].Type
	stage := jsonData.Data.Subtypes[0]

if stat == 1 {
	fmt.Println("Attacks:", attack1, "and", attack2)
} else if stat == 2 {
	fmt.Println("Type:", types)
} else if stat == 3 {
	fmt.Println("Weakness:", weakness)
} else if stat == 4 {
	fmt.Println("Evolves To:", evolvesTo)
} else if stat == 5 {
	fmt.Println("Health:", hp, "HP")
} else if stat == 6 {
	fmt.Println("Stage:", stage)
} else {
	fmt.Println("Invalid Input")
}

	fmt.Println()
	fmt.Println("\nDone.")
}
