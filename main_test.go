package main

import (
	"fmt"
	"sort"
	"testing"
)

// PKG1 50 30 OFR001
// PKG2 75 125 OFR0008
// PKG3 175 100 OFR003
// PKG4 110 60 OFR002
// PKG5 155 95 NA

func getAllPackages_test() Packages {

	allPackages := make(Packages, 0)
	allPackages = append(allPackages, NewPackage("PKG1", 100, 30, 50, allDiscounts().getDiscountByCoupon("OFR001")))
	allPackages = append(allPackages, NewPackage("PKG2", 100, 125, 75, allDiscounts().getDiscountByCoupon("OFR008")))
	allPackages = append(allPackages, NewPackage("PKG3", 100, 100, 175, allDiscounts().getDiscountByCoupon("OFR003")))
	allPackages = append(allPackages, NewPackage("PKG4", 100, 60, 110, allDiscounts().getDiscountByCoupon("OFR002")))
	allPackages = append(allPackages, NewPackage("PKG5", 100, 95, 155, allDiscounts().getDiscountByCoupon("NA")))

	return allPackages
}

//Test case to debug problem2

func TestGetPackagesForVehicle(t *testing.T) {
	allPackages := getAllPackages_test()
	sort.Sort(allPackages)

	p := calculatePackagesGroups(200, allPackages)
	sort.Sort(p)
	vehicles := make(Vehicles, 0)
	vehicles = append(vehicles, NewVehicle(1, 70, 200))
	vehicles = append(vehicles, NewVehicle(2, 70, 200))

	index := 0
	for {
		if len(p) == index {
			break
		}
		for _, v := range vehicles {
			fmt.Println(v.id, v.nextDeliveryTime)
			p[index].SetDeliveryTimeForPackages(v.maxSpeed, v.nextDeliveryTime)
			v.SetNextDeliveryTime(p[index].GetTotalDeliveryTime())
			index++
		}
		sort.Sort(vehicles)
	}

	for _, groups := range p {
		for _, v := range groups {
			fmt.Println(v.id, v.deliveryTime)
		}
	}

}
