package main

import (
	"fmt"
	"golang_towerofhanoi/TOHPackage"
)

func main() {
	var disc int
	fmt.Print("Enter Disc : ")
	fmt.Scan(&disc)
	TOHPackage.TOH(disc, 1, 2, 3)
}
