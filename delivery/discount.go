package delivery

import "github.com/anil-appface/go-kikidelivery/globals"

type Discount struct {
	Coupon                 string
	Percentage             int
	MinDestinationDistance int
	MaxDestinationDistance int
	MinPackageWeight       int
	MaxPackageWeight       int
}

func NewDiscount(coupon string, percentage, minDestinationDistance, maxDestinationDistance, minPackageWeight, maxPackageWeight int) *Discount {
	return &Discount{
		Coupon:                 coupon,
		Percentage:             percentage,
		MinDestinationDistance: minDestinationDistance,
		MaxDestinationDistance: maxDestinationDistance,
		MinPackageWeight:       minPackageWeight,
		MaxPackageWeight:       maxPackageWeight,
	}
}

func (me Discount) calculateDiscountAmount(cost int) int {
	return (cost / 100 * me.Percentage)
}

type Discounts []*Discount

func NewDiscounts(d ...*Discount) Discounts {
	discounts := Discounts{}
	for _, disc := range d {
		discounts = append(discounts, disc)
	}
	return discounts
}

func (me Discounts) calculateDiscountAmount(cost int) int {
	totalDiscountAmount := 0
	for _, v := range me {
		totalDiscountAmount += v.calculateDiscountAmount(cost)
	}
	return totalDiscountAmount
}

func (me Discounts) GetDiscountByCoupon(coupons []string) Discounts {
	discounts := Discounts{}
	for _, discount := range me {
		if contains(coupons, discount.Coupon) {
			discounts = append(discounts, discount)
		}
	}
	return discounts
}

//utitlity method
func contains(arr []string, value string) bool {
	for _, v := range arr {
		if value == v {
			return true
		}
	}
	return false
}

// This function is just a mock to get all discounts
func MockAllDiscounts() Discounts {
	allDiscounts := Discounts{}
	allDiscounts = append(allDiscounts, NewDiscount(globals.Coupon1, 10, 0, 200, 70, 200))
	allDiscounts = append(allDiscounts, NewDiscount(globals.Coupon2, 7, 50, 150, 100, 250))
	allDiscounts = append(allDiscounts, NewDiscount(globals.Coupon3, 7, 50, 150, 100, 250))
	allDiscounts = append(allDiscounts, NewDiscount(globals.Coupon4, 5, 50, 250, 10, 150))
	return allDiscounts
}
