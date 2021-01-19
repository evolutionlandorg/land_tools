package elem

import (
	"math/rand"
	"time"
)

var elems []Resource

func Random() {
	d := rand.New(rand.NewSource(time.Now().UnixNano()))
	elems = FillResource()
	for i := 0; i < len(elems); i++ {
		r := proRandOne()
		sum := 0
		if elems[i].Gold == 0 {
			elems[i].Gold = r[0]
			sum = sum + r[0]
		}

		if elems[i].Wood == 0 {
			elems[i].Wood = d.Intn(3)
			sum = sum + elems[i].Wood
		}

		if elems[i].Water == 0 {
			elems[i].Water = r[2]
			sum = sum + r[2]
		}

		if elems[i].Fire == 0 {
			elems[i].Fire = r[3]
			sum = sum + r[3]
		}

		if elems[i].Earth == 0 {
			elems[i].Earth = r[4]
			sum = sum + r[4]
		}

		sub := [5]int{-1, -2, -3, -4, -5}
		num := 0
		if elems[i].Gold == 100 || elems[i].Gold == 200 || elems[i].Gold == 400 {
			sub[0] = 0
			num++
		}
		if elems[i].Wood == 100 || elems[i].Wood == 200 || elems[i].Wood == 400 {
			sub[1] = 1
			num++
		}
		if elems[i].Water == 100 || elems[i].Water == 200 || elems[i].Water == 400 {
			sub[2] = 2
			num++
		}
		if elems[i].Fire == 100 || elems[i].Fire == 200 || elems[i].Fire == 400 {
			sub[3] = 3
			num++
		}
		if elems[i].Earth == 100 || elems[i].Earth == 200 || elems[i].Earth == 400 {
			sub[4] = 4
			num++
		}

		eachNum := 0
		if num != 0 {
			eachNum = sum / num

		}

		if elems[i].Gold == 100 || elems[i].Gold == 200 || elems[i].Gold == 400 {
			elems[i].Gold = elems[i].Gold - eachNum - d.Intn(5)

		}

		if elems[i].Wood == 100 || elems[i].Wood == 200 || elems[i].Wood == 400 {
			elems[i].Wood = elems[i].Wood - eachNum - d.Intn(2)

		}

		if elems[i].Water == 100 || elems[i].Water == 200 || elems[i].Water == 400 {
			elems[i].Water = elems[i].Water - eachNum - d.Intn(2)

		}

		if elems[i].Fire == 100 || elems[i].Fire == 200 || elems[i].Fire == 400 {
			elems[i].Fire = elems[i].Fire - eachNum - d.Intn(6)

		}

		if elems[i].Earth == 100 || elems[i].Earth == 200 || elems[i].Earth == 400 {
			elems[i].Earth = elems[i].Earth - eachNum - d.Intn(5)

		}

	}

	n := d.Intn(2) + 1
	for i := 0; i < len(elems); i++ {
		if elems[i].Wood == 0 {
			elems[i].Wood = n
		}
		if elems[i].Wood > 10 {
			elems[i].Wood = elems[i].Wood - n
		}
	}

}

func proRandOne() []int {
	d := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 存放结果的slice
	nums := make([]int, 0)
	for i := 0; i < 10; i++ {
		// r 是生成0~10的随机数
		r := d.Intn(5)
		// num 是对r进行取整 , 产生1，2。。。。10
		num := int(r) + 1

		nums = append(nums, num)
	}
	return nums
}
