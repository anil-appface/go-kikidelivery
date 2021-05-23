package main

import (
	"fmt"
	"os"
)

//TODO:: To use buffer reader for input reading. bufio.NewReader(os.Stdin)
func main() {

	problem := os.Args[1]

	if problem == "1" {
		// problem1 -> calulates delivery cost
		allPackages := findDeliveryCostForPackages()

		//print output format
		for _, v := range allPackages {
			fmt.Printf("%s %d %d", v.Id, v.DiscountCost, v.TotalCost)
			fmt.Println()
		}
	}

	//Problem 2 -> calculates delivery estimation time
	if problem == "2" {
		packages := calculateDeliveryTimeEstimation()

		//print output format
		for _, v := range packages {
			fmt.Printf("%s %d %d %f", v.Id, v.DiscountCost, v.TotalCost, v.DeliveryTime)
			fmt.Println()
		}

	}

}

// 100 5
// PKG1 50 30 OFR001
// PKG2 75 125 OFFR0008
// PKG3 175 100 OFFR003
// PKG4 110 60 OFFR002
// PKG5 155 95 NA
// 2 70 200
