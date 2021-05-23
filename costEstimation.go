package main

import (
	"fmt"

	"github.com/anil-appface/go-kikidelivery/delivery"
)

//TODO:: Try reading the scanf from main.go file.
func findDeliveryCostForPackages() delivery.Packages {
	var baseDeliveryPrice int
	var totalPackages int
	//fmt.Println("Enter Base Delivery Price & Total Packages in format (base_delivery_cost no_of_packges)")
	fmt.Scanf("%d %d", &baseDeliveryPrice, &totalPackages)

	//fmt.Println("Enter Packages details in format (pkg_id pkg_weight_in_kg distance_in_km offer_code):")
	allPackages := delivery.Packages{}

	//getting all the discounts.
	discounts := delivery.MockAllDiscounts()

	//Get the all input packages details
	for i := 0; i < totalPackages; i++ {
		var packageID, coupon string
		var weight, distance int
		fmt.Scanf("%s %d %d %s", &packageID, &weight, &distance, &coupon)

		//Get the discount details based on coupon
		discounts := discounts.GetDiscountByCoupon([]string{coupon})

		//create a new package
		p := delivery.NewPackage(packageID, baseDeliveryPrice, distance, weight, discounts)

		allPackages = append(allPackages, p)
	}

	//calculate & set total delivery cost to the package
	allPackages.CalculateDeliveryCost()

	return allPackages
}
