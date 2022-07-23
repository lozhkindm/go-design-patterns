package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"strconv"
	"sync"
)

// sync.Once - thread safety, laziness
// init() - thread safety, no laziness

//go:embed cities.txt
var fs embed.FS

var (
	once     sync.Once
	instance *singletonDB
)

type singletonDB struct {
	cities map[string]int
}

func (s *singletonDB) GetPopulation(name string) int {
	return s.cities[name]
}

func GetSingletonDB() *singletonDB {
	once.Do(func() {
		cities, err := readData("cities.txt")
		if err != nil {
			log.Fatal(err)
		}
		db := singletonDB{cities: cities}
		instance = &db
	})
	return instance
}

func readData(path string) (map[string]int, error) {
	file, err := fs.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, err := strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		}
		result[k] = v
	}

	return result, nil
}

func GetTotalPopulation(cities []string) int {
	result := 0
	for _, city := range cities {
		result += GetSingletonDB().GetPopulation(city) // problem, depending upon implementation
	}
	return result
}

func main() {
	cities := []string{"Moscow", "SPB"}
	total := GetTotalPopulation(cities)
	ok := total == (12655050 + 5377503) // this test will break after the data will be changed
	fmt.Println(ok)
}
