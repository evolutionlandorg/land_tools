package elem

var other = []Coordinate{}


func FindOtherLand()  {

	resources := MergeSlice(gold,wood,lake,fire,earth,reservation_center,closed)
	for i := cordRange.minx; i <= cordRange.maxx; i++ {
		for j := cordRange.miny; j <= cordRange.maxy; j++ {
			if !IsExist(Coordinate{i,j},resources) {
				other = append(other, Coordinate{i, j})
			}
		}
	}

}
