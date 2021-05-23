package main

import (
	"fmt"
	"sort"
	"testing"

	"github.com/anil-appface/go-kikidelivery/delivery"
)

// PKG1 50 30 OFR001
// PKG2 75 125 OFR0008
// PKG3 175 100 OFR003
// PKG4 110 60 OFR002
// PKG5 155 95 NA

func getAllPackages_test1() delivery.Packages {

	allPackages := make(delivery.Packages, 0)
	allPackages = append(allPackages, delivery.NewPackage("PKG1", 100, 30, 50, delivery.MockAllDiscounts().GetDiscountByCoupon([]string{"OFR001"})))
	allPackages = append(allPackages, delivery.NewPackage("PKG2", 100, 125, 75, delivery.MockAllDiscounts().GetDiscountByCoupon([]string{"OFR008"})))
	allPackages = append(allPackages, delivery.NewPackage("PKG3", 100, 100, 175, delivery.MockAllDiscounts().GetDiscountByCoupon([]string{"OFR003"})))
	allPackages = append(allPackages, delivery.NewPackage("PKG4", 100, 60, 110, delivery.MockAllDiscounts().GetDiscountByCoupon([]string{"OFR002"})))
	allPackages = append(allPackages, delivery.NewPackage("PKG5", 100, 95, 155, delivery.MockAllDiscounts().GetDiscountByCoupon([]string{"NA"})))

	return allPackages
}

func getAllPackages_test2() delivery.Packages {

	allPackages := make(delivery.Packages, 0)
	allPackages = append(allPackages, delivery.NewPackage("PKG1", 100, 30, 50, delivery.MockAllDiscounts().GetDiscountByCoupon([]string{"OFR001"})))
	allPackages = append(allPackages, delivery.NewPackage("PKG2", 100, 30, 110, delivery.MockAllDiscounts().GetDiscountByCoupon([]string{"OFR008"})))
	allPackages = append(allPackages, delivery.NewPackage("PKG3", 100, 30, 90, delivery.MockAllDiscounts().GetDiscountByCoupon([]string{"OFR003"})))
	allPackages = append(allPackages, delivery.NewPackage("PKG5", 100, 30, 20, delivery.MockAllDiscounts().GetDiscountByCoupon([]string{"NA"})))

	return allPackages
}

//Test case to debug problem2

func TestGetPackagesForVehicle(t *testing.T) {
	allPackages := getAllPackages_test2()
	//sort.Sort(allPackages)

	p := calculatePackagesGroups(200, allPackages)
	sort.Sort(p)
	vehicles := make(delivery.Vehicles, 0)
	vehicles = append(vehicles, delivery.NewVehicle(1, 70, 200))
	vehicles = append(vehicles, delivery.NewVehicle(2, 70, 200))

	index := 0
	for {
		if len(p) == index {
			break
		}
		for _, v := range vehicles {
			fmt.Println(v.Id, v.NextDeliveryTime)
			p[index].SetDeliveryTimeForPackages(v.MaxSpeed, v.NextDeliveryTime)
			v.SetNextDeliveryTime(p[index].GetTotalDeliveryTime())
			index++
		}
		sort.Sort(vehicles)
	}

	for _, groups := range p {
		for _, v := range groups {
			fmt.Println(v.Id, v.DeliveryTime)
		}
	}

}
