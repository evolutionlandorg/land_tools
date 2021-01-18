package elem

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
)

var putMap = make(map[Coordinate] [5]int)

var leftNum = [5]int{100000,100000,100000,100000,100000}


func Reservation( n Coordinate){

	elemPut := [5]int{0}
	//存放结果的slice
	nums := GenerateRandSlice(5,3)

	for i := 0; i < len(nums); i++ {
		//对elemPut进行赋值
		r := nums[i]
		if i <= 1 {
			if r != 3{
				elemPut[r] = 100
			}else {
				elemPut[r] = 200
			}

		} else {
			sum := 0
			for i := 0; i < len(elemPut);i++{
				sum = sum + elemPut[i]
			}
			//	 用400减去 前两个的和
			elemPut[r] = 400 - sum
		}

	}

	putMap[n] = elemPut

	calculateLeft(n)

}



func ClosedLand( n Coordinate)  {

	elemPut := closedElementMap[n]

	//存放结果的slice
	nums := GenerateRandSlice(5,2)


 // 距离金木水火土小于等于3的地块，还会随机输出其他元素
	tag := 0
	var k []int //保存输出不是0的下标
	for i := 0; i < len(elemPut); i++{
		if elemPut[i] != 0{
			k = append(k,i)
			tag++
		}
	}

	if tag == 1{
		for i := 0; i < 2; i++{
			// 不输出木元素
			if nums[i] != k[0] && nums[i] != 1{
				elemPut[nums[i]] = 100
			}
		}

	}else if tag == 2{
		for i := 0; i < 1; i++{
			for j := 0; j < len(k); j++{
				if nums[i] != k[j] && nums[i] != 1{
					elemPut[nums[i]] = 100
					break
				}
			}

		}

	}else if tag == 3{
		// 最多输出3种元素，所以tag==3时，什么也不做
	}

	putMap[n] = elemPut
	calculateLeft(n)

}
func OtherLandSect() [5]int {
	elemPut := [5]int{0}
	nums := GenerateRandSlice(5,2)

	//再产生0~2的随机数
	d := rand.New(rand.NewSource(time.Now().UnixNano()))
	r := d.Intn(2)
	num := int(r)

	// fmt.Println(leftNum)
	if num == 0 {
		l := nums[0]
		if l != 1 {
			elemPut[l] = 100
		}else if  leftNum[1]-100 > 0{
			elemPut[1] = 100
		}
	} else if num == 1{
		for i := 0; i < 2; i++{
			r := nums[i]
			if r != 1 {
				elemPut[r] = 100
			}else if leftNum[1]-100 > 0{
				elemPut[1] = 100
			}
		}
	}
	return elemPut
}

func OtherLandSect_() [5]int {
	elemPut := [5]int{0,0,0,0,0}
	d := rand.New(rand.NewSource(time.Now().UnixNano()))
	r := d.Intn(2)
	num := int(r)
	if num == 0{
		elemPut[3] = 100
	}else if num == 1{
		elemPut = [5]int{100,0,0,100,0}
	}

	return elemPut
}

func OtherLand(n Coordinate)  {

	elemPut  := OtherLandSect()
	// 如果出现(0 0 0 0 0)，则重新进行计算。产生(0 0 0 100 0 ) (100 0 0 100 0)两种情况，因为金和火剩余的比较多
	for elemPut == [5]int{0,0,0,0,0}{
		elemPut = OtherLandSect_()
	}
	putMap[n] = elemPut

	calculateLeft(n)

}

func GoldLand(n Coordinate)  {
	putMap[n] = [5]int{400,0,0,0,0}

	calculateLeft(n)
}
func WoodLand(n Coordinate)  {
	putMap[n] = [5]int{0,400,0,0,0}

	calculateLeft(n)
}
func LakeLand(n Coordinate)  {
	putMap[n] = [5]int{0,0,400,0,0}

	calculateLeft(n)
}
func FireLand(n Coordinate)  {
	putMap[n] = [5]int{0,0,0,400,0}

	calculateLeft(n)
}
func EarthLand(n Coordinate)  {
	putMap[n] = [5]int{0,0,0,0,400}

	calculateLeft(n)
}




func Fill(){
	//fmt.Println("----------中心保留地-----------")
	for i := 0; i < len(reserved); i++{
		Reservation(reserved[i])
	}

	/*fmt.Println("----------其余保留地-----------")
	for i := 0; i < len(reservation_other); i++{
		ClosedLand(reservation_other[i])
	}*/
	//ClosedLand(Coordinate{-71,1})

	//fmt.Println("----------金矿山-----------")
	for i := 0; i < len(gold); i++{
		GoldLand(gold[i])
	}

	//fmt.Println("----------森林-----------")
	for i := 0; i < len(wood); i++{
		WoodLand(wood[i])
	}

	//fmt.Println("----------湖泊-----------")
	for i := 0; i < len(water); i++{
		LakeLand(water[i])
	}

	//fmt.Println("----------火山-----------")
	for i := 0; i < len(fire); i++{
		FireLand(fire[i])
	}

	//fmt.Println("----------土山-----------")
	for i := 0; i < len(earth); i++{
		EarthLand(earth[i])
	}


	//fmt.Println("----------有加成的土地-----------")
	for i := 0; i < len(closed); i++{
		ClosedLand(closed[i])
	}

	//fmt.Println("----------其余的土地-----------")

	for i := 0; i < len(other); i++{
		OtherLand(other[i])
	}


	for i := cordRange.minx; i <= cordRange.maxx; i++{
		Addition(Coordinate{i,cordRange.maxy})
	}

	for i := cordRange.minx; i <= cordRange.maxx; i++{
		Addition(Coordinate{i,cordRange.miny})
	}

	for i := cordRange.miny+1; i <= cordRange.maxy-1; i++{
		Addition(Coordinate{cordRange.minx,i})
	}


	for i := cordRange.maxy-1; i >= cordRange.miny+1; i--{
		Addition(Coordinate{cordRange.maxx,i})
	}

}


func calculateLeft(n Coordinate)  {

	elemPut, _ := putMap[n]

	// 计算剩余的元素量
	for i := 0; i < 5; i++{
		leftNum[i] = leftNum[i] - elemPut[i]
	}
	//fmt.Printf("[ %+v=>%+v ]\n",n,elemPut)
}

var dataFilePath = "./data/resource.json"

func SaveFile() {
	var lands = []Resource{}
	for j := cordRange.miny; j <= cordRange.maxy; j++ {
		for i := cordRange.minx; i <= cordRange.maxx; i++ {
			v := putMap[Coordinate{i,j}]
			lands = append(lands,Resource{Gold: v[GOLD],Wood: v[WOOD],Water: v[WATER],Fire: v[FIRE],Earth: v[EARTH],Coordinate:Coordinate{i,j}})
		}
	}
	data, _ := json.MarshalIndent(lands, "", "  ")
	_ = ioutil.WriteFile(dataFilePath, data, 0644)
}

func Addition(n Coordinate)  {

	elemPut := putMap[n]

	d := rand.New(rand.NewSource(time.Now().UnixNano()))
	//存放结果的slice
	nums := GenerateRandSlice(5,3)

	//再产生0~2的随机数
	r := d.Intn(2)
	num := int(r)
	// 原来元素个数要加的数量
	nu := [5]int{150,0,150,150,150}
	if num == 0{
		l := nums[0]

		leftNum[l] = leftNum[l] - nu[l]
		elemPut[l] = elemPut[l] + nu[l]

	}else if num == 1{
		for i := 0; i < 2; i++{
			l := nums[i]

			leftNum[l] = leftNum[l] - nu[l]
			elemPut[l] = elemPut[l] + nu[l]
		}
	}

	putMap[n] = elemPut

}


