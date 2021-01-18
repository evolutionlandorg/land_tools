package elem

const (
	GOLD = iota
	WOOD
	WATER
	FIRE
	EARTH
)

var Elements = []int{GOLD,WOOD,WATER,FIRE,EARTH}

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