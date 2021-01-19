package elem

import (
	"encoding/json"
	"io/ioutil"
)

var gold, wood, water, fire, earth, reserved []Coordinate // full resource land

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
}
