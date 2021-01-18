package elem

import (
	"math"
	"math/rand"
	"time"
)

func CalculateDistance(from,to Coordinate) (distance float64){
	distance = math.Sqrt(math.Pow(float64(from.X) - float64(to.X),2) + math.Pow(float64(from.Y) - float64(to.Y),2))
	return
}

func MinVal(array []float64) float64 {
	var min = math.MaxFloat64
	if len(array) == 0 {
		return 0
	}
	for _, value := range array {
		if min > value {
			min = value
		}
	}
	return min
}

func IsExist(cord Coordinate,list []Coordinate) (exist bool){
	for n := 0; n < len(list); n++ {
		if  cord.X == list[n].X && cord.Y == list[n].Y {
			exist = true
			break
		}
	}
	return
}

func MergeSlice(resources ...[]Coordinate) (merged []Coordinate){
	for _, r := range resources {
		merged = append(merged, r...)
	}
	return
}

func GenerateRandSlice(randNum,length int) (nums []int){
	d := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		num := d.Intn(randNum)
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
		if len(nums)==length{
			break
		}
	}
	return
}
