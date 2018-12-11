package main

import (
	"fmt"
	cs "countries/countries"
)

func main() {
	fmt.Println("hello world")
	fmt.Println("Pretty cool isn't it")

	countries := cs.Read("./countries.csv")

	fmt.Println(countries)
	fmt.Println(cs.Lev([]rune("abrakadabr"), []rune("asbarakak")))
}
