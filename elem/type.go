package elem

const (
	GOLD = iota
	WOOD
	WATER
	FIRE
	EARTH
	RESERVED
)

var ElementsMap = map[string]int {"gold":GOLD,"wood":WOOD,"water":WATER,"fire":FIRE,"earth":EARTH}
var AllMap = map[string]int {"gold":GOLD,"wood":WOOD,"water":WATER,"fire":FIRE,"earth":EARTH,"reserved":RESERVED}


type Range struct {
	minx, maxx,miny,maxy int
}

type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Resource struct {
	Gold int `json:"gold"`
	Wood int `json:"wood"`
	Water int `json:"water"`
	Fire int `json:"fire"`
	Earth int `json:"earth"`
	Coordinate Coordinate `json:"coordinate"`
}