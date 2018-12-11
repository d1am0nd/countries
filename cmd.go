package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"

	cs "countries/countries"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	trimmed := strings.TrimSuffix(string(text), "\n")
	fmt.Println(trimmed)

	countries := cs.Read("./countries.csv")

	country, distance := countries.MatchClosest([]rune(trimmed))
	fmt.Println(country, distance)
}