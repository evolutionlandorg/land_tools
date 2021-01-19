package elem

const DistanceThreshhold = 3

var (
	closed           []Coordinate
	other            []Coordinate
	closedElementMap = make(map[Coordinate][5]int)
)

func findClosedLand(n Coordinate) {

	var distancesGold = make([]float64, len(gold))
	var distancesWood = make([]float64, len(wood))
	var distancesLake = make([]float64, len(water))
	var distancesFire = make([]float64, len(fire))
	var distancesEarth = make([]float64, len(earth))

	for i := 0; i < len(wood); i++ {
		distancesWood[i] = CalculateDistance(n, wood[i])
	}

	for i := 0; i < len(gold); i++ {
		distancesGold[i] = CalculateDistance(n, gold[i])
	}

	for i := 0; i < len(water); i++ {
		distancesLake[i] = CalculateDistance(n, water[i])
	}

	for i := 0; i < len(fire); i++ {
		distancesFire[i] = CalculateDistance(n, fire[i])
	}

	for i := 0; i < len(earth); i++ {
		distancesEarth[i] = CalculateDistance(n, earth[i])
	}

	minGold := MinVal(distancesGold)
	minWood := MinVal(distancesWood)
	minLake := MinVal(distancesLake)
	minFire := MinVal(distancesFire)
	minEarth := MinVal(distancesEarth)

	elemPut := [5]int{0}

	weight := func(val float64) int {
		return int(100*1 + 1.0/val)
	}

	if minGold <= DistanceThreshhold {
		elemPut[GOLD] = weight(minGold)
	}
	if minWood <= DistanceThreshhold {
		elemPut[WOOD] = weight(minWood)
	}
	if minLake <= DistanceThreshhold {
		elemPut[WATER] = weight(minLake)
	}
	if minFire <= DistanceThreshhold {
		elemPut[FIRE] = weight(minFire)
	}
	if minEarth <= DistanceThreshhold {
		elemPut[EARTH] = weight(minEarth)
	}

	if minGold <= DistanceThreshhold || minWood <= DistanceThreshhold || minLake <= DistanceThreshhold || minFire <= DistanceThreshhold || minEarth <= DistanceThreshhold {
		closed = append(closed, n)
		closedElementMap[n] = elemPut
	} else {
		other = append(other, n)
	}
}

func FindBarren() {
	resources := MergeSlice(gold, wood, water, fire, earth, reserved)

	for i := cordRange.minx; i <= cordRange.maxx; i++ {
		for j := cordRange.miny; j <= cordRange.maxy; j++ {
			if !IsExist(Coordinate{i, j}, resources) {
				findClosedLand(Coordinate{i, j})
			}
		}
	}
}
