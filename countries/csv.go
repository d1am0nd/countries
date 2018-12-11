package countries

import (
	"os"
	"fmt"
	"math"
	"unicode"
	"encoding/csv"	
)

type Country struct {
	Name string
	Alpha2 string
	Alpha3 string
	Iso3166 string
	Region string
	SubRegion string
	IntermediateRegion string
	RegionCode string
	SubRegionCode string
	IntermediateRegionCode string
}

type Countries struct {
	Countries []Country
}

func Read(fn string) Countries {
	f, err := os.Open(fn)
	if err != nil {
		fmt.Println("---OS problem---")
		panic(err)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Println("---Reader problem---")
		panic(err)
	}

	countries := make([]Country, len(lines))

	for i, line := range lines {
		countries[i] = Country{
			line[0],
			line[1],
			line[2],
			line[3],
			line[4],
			line[5],
			line[6],
			line[7],
			line[8],
			line[9],
		}
	}

	return Countries{countries}
}

func (c Countries) MatchClosest(name []rune) (match Country, distance int) {
	distance = math.MaxInt32

	toLower(name)

	for _, country := range c.Countries {
		cName := []rune(country.Name)
		toLower(cName)
		d := Lev(name, cName)
		if d < distance {
			distance = d
			match = country
		}
	}

	return match, distance
}

func toLower(runes []rune) {
	for i := range runes {
		runes[i] = unicode.ToLower(runes[i])
	}
}