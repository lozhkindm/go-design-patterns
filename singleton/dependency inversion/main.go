package main

import (
	"fmt"
)

type DB interface {
	GetPopulation(name string) int
}

type dummyDB struct {
	cities map[string]int
}

func (d *dummyDB) GetPopulation(name string) int {
	if len(d.cities) == 0 {
		d.cities = map[string]int{
			"Moscow": 100,
			"SPB":    250,
		}
	}
	return d.cities[name]
}

func GetTotalPopulation(db DB, cities []string) int {
	result := 0
	for _, city := range cities {
		result += db.GetPopulation(city) // problem solved, depending upon interface
	}
	return result
}

func main() {
	cities := []string{"Moscow", "SPB"}
	total := GetTotalPopulation(&dummyDB{}, cities)
	ok := total == (100 + 250) // this test is ok and will not break
	fmt.Println(ok)
}
