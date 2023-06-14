// Copyright (c) 2023 Dominic M. All rights reserved
//
// Created by: Dominic M
// Created on: June 2023
// This program shows you your selected stat of pikachu

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type StatData struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Stat []struct {
		Icon string `json:"icon"`
	} `json:"Stat"`
}

func main() {
	var stat string

	var jsonData StatData

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

	fmt.Println("This program shows you your selected stat of pikachu.")
	fmt.Println()
	fmt.Print("Please enter a from the following: ")
	fmt.Println("Type")
	fmt.Println("Evolution")
	fmt.Println("Health")
	fmt.Println("Attacks")
	fmt.Println("Weaknesses")
	fmt.Println("Resistances")
	fmt.Println()
	fmt.Scanln(&stat)

	// temperature := jsonData.Main.Temp
	// temperatureCelcius := temperature - 273.15
	// temperatureFormatted := fmt.Sprintf("%.2f", temperatureCelcius)

	if stat == "Type" {
		fmt.Println(jsonData.Stat[0].Icon)
	} else if stat == "Evolution" {
		fmt.Println(jsonData.Stat[1].Icon)
	}

	fmt.Println()
	fmt.Print("The temperature outside is ", temperatureFormatted, "°C!")
	fmt.Println()
	fmt.Println("\nDone.")
}
