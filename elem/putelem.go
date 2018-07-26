package src

import (
	"math/rand"
	"time"
	"fmt"
	"math"

)

var putMap = make(map[coordinate] [5]int)

var leftNum = [5]int{100000,100000,100000,100000,100000}

var uesdNum [5]int


//var addiNum [5]int

var all = [5]int{0,0,0,0,0}


func Reservation( n coordinate){

	elemPut := [5]int{0}
	// 利用时间作为随机数种子
	//rand.Seed(time.Now().Unix())
	d := rand.New(rand.NewSource(time.Now().UnixNano()))
	//存放结果的slice
	nums := make([]int, 0)
	for i := 0; i < 10; i++{
		// r 是生成0~5的随机数
		r := d.Intn(5)
		// num 是对r进行取整
		num := int(r)

		//查重
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
	}

	for i := 0; i < 3; i++ {
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

	PrintMap(n)

}



func ClosedLand( n coordinate)  {

	elemPut := [5]int{0}

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
	if min_gold <= 3.0{
		addition := 1 + 1.0/min_gold
		elemPut[0] = int(100 * addition)
	}

	min_wood := 1000.0
	for i := 0; i < len(distances_wood); i++{
		if min_wood > distances_wood[i] {
			min_wood = distances_wood[i]
		}
	}
	if min_wood <= 3.0{
		addition := 1 + 1.0/min_wood
		elemPut[1] = int(100 * addition)
	}

	min_lake := 1000.0
	for i := 0; i < len(distances_lake); i++{
		if min_lake > distances_lake[i]{
			min_lake = distances_lake[i]
		}
	}
	if min_lake <= 3.0{
		addition := 1 + 1.0/min_lake
		elemPut[2] = int(100 * addition)
	}

	min_fire := 1000.0
	for i := 0; i < len(distances_fire); i++{
		if min_fire > distances_fire[i]{
			min_fire = distances_fire[i]
		}
	}
	if min_fire <= 3.0{
		addition := 1 + 1.0/min_fire
		elemPut[3] = int(100 * addition)
	}


	min_earth := 1000.0
	for i := 0; i < len(distances_earth); i++{
		if min_earth > distances_earth[i]{
			min_earth = distances_earth[i]
		}
	}
	if min_earth <= 3.0{
		addition := 1 + 1.0/min_earth
		elemPut[4] = int(100 * addition)
	}



	// 利用时间作为随机数种子
	//rand.Seed(time.Now().Unix())
	d := rand.New(rand.NewSource(time.Now().UnixNano()))
	//存放结果的slice
	nums := make([]int, 0)
	for i := 0; i < 10; i++{
		// r 是生成0~5的随机数
		r := d.Intn(5)
		// num 是对r进行取整
		num := int(r)

		//查重
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
	}


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
	PrintMap(n)

}
func OtherLandSect() [5]int {
	elemPut := [5]int{0}
	// 利用时间作为随机数种子
	//rand.Seed(time.Now().Unix())
	d := rand.New(rand.NewSource(time.Now().UnixNano()))
	//存放结果的slice
	nums := make([]int, 0)
	for i := 0; i < 10; i++{
		// r 是生成0~5的随机数
		r := d.Intn(5)
		// num 是对r进行取整
		num := int(r)

		//查重
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
	}

	//再产生0~2的随机数
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

func OtherLand(n coordinate)  {

	elemPut  := OtherLandSect()
	// 如果出现(0 0 0 0 0)，则重新进行计算。产生(0 0 0 100 0 ) (100 0 0 100 0)两种情况，因为金和火剩余的比较多
	for elemPut == [5]int{0,0,0,0,0}{
		elemPut = OtherLandSect_()
	}
	putMap[n] = elemPut

	PrintMap(n)

}

func GoldLand(n coordinate)  {
	putMap[n] = [5]int{400,0,0,0,0}

	PrintMap(n)
}
func WoodLand(n coordinate)  {
	putMap[n] = [5]int{0,400,0,0,0}

	PrintMap(n)
}
func LakeLand(n coordinate)  {
	putMap[n] = [5]int{0,0,400,0,0}

	PrintMap(n)
}
func FireLand(n coordinate)  {
	putMap[n] = [5]int{0,0,0,400,0}

	PrintMap(n)
}
func EarthLand(n coordinate)  {
	putMap[n] = [5]int{0,0,0,0,400}

	PrintMap(n)
}




func	Traverse(){
	//fmt.Println("----------中心保留地-----------")
	for i := 0; i < len(reservation_center); i++{
		Reservation(reservation_center[i])
	}

	/*fmt.Println("----------其余保留地-----------")
	for i := 0; i < len(reservation_other); i++{
		ClosedLand(reservation_other[i])
	}*/
	//ClosedLand(coordinate{-71,1})

	//fmt.Println("----------金矿山-----------")
	for i := 0; i < len(gold); i++{
		GoldLand(gold[i])
	}

	//fmt.Println("----------森林-----------")
	for i := 0; i < len(wood); i++{
		WoodLand(wood[i])
	}

	//fmt.Println("----------湖泊-----------")
	for i := 0; i < len(lake); i++{
		LakeLand(lake[i])
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



	for i := -112; i <= -68; i++{
		Addition(coordinate{i,22})
	}


	for i := 21; i >= -22; i--{
		Addition(coordinate{-68,i})
	}

	for i := -112; i <= -67; i++{
		Addition(coordinate{i,-22})
	}

	for i := -21; i <= 21; i++{
		Addition(coordinate{-112,i})
	}

}


func PrintMap(n coordinate)  {

	elemPut, _ := putMap[n]

	// 计算剩余的元素量
	for i := 0; i < 5; i++{
		leftNum[i] = leftNum[i] - elemPut[i]
	}

	//fmt.Printf("[ %+v=>%+v ]\n",n,elemPut)

}



func PrintAll()  {

	for i := 22 ; i >= -22; i--{
		for j := -112; j <= -68; j++{
			elemPut, _ := putMap[coordinate{j,i}]
			fmt.Printf("(")
			for n := 0; n < len(elemPut); n++{
				if elemPut[n] == 0{
					fmt.Printf("  %d ",elemPut[n])
				}else {
					fmt.Printf("%d ", elemPut[n])
				}
			}
			fmt.Printf(")  ")
			//fmt.Printf("%+v, ",elemPut)
			/*if elemPut == [5]int{0,0,0,0,0}{
				fmt.Println(coordinate{j,i})
			}*/
		}
		fmt.Println()
	}


}

func Addition(n coordinate)  {

	elemPut := putMap[n]

	d := rand.New(rand.NewSource(time.Now().UnixNano()))
	//存放结果的slice
	nums := make([]int, 0)
	for i := 0; i < 10; i++{
		// r 是生成0~5的随机数
		r := d.Intn(5)
		// num 是对r进行取整
		num := int(r)

		//查重
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
	}

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

func FindNum()  {

	for i := 0; i < 5; i++{
		uesdNum[i] = 100000-leftNum[i]
		fmt.Println(uesdNum[i])
	}
	fmt.Println("----------")
	for i := 0; i < 5; i++{
		fmt.Println(leftNum[i])
	}
}

// 把总的元素加起来，打印出来
func SumAll()  {
	for i := -112 ; i <= -68; i++{
		for j := -22; j <= 22; j++{
			elemPut, _ := putMap[coordinate{i,j}]
			for n := 0; n < 5; n++{
				all[n] = all[n] + elemPut[n]
			}
		}

	}

	fmt.Println(all)

}

