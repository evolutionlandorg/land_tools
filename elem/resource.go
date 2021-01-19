package elem

import (
	"encoding/json"
	"io/ioutil"
)

var gold,wood, water,fire,earth,reserved []Coordinate

var cordRange Range

var inputFile = "./input.json"

var FormatToJson = func(){
	var goldList,woodList,waterList,fireList,earthList,reservedList [][2]int
	for _,v := range gold{
		goldList = append(goldList,[2]int{v.X,v.Y})
	}
	for _,v := range wood{
		woodList = append(woodList,[2]int{v.X,v.Y})
	}
	for _,v := range water {
		waterList = append(waterList,[2]int{v.X,v.Y})
	}
	for _,v := range fire{
		fireList = append(fireList,[2]int{v.X,v.Y})
	}
	for _,v := range earth{
		earthList = append(earthList,[2]int{v.X,v.Y})
	}
	for _,v := range reserved {
		reservedList = append(reservedList,[2]int{v.X,v.Y})
	}
	var resource = make(map[string]interface{})
	resource["range"] = []int{cordRange.minx,cordRange.maxx,cordRange.miny,cordRange.maxy}
	resource["gold"] = goldList
	resource["wood"] = woodList
	resource["water"] = waterList
	resource["fire"] = fireList
	resource["earth"] = earthList
	resource["reserved"] = reservedList
	data, _ := json.MarshalIndent(resource, "", "  ")
	_ = ioutil.WriteFile(inputFile, data, 0644)
}

func LoadResource(){
	var resource map[string]interface{}

	data, _ := ioutil.ReadFile(inputFile)
	json.Unmarshal(data, &resource)

	rect := resource["range"].([]interface{})
	cordRange.minx = Float2Int(rect[0])
	cordRange.maxx = Float2Int(rect[1])
	cordRange.miny = Float2Int(rect[2])
	cordRange.maxy = Float2Int(rect[3])

	var elementsMap = make(map[string][]Coordinate)
	for key,_:= range AllMap{
		var list []Coordinate
		listRaw := resource[key].([]interface{})
		for _,raw := range listRaw{
			v := raw.([]interface{})
			list = append(list,Coordinate{Float2Int(v[0]), Float2Int(v[1])})
		}
		elementsMap[key] = list
	}

	gold = elementsMap["gold"]
	wood = elementsMap["wood"]
	water = elementsMap["water"]
	fire = elementsMap["fire"]
	earth = elementsMap["earth"]
	reserved = elementsMap["reserved"]
}

func FillResource() []Resource{
	var lands = []Resource{}
	for j := cordRange.miny; j <= cordRange.maxy; j++ {
		for i := cordRange.minx; i <= cordRange.maxx; i++ {
			v := putMap[Coordinate{i,j}]
			lands = append(lands,Resource{Gold: v[GOLD],Wood: v[WOOD],Water: v[WATER],Fire: v[FIRE],Earth: v[EARTH],Coordinate:Coordinate{i,j}})
		}
	}
	return lands
}

var dataFilePath = "./data/resource.json"

func SaveFile() {
	data, _ := json.MarshalIndent(elems, "", "  ")
	_ = ioutil.WriteFile(dataFilePath, data, 0644)
}

