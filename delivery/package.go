package delivery

// Package
// Package details
type Package struct {
	Id           string
	Distance     int
	Weight       int
	BaseCost     int
	TotalCost    int
	DiscountCost int
	DeliveryTime float32
	Discounts    Discounts
}

// NewPackage
// Creates the New Package and returns the Package.
func NewPackage(id string, baseCost, distance, weight int, discounts Discounts) *Package {
	return &Package{
		Id:        id,
		BaseCost:  baseCost,
		Distance:  distance,
		Weight:    weight,
		Discounts: discounts,
	}
}

// IsDiscountApplicable
// Checks whether the package has discount or not based on conditions
// 1. package weight should be greater or equals to discount minimun package weight
// 2. package weight should be lesser or equals to discount maximum package weight
// 4. package distance should be greater or equals to discount minimum package distance
// 3. package distance should be lesser or equals to discount maximum package distance
func (me *Package) IsDiscountApplicable() bool {

	for _, discount := range me.Discounts {
		if discount.MinPackageWeight <= me.Weight &&
			me.Weight <= discount.MaxPackageWeight &&
			discount.MinDestinationDistance <= me.Distance &&
			me.Distance <= discount.MaxDestinationDistance {
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
	totalCost := me.BaseCost + (me.Weight * 10) + (me.Distance * 5)
	if me.IsDiscountApplicable() {
		me.DiscountCost = me.Discounts.calculateDiscountAmount(totalCost)
	}
	me.TotalCost = totalCost - me.DiscountCost
}
