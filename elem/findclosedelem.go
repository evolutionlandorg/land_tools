package src





import (
		"math"
)

type coordinate struct {
	x int
	y int
}


var gold = []coordinate{coordinate{-108,-4},coordinate{-107,-3},coordinate{-107,-4},
	coordinate{-107,-5},coordinate{-106,-4},coordinate{-106,-5},
	coordinate{-106,-6},coordinate{-105,-5},coordinate{-105,-6},
	coordinate{-105,-7},coordinate{-104,-4},coordinate{-104,-5},
	coordinate{-104,-6},coordinate{-75,2},coordinate{-75,3},
	coordinate{-74,1},coordinate{-74,2},coordinate{-74,3},coordinate{-74,4},
	coordinate{-73,1},coordinate{-73,2},coordinate{-73,3},coordinate{-72,2},
}
var wood = []coordinate{coordinate{-92,-19},coordinate{-92,-18},coordinate{-93,-18},
	coordinate{-94,-18},coordinate{-94,-17},coordinate{-95,-17},coordinate{-75,-19},
	coordinate{-74,-19},coordinate{-74,-18},coordinate{-74,-17},coordinate{-73,-18},
	coordinate{-73,-17},coordinate{-73,-16},coordinate{-72,-17},coordinate{-72,-16},
	coordinate{-72,-15},coordinate{-77,17},coordinate{-78,17},coordinate{-78,18},
	coordinate{-79,17},coordinate{-79,18},coordinate{-80,17},coordinate{-74,-8},
	coordinate{-75,-7},coordinate{-74,-7},coordinate{-73,-7},coordinate{-74,-6},
	coordinate{-73,-6},coordinate{-73,-4},coordinate{-74,-3},coordinate{-73,-3},
	coordinate{-72,-3},coordinate{-74,-2},coordinate{-73,-2},coordinate{-72,-2},
	coordinate{-73,-1},coordinate{-72,-1},coordinate{-71,-1},coordinate{-73,0},
	coordinate{-72,0},coordinate{-71,0},coordinate{-72,1},coordinate{-71,2},coordinate{-102,16},coordinate{-101,16},
	coordinate{-101,17},coordinate{-100,17},coordinate{-99,17},coordinate{-98,17},
	coordinate{-98,18},coordinate{-97,17},coordinate{-97,18},coordinate{-97,19},
	coordinate{-96,18},coordinate{-95,18},coordinate{-105,-18},coordinate{-107,-15},
	coordinate{-106,-16},coordinate{-106,-15},coordinate{-106,-14},coordinate{-105,-16},
	coordinate{-105,-15},coordinate{-105,-14},coordinate{-104,-16},coordinate{-104,-15},
	coordinate{-104,-14},coordinate{-103,-18},coordinate{-103,-17},coordinate{-103,-16},
	coordinate{-103,-15},coordinate{-103,-14},coordinate{-103,-13},coordinate{-102,-18},
	coordinate{-102,-17},coordinate{-102,-16},coordinate{-102,-15},coordinate{-102,-14},
	coordinate{-102,-13},coordinate{-101,-17},coordinate{-101,-16},coordinate{-101,-15},
	coordinate{-101,-14},coordinate{-100,-16},coordinate{-100,-15},coordinate{-100,-14},}

var lake = []coordinate{coordinate{-100,7},coordinate{-100,8},coordinate{-99,6},coordinate{-99,7},
	coordinate{-99,8},coordinate{-99,9},coordinate{-98,7},coordinate{-98,8},
	coordinate{-98,9},coordinate{-75,-6},coordinate{-75,-5},coordinate{-74,-5},
	coordinate{-74,-4},coordinate{-75,17},coordinate{-75,18},coordinate{-76,17},
	coordinate{-76,18},coordinate{-76,19},coordinate{-77,18},coordinate{-77,13},
	coordinate{-78,13},coordinate{-78,14},coordinate{-79,13},coordinate{-79,14},
	coordinate{-79,15},coordinate{-80,14},coordinate{-80,15},coordinate{-80,16},
	coordinate{-81,15},coordinate{-81,16},}

var fire = []coordinate{coordinate{-87,-16},coordinate{-86,-17},coordinate{-86,-16},coordinate{-86,-15},
	coordinate{-85,-17},coordinate{-85,-16},coordinate{-85,-15},coordinate{-85,-14},
	coordinate{-84,-16},coordinate{-84,-15},coordinate{-84,-14},coordinate{-83,-15},
	}

var earth = []coordinate{coordinate{-75,15},coordinate{-75,16},coordinate{-76,14},coordinate{-76,15},
	coordinate{-76,16},coordinate{-77,14},coordinate{-77,15},coordinate{-77,16},
	coordinate{-78,15},coordinate{-78,16},coordinate{-79,16},coordinate{-108,14},
	coordinate{-108,15},coordinate{-107,15},coordinate{-107,16},coordinate{-106,16},
	coordinate{-106,17},coordinate{-105,16},coordinate{-105,17},coordinate{-104,17},
	coordinate{-104,18},coordinate{-103,17},coordinate{-103,18},coordinate{-102,17},
	coordinate{-102,18},coordinate{-101,18},coordinate{-101,19},coordinate{-100,18},
	coordinate{-100,19},coordinate{-99,18},coordinate{-99,19},coordinate{-98,19},}


var reservation_center = []coordinate{coordinate{-86,-4},coordinate{-87,-4},coordinate{-88,-4},coordinate{-89,-4},coordinate{-90,-4},coordinate{-91,-4},coordinate{-92,-4},coordinate{-93,-4},coordinate{-94,-4},
	coordinate{-86,-3},coordinate{-87,-3},coordinate{-88,-3},coordinate{-89,-3},coordinate{-90,-3},coordinate{-91,-3},coordinate{-92,-3},coordinate{-93,-3},coordinate{-94,-3},
	coordinate{-86,-2},coordinate{-87,-2},coordinate{-88,-2},coordinate{-89,-2},coordinate{-90,-2},coordinate{-91,-2},coordinate{-92,-2},coordinate{-93,-2},coordinate{-94,-2},
	coordinate{-86,-1},coordinate{-87,-1},coordinate{-88,-1},coordinate{-89,-1},coordinate{-90,-1},coordinate{-91,-1},coordinate{-92,-1},coordinate{-93,-1},coordinate{-94,-1},
	coordinate{-86,0},coordinate{-87,0},coordinate{-88,0},coordinate{-89,0},coordinate{-90,0},coordinate{-91,0},coordinate{-92,0},coordinate{-93,0},coordinate{-94,0},
	coordinate{-86,1},coordinate{-87,1},coordinate{-88,1},coordinate{-89,1},coordinate{-90,1},coordinate{-91,1},coordinate{-92,1},coordinate{-93,1},coordinate{-94,1},
	coordinate{-86,2},coordinate{-87,2},coordinate{-88,2},coordinate{-89,2},coordinate{-90,2},coordinate{-91,2},coordinate{-92,2},coordinate{-93,2},coordinate{-94,2},
	coordinate{-86,3},coordinate{-87,3},coordinate{-88,3},coordinate{-89,3},coordinate{-90,3},coordinate{-91,3},coordinate{-92,3},coordinate{-93,3},coordinate{-94,3},
	coordinate{-86,4},coordinate{-87,4},coordinate{-88,4},coordinate{-89,4},coordinate{-90,4},coordinate{-91,4},coordinate{-92,4},coordinate{-93,4},coordinate{-94,4},
}


var closed = []coordinate{}


func FindClosedLand( n coordinate)  {



	var distances_gold  [23]float64
	var distances_wood  [85]float64
	var distances_lake  [30]float64
	var distances_fire  [12]float64
	var distances_earth [32]float64

	for i :=0; i < len(wood); i++{
		// 计算两点之间的距离   距离木
		x := float64(n.x) - float64(wood[i].x)
		y := float64(n.y) - float64(wood[i].y)
		x_o := math.Pow(x,2)
		y_o := math.Pow(y,2)
		r := math.Sqrt(x_o + y_o)

		distances_wood[i] = r
	}

	for i :=0; i < len(gold); i++{
		// 计算两点之间的距离   距离金
		x := float64(n.x) - float64(gold[i].x)
		y := float64(n.y) - float64(gold[i].y)
		x_o := math.Pow(x,2)
		y_o := math.Pow(y,2)
		r := math.Sqrt(x_o + y_o)

		distances_gold[i] = r
	}

	for i :=0; i < len(lake); i++{
		// 计算两点之间的距离      距离水
		x := float64(n.x) - float64(lake[i].x)
		y := float64(n.y) - float64(lake[i].y)
		x_o := math.Pow(x,2)
		y_o := math.Pow(y,2)
		r := math.Sqrt(x_o + y_o)

		distances_lake[i] = r
	}


	for i :=0; i < len(fire); i++{
		// 计算两点之间的距离        距离火
		x := float64(n.x) - float64(fire[i].x)
		y := float64(n.y) - float64(fire[i].y)
		x_o := math.Pow(x,2)
		y_o := math.Pow(y,2)
		r := math.Sqrt(x_o + y_o)

		distances_fire[i] = r
	}


	for i :=0; i < len(earth); i++{
		// 计算两点之间的距离     距离土    j
		x := float64(n.x) - float64(earth[i].x)
		y := float64(n.y) - float64(earth[i].y)
		x_o := math.Pow(x,2)
		y_o := math.Pow(y,2)
		r := math.Sqrt(x_o + y_o)

		distances_earth[i] = r
	}

	min_gold := 1000.0
	for i := 0; i < len(distances_gold); i++{
		if min_gold > distances_gold[i] {
			min_gold = distances_gold[i]
		}
	}


	min_wood := 1000.0
	for i := 0; i < len(distances_wood); i++{
		if min_wood > distances_wood[i] {
			min_wood = distances_wood[i]
		}
	}


	min_lake := 1000.0
	for i := 0; i < len(distances_lake); i++{
		if min_lake > distances_lake[i]{
			min_lake = distances_lake[i]
		}
	}


	min_fire := 1000.0
	for i := 0; i < len(distances_fire); i++ {
		if min_fire > distances_fire[i] {
			min_fire = distances_fire[i]
		}

	}
	min_earth := 1000.0
	for i := 0; i < len(distances_earth); i++{
		if min_earth > distances_earth[i]{
			min_earth = distances_earth[i]
		}
	}


	if min_gold <= 3.0 || min_wood <= 3.0 || min_lake <= 3.0 || min_fire <= 3.0 || min_earth <= 3.0{
		closed = append(closed,n)
	}

}

func MergeSliceSix(s1 []coordinate, s2 []coordinate, s3 []coordinate,s4 []coordinate,s5 []coordinate,s6 []coordinate ) []coordinate {
	slice := make([]coordinate, len(s1)+len(s2)+len(s3)+len(s4)+len(s5)+len(s6))
	copy(slice, s1)
	copy(slice[len(s1):], s2)
	copy(slice[len(s1)+len(s2):], s3)
	copy(slice[len(s1)+len(s2)+len(s3):], s4)
	copy(slice[len(s1)+len(s2)+len(s3)+len(s4):], s5)
	copy(slice[len(s1)+len(s2)+len(s3)+len(s4)+len(s5):], s6)

	return slice
}




func FindClosed()  {
	all_land := MergeSliceSix(gold,wood,lake,fire,earth,reservation_center)

	tag := false
	for i := -112; i <= -68; i++ {
		for j := -22; j <= 22; j++ {

			// 如果不和特殊地块相同，则才带入方程
			for n := 0; n < len(all_land); n++ {
				if  i != all_land[n].x || j != all_land[n].y{
					tag = true

				}else {
					tag = false
					break
				}

			}
			if tag {
				FindClosedLand(coordinate{i,j})
			}

		}
	}


}
