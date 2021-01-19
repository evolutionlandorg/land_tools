package main

import (
	"land_tools/elem"
)

func main() {
	elem.LoadResource()
	elem.FindBarren()
	elem.Fill()
	elem.Random()
	elem.SaveFile()
}
