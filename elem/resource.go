package elem

import (
	"encoding/json"
	"io/ioutil"
)

var gold, wood, water, fire, earth, reserved []Coordinate // full resource land

var locationMap map[string]int

var cordRange Range

var inputFile = "./input.json"

type LandResource struct {
	gold      uint
	wood      uint
	water     uint
	fire      uint
	earth     uint
	isSpecial uint
}

type InputResource struct {
	Range    []int
	Earth    [][]int
	Fire     [][]int
	Gold     [][]int
	Reserved [][]int
	Water    [][]int
	Wood     [][]int
	Location map[string]int
}

func LoadResource() {
	var resource InputResource

	data, _ := ioutil.ReadFile(inputFile)
	_ = json.Unmarshal(data, &resource)

	rect := resource.Range
	cordRange.minx = rect[0]
	cordRange.maxx = rect[1]
	cordRange.miny = rect[2]
	cordRange.maxy = rect[3]

	var elementDeal = func(elements [][]int) []Coordinate {
		var list []Coordinate
		for _, raw := range elements {
			list = append(list, Coordinate{raw[0], raw[1]})
		}
		return list
	}

	gold = elementDeal(resource.Gold)
	wood = elementDeal(resource.Wood)
	water = elementDeal(resource.Water)
	fire = elementDeal(resource.Fire)
	earth = elementDeal(resource.Earth)
	reserved = elementDeal(resource.Reserved)
	locationMap = resource.Location
}

var dataFilePath = "land_resource.json"

func SaveFile() {
	// var singleLand [][]int
	// for _, elem := range elems {
	// 	singleLand = append(singleLand, []int{elem.Gold, elem.Wood, elem.Water, elem.Fire, elem.Earth})
	// }
	// b, _ := json.Marshal(singleLand)
	// fmt.Println(string(b))
	data, _ := json.MarshalIndent(elems, "", "  ")
	_ = ioutil.WriteFile(dataFilePath, data, 0644)
}
