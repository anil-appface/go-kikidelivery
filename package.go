package main

// Package
// Package details
type Package struct {
	id           string
	distance     int
	weight       int
	baseCost     int
	totalCost    int
	discountCost int
	deliveryTime float32
	discount     *Discount
}

// NewPackage
// Creates the New Package and returns the Package.
func NewPackage(id string, baseCost, distance, weight int, discount *Discount) *Package {
	return &Package{
		id:       id,
		baseCost: baseCost,
		distance: distance,
		weight:   weight,
		discount: discount,
	}
}

// IsDiscountApplicable
// Checks whether the package has discount or not based on conditions
// 1. package weight should be greater or equals to discount minimun package weight
// 2. package weight should be lesser or equals to discount maximum package weight
// 4. package distance should be greater or equals to discount minimum package distance
// 3. package distance should be lesser or equals to discount maximum package distance
func (me *Package) IsDiscountApplicable() bool {

	if me.discount != nil {
		if me.discount.minPackageWeight <= me.weight && me.weight <= me.discount.maxPackageWeight && me.discount.minDestinationDistance <= me.distance && me.distance <= me.discount.maxDestinationDistance {
			return true
		}
	}
	return false
}

// CalculateDeliveryCost
// Calculates delivery cost of the package and if there is any discount, it will be discounted from the totalCost of package.
// Formulae:
// totalCost_of_Package = (base_cost + (package_weight * 10) + (package_distance * 5)) - discountPrice
func (me *Package) CalculateDeliveryCost() {
	totalCost := me.baseCost + (me.weight * 10) + (me.distance * 5)
	if me.IsDiscountApplicable() {
		me.discountCost = me.discount.calculateDiscountAmount(totalCost)
	}
	me.totalCost = totalCost - me.discountCost
}
