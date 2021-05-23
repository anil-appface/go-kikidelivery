package delivery

import (
	"testing"
)

// Test to check whether the package has valid discount or not
func TestCheckIsDiscountApplicableForPackage(t *testing.T) {

	//negative test case
	//-> package distance is within range
	//-> package weight is less than discount_min_weight
	discount1 := NewDiscount("OFR001", 10, 10, 200, 70, 200)
	pkg1 := NewPackage("PKG1", 100, 30, 50, NewDiscounts(discount1))
	if pkg1.IsDiscountApplicable() {
		t.Fatal("something went wrong while calculating the discount applicable for packages for PKG1")
	}

	//negative test case
	//-> package distance is within range
	//-> package weight is greater than discount_max_weight
	pkg2 := NewPackage("PKG2", 100, 30, 300, NewDiscounts(discount1))
	if pkg2.IsDiscountApplicable() {
		t.Fatal("something went wrong while calculating the discount applicable for packages for PKG2")
	}

	//negative test case
	//-> package distance is less than discount_minimum_distance
	//-> weight is within range
	pkg3 := NewPackage("PKG3", 100, 5, 100, NewDiscounts(discount1))
	if pkg3.IsDiscountApplicable() {
		t.Fatal("something went wrong while calculating the discount applicable for packages for PKG3")
	}

	//negative test case
	//-> package distance is greater than discount_max_distance
	//-> weight is within range
	pkg4 := NewPackage("PKG4", 100, 300, 100, NewDiscounts(discount1))
	if pkg4.IsDiscountApplicable() {
		t.Fatal("something went wrong while calculating the discount applicable for packages for PKG4")
	}

	//negative test case
	//-> package distance is greater than discount_max_distance
	//-> package weight is greater than discount_max_weight
	pkg5 := NewPackage("PKG5", 100, 300, 300, NewDiscounts(discount1))
	if pkg5.IsDiscountApplicable() {
		t.Fatal("something went wrong while calculating the discount applicable for packages for PKG5")
	}

	//negative test case
	//-> package distance is lesser than discount_min_distance
	//-> package weight is lesser than discount_min_weight
	pkg6 := NewPackage("PKG6", 100, 9, 50, NewDiscounts(discount1))
	if pkg6.IsDiscountApplicable() {
		t.Fatal("something went wrong while calculating the discount applicable for packages for PKG6")
	}

	//negative test case
	//-> if discount is nil
	pkg7 := NewPackage("PKG6", 100, 9, 50, nil)
	if pkg7.IsDiscountApplicable() {
		t.Fatal("something went wrong while calculating the discount applicable for packages for PKG7")
	}

	//positive test case
	//-> package distance is within range
	//-> package weight is within range
	pkg8 := NewPackage("PKG6", 100, 50, 100, NewDiscounts(discount1))
	if !pkg8.IsDiscountApplicable() {
		t.Fatal("something went wrong while calculating the discount applicable for packages for PKG8")
	}
}

// Test to check whether the delivery cost calculation is correct
func TestCheckDeliveryCostForPackage(t *testing.T) {

	discount1 := NewDiscount("OFR001", 10, 10, 200, 70, 200)
	pkg1 := NewPackage("PKG1", 100, 30, 100, NewDiscounts(discount1))
	pkg1.CalculateDeliveryCost()
	expectedPKG1DeliveryCost := 1130
	if pkg1.TotalCost != expectedPKG1DeliveryCost {
		t.Fatal("something went wrong while calculating the delivery cost package PKG1")
	}

	pkg2 := NewPackage("PKG1", 100, 30, 100, nil)
	pkg2.CalculateDeliveryCost()
	expectedPKG2DeliveryCost := 1250
	if pkg2.TotalCost != expectedPKG2DeliveryCost {
		t.Fatal("something went wrong while calculating the delivery cost package PKG2")
	}
}
