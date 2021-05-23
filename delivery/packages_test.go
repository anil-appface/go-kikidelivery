package delivery

import (
	"sort"
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

// Test sorting of packages.
func TestSortPackages(t *testing.T) {

	expectedMaxWeighPackage := 175
	expectedMinWeighPackage := 50

	pkg1 := NewPackage("PKG1", 100, 30, 50, MockAllDiscounts().GetDiscountByCoupon([]string{"OFR001"}))
	pkg2 := NewPackage("PKG2", 100, 125, 75, MockAllDiscounts().GetDiscountByCoupon([]string{"OFR008"}))
	pkg3 := NewPackage("PKG3", 100, 100, 175, MockAllDiscounts().GetDiscountByCoupon([]string{"OFR003"}))
	pkg4 := NewPackage("PKG4", 100, 60, 110, MockAllDiscounts().GetDiscountByCoupon([]string{"OFR002"}))
	pkg5 := NewPackage("PKG5", 100, 95, 155, MockAllDiscounts().GetDiscountByCoupon([]string{"NA"}))

	allPackages := NewPackages(pkg1, pkg2, pkg3, pkg4, pkg5)
	sort.Sort(allPackages)
	if allPackages[0].Weight != expectedMaxWeighPackage {
		t.Fatal("sorting of descending failed")
	}

	if allPackages[4].Weight != expectedMinWeighPackage {
		t.Fatal("sorting of descending failed")
	}
}
