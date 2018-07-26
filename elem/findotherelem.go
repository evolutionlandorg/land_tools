package src

var other = []coordinate{}


func MergeSliceSeven(s1 []coordinate, s2 []coordinate, s3 []coordinate,s4 []coordinate,s5 []coordinate,s6 []coordinate, s7 []coordinate ) []coordinate {
	slice := make([]coordinate, len(s1)+len(s2)+len(s3)+len(s4)+len(s5)+len(s6)+len(s7))
	copy(slice, s1)
	copy(slice[len(s1):], s2)
	copy(slice[len(s1)+len(s2):], s3)
	copy(slice[len(s1)+len(s2)+len(s3):], s4)
	copy(slice[len(s1)+len(s2)+len(s3)+len(s4):], s5)
	copy(slice[len(s1)+len(s2)+len(s3)+len(s4)+len(s5):], s6)
	copy(slice[len(s1)+len(s2)+len(s3)+len(s4)+len(s5)+len(s6):], s7)

	return slice
}


func FindOtherLand()  {

	// 寻找普通地块的坐标，并且添加到other切片中
	all_land := MergeSliceSeven(gold,wood,lake,fire,earth,reservation_center,closed)

	tag := false

	for i := 22; i >= -22; i--{
		for j := -112; j <= -68; j++ {

			for n :=0 ; n < len(all_land); n++{


				if (j != all_land[n].x) || (i != all_land[n].y){

					tag = true
				}else {
					tag = false
					break
				}

			}
			if tag {
				other = append(other, coordinate{j, i})
			}

		}
	}

	// fmt.Println(len(all_land))
	// fmt.Println(len(other))


}
