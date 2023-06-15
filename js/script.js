// Copyright (c) 2023 Dominic M. All rights reserved
//
// Created by: Dominic M.
// Created on: June 2023
// This file contains the JS functions for index.html

"use strict"

if (navigator.serviceWorker) {
  navigator.serviceWorker.register("/ICS2O-Assignment-6-Dominic-Madeira/sw.js", {
    scope: "/ICS2O-Assignment-6-Dominic-Madeira/sw.js",
  })
}

/**
 * Get API info.
 */
// code from: https://www.youtube.com/watch?v=670f71LTWpM

function myButtonClicked() {

const getStat = async (URLAddress) => {
  try {
    console.log(URLAddress)
    const stat = document.getElementById("user-input").value
    console.log(stat)
    const result = await fetch(URLAddress)
    const jsonData = await result.json()
    const pokemonHP = jsonData.data.hp
    const pokemonStage = jsonData.data.subtypes[0]
    const pokemonType = jsonData.data.types
    const pokemonAttack1 = jsonData.data.attacks[0].name
    const pokemonAttack2 = jsonData.data.attacks[1].name
    const pokemonWeakness = jsonData.data.weaknesses[0].type
    const pokemonEvolution = jsonData.data.evolvesTo[0]

    if (stat == "Health") {
      document.getElementById("answer").innerHTML = pokemonHP + " HP"
    } else if (stat == "Stage") {
      document.getElementById("answer").innerHTML = pokemonStage
    } else if (stat == "Type") {
      document.getElementById("answer").innerHTML = pokemonType
    } else if (stat == "Attacks") {
      document.getElementById("answer").innerHTML = pokemonAttack1 + " and " + pokemonAttack2
    } else if (stat == "Weakness") {
      document.getElementById("answer").innerHTML = pokemonWeakness
    } else if (stat == "Evolution") {
      document.getElementById("answer").innerHTML = pokemonEvolution
    } else {
      document.getElementById("answer").innerHTML = "Invalid Input"
    }
  } catch (err) {
    console.log(err)
  }
}

getStat(
  "https://api.pokemontcg.io/v2/cards/basep-1"
)
}
