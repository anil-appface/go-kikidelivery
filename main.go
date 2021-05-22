package main

import (
	"fmt"
	"os"
	"sort"
)

//TODO:: To use buffer reader for input reading. bufio.NewReader(os.Stdin)
func main() {

	problem := os.Args[1]

	if problem == "1" {
		// problem1 -> calulates delivery cost
		allPackages := problem1()

		//print output format
		for _, v := range allPackages {
			fmt.Printf("%s %d %d", v.id, v.discountCost, v.totalCost)
			fmt.Println()
		}
	}

	//Problem 2 -> calculates delivery estimation time
	if problem == "2" {
		allPackages := problem1()

		//fmt.Println("Enter Number of vehicles , Max Speed & Max carriable weight in format: (no_of_vehicles max_speed max_carriable_weight)")
		var vehicleCount, vehicleMaxSpeed, vehicleMaxLoad int
		fmt.Scanf("%d %d %d", &vehicleCount, &vehicleMaxSpeed, &vehicleMaxLoad)

		//initialising new vehicles
		vehicles := NewVehicles(vehicleCount, vehicleMaxSpeed, vehicleMaxLoad)

		//sort packages with maximum weight
		sort.Sort(allPackages)

		// calculate packagegroups
		packageGroups := calculatePackagesGroups(vehicleMaxLoad, allPackages)
		sort.Sort(packageGroups)

		index := 0
		for {
			if len(packageGroups) == index {
				break
			}
			for _, v := range vehicles {
				packageGroups[index].SetDeliveryTimeForPackages(v.maxSpeed, v.nextDeliveryTime)
				v.SetNextDeliveryTime(packageGroups[index].GetTotalDeliveryTime())
				index++
			}
			sort.Sort(vehicles)
		}

		// Output formats
		packages := packageGroups.ConvertToPackages()
		for _, v := range packages {
			fmt.Printf("%s %d %d %f", v.id, v.discountCost, v.totalCost, v.deliveryTime)
			fmt.Println()
		}

	}

}

func problem1() Packages {
	var baseDeliveryPrice int
	var totalPackages int
	//fmt.Println("Enter Base Delivery Price & Total Packages in format (base_delivery_cost no_of_packges)")
	fmt.Scanf("%d %d", &baseDeliveryPrice, &totalPackages)

	//fmt.Println("Enter Packages details in format (pkg_id pkg_weight_in_kg distance_in_km offer_code):")
	allPackages := Packages{}

	//getting all the discounts.
	discounts := allDiscounts()

	//Get the all input packages details
	for i := 0; i < totalPackages; i++ {
		var packageID, coupon string
		var weight, distance int
		fmt.Scanf("%s %d %d %s", &packageID, &weight, &distance, &coupon)

		//Get the discount details based on coupon
		discount := discounts.getDiscountByCoupon(coupon)

		//create a new package
		p := NewPackage(packageID, baseDeliveryPrice, distance, weight, discount)

		allPackages = append(allPackages, p)
	}

	//calculate & set total delivery cost to the package
	allPackages.CalculateDeliveryCost()

	return allPackages
}

// calculatePackagesGroups
// TODO:: Make this method more readable.
func calculatePackagesGroups(maxLoad int, packages Packages) PackagesGroup {

	// small piece of function helps to get the packages according to their weights
	getMaxWeightPackages := func(packages Packages) (Packages, Packages) {
		totalPackages := make(Packages, 0, len(packages))
		remainingPackages := make(Packages, 0, len(packages))
		totalWeight := 0
		for _, pack := range packages {

			if pack.weight <= maxLoad {
				if (totalWeight + pack.weight) > maxLoad {
					remainingPackages = append(remainingPackages, pack)
					continue
				}
			}
			totalWeight += pack.weight
			totalPackages = append(totalPackages, pack)
		}
		return totalPackages, remainingPackages
	}

	var p Packages
	group := make(PackagesGroup, 0)

	// Creating Package Groups according to the package count & packages weight
	for {
		if len(packages) == 0 {
			break
		}
		p, packages = getMaxWeightPackages(packages)
		group = append(group, p)
	}

	return group
}

// 100 5
// PKG1 50 30 OFR001
// PKG2 75 125 OFFR0008
// PKG3 175 100 OFFR003
// PKG4 110 60 OFFR002
// PKG5 155 95 NA
// 2 70 200
