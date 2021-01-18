package main

import (
	"land_tools/elem"
)


func main() {
	elem.LoadResource()
	elem.FindClosed()
	elem.FindOtherLand()
	elem.Fill()
	elem.SaveFile()
}
