package main

type Discount struct {
	coupon                 string
	percentage             int
	minDestinationDistance int
	maxDestinationDistance int
	minPackageWeight       int
	maxPackageWeight       int
}

func newDiscount(coupon string, percentage, minDestinationDistance, maxDestinationDistance, minPackageWeight, maxPackageWeight int) Discount {
	return Discount{
		coupon:                 coupon,
		percentage:             percentage,
		minDestinationDistance: minDestinationDistance,
		maxDestinationDistance: maxDestinationDistance,
		minPackageWeight:       minPackageWeight,
		maxPackageWeight:       maxPackageWeight,
	}
}

func (me Discount) calculateDiscountAmount(cost int) int {
	return (cost / 100 * me.percentage)
}

type discounts []Discount

func (me discounts) getDiscountByCoupon(coupon string) *Discount {
	for _, discount := range me {
		if discount.coupon == coupon {
			return &discount
		}
	}
	return nil
}

func allDiscounts() discounts {
	allDiscounts := make([]Discount, 0, 3)
	allDiscounts = append(allDiscounts, newDiscount("OFR001", 10, 0, 200, 70, 200))
	allDiscounts = append(allDiscounts, newDiscount("OFR002", 7, 50, 150, 100, 250))
	allDiscounts = append(allDiscounts, newDiscount("OFFR002", 7, 50, 150, 100, 250))
	allDiscounts = append(allDiscounts, newDiscount("OFR003", 5, 50, 250, 10, 150))
	return allDiscounts
}
