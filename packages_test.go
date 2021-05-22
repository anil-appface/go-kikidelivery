package main

import (
	"testing"
)

// Get total weight
func TestGetTotalWeight(t *testing.T) {

	pkg1 := NewPackage("PKG1", 100, 30, 100, nil)
	pkg2 := NewPackage("PKG2", 100, 30, 200, nil)
	packages := Packages{pkg1, pkg2}
	totalWeightOfPackages := packages.GetTotalWeight()
	expectedTotalWeightOfPackages := 300
	if totalWeightOfPackages != expectedTotalWeightOfPackages {
		t.Fatal("something went wrong while calculating the total weight of the packages")
	}
}

// Setting and Getting delivery time for packages
func TestSetAndGetDeliveryTimeForPackages(t *testing.T) {

	pkg1 := NewPackage("PKG1", 100, 40, 100, nil)
	pkg2 := NewPackage("PKG2", 100, 30, 200, nil)
	packages := Packages{pkg1, pkg2}
	packages.SetDeliveryTimeForPackages(10, 0)
	totalDeliveryTime := packages.GetTotalDeliveryTime()
	expectedTotalDeliveryTime := float32(40 / 10) // (maximum distance)/ (maxSpeed)
	if totalDeliveryTime != expectedTotalDeliveryTime {
		t.Fatal("something went wrong while calculating the total delivery time of packages")
	}
}
