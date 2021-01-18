package elem

const DistanceThreshhold = 3

var closed []Coordinate

var closedElementMap = make(map[Coordinate] [5]int)


func findClosedLand(n Coordinate) {

	var distances_gold  = make([]float64,len(gold))
	var distances_wood  = make([]float64,len(wood))
	var distances_lake  = make([]float64,len(water))
	var distances_fire  = make([]float64,len(fire))
	var distances_earth = make([]float64,len(earth))

	for i :=0; i < len(wood); i++{
		distances_wood[i] = CalculateDistance(n,wood[i])
	}

	for i :=0; i < len(gold); i++{
		distances_gold[i] = CalculateDistance(n,gold[i])
	}

	for i :=0; i < len(water); i++{
		distances_lake[i] = CalculateDistance(n, water[i])
	}

	for i :=0; i < len(fire); i++{
		distances_fire[i] = CalculateDistance(n,fire[i])
	}

	for i :=0; i < len(earth); i++{
		distances_earth[i] = CalculateDistance(n,earth[i])
	}

	min_gold := MinVal(distances_gold)
	min_wood := MinVal(distances_wood)
	min_lake := MinVal(distances_lake)
	min_fire := MinVal(distances_fire)
	min_earth := MinVal(distances_earth)

	elemPut := [5]int{0}

	weight := func(val float64) int{
		return int(100 * 1 + 1.0/val)
	}

	if min_gold <= DistanceThreshhold{
		elemPut[GOLD] = weight(min_gold)
	}
	if min_wood <= DistanceThreshhold{
		elemPut[WOOD] = weight(min_wood)
	}
	if min_lake <= DistanceThreshhold{
		elemPut[WATER] = weight(min_lake)
	}
	if min_fire <= DistanceThreshhold{
		elemPut[FIRE] = weight(min_fire)
	}
	if min_earth <= DistanceThreshhold{
		elemPut[EARTH] = weight(min_earth)
	}

	if min_gold <= DistanceThreshhold || min_wood <= DistanceThreshhold || min_lake <= DistanceThreshhold || min_fire <= DistanceThreshhold || min_earth <= DistanceThreshhold{
		closed = append(closed,n)
		closedElementMap[n] = elemPut
	}
}


func FindClosed()  {
	resources := MergeSlice(gold,wood, water,fire,earth, reserved)
	for i := cordRange.minx; i <= cordRange.maxx; i++ {
		for j := cordRange.miny; j <= cordRange.maxy; j++ {
			if !IsExist(Coordinate{i,j},resources) {
				findClosedLand(Coordinate{i,j})
			}
		}
	}
}
