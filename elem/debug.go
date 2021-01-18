package elem

import "fmt"

var usedNum [5]int


//var addiNum [5]int

var all = [5]int{0,0,0,0,0}

func FindNum()  {

	for i := 0; i < 5; i++{
		usedNum[i] = 100000-leftNum[i]
		fmt.Println(usedNum[i])
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
			elemPut, _ := putMap[Coordinate{i,j}]
			for n := 0; n < 5; n++{
				all[n] = all[n] + elemPut[n]
			}
		}

	}

	fmt.Println(all)

}

func PrintAll()  {

	for i := cordRange.maxy ; i >= cordRange.miny; i--{
		for j := cordRange.minx; j <= cordRange.maxx; j++{
			elemPut, _ := putMap[Coordinate{j,i}]
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
				fmt.Println(Coordinate{j,i})
			}*/
		}
		fmt.Println()
	}


}
