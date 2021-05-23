package main

import (
	"fmt"
	"sort"

	"github.com/anil-appface/go-kikidelivery/delivery"
)

//TODO:: read scanf from main.go
func calculateDeliveryTimeEstimation() delivery.Packages {
	// calculate delivery cost
	allPackages := findDeliveryCostForPackages()

	//fmt.Println("Enter Number of vehicles , Max Speed & Max carriable weight in format: (no_of_vehicles max_speed max_carriable_weight)")
	var vehicleCount, vehicleMaxSpeed, vehicleMaxLoad int
	fmt.Scanf("%d %d %d", &vehicleCount, &vehicleMaxSpeed, &vehicleMaxLoad)

	//initialising new vehicles
	vehicles := delivery.NewVehicles(vehicleCount, vehicleMaxSpeed, vehicleMaxLoad)

	//sort packages with maximum weight
	sort.Sort(allPackages)

	// calculate packagegroups
	packageGroups := calculatePackagesGroups(vehicleMaxLoad, allPackages)
	//sort.Sort(packageGroups)

	index := 0
	for {
		if len(packageGroups) == index {
			break
		}
		for _, v := range vehicles {
			packageGroups[index].SetDeliveryTimeForPackages(v.MaxSpeed, v.NextDeliveryTime)
			v.SetNextDeliveryTime(packageGroups[index].GetTotalDeliveryTime())
			index++
		}
		sort.Sort(vehicles)
	}

	// Output formats
	return packageGroups.ConvertToPackages()
}

// calculatePackagesGroups
// TODO:: Make this method more readable.
func calculatePackagesGroups(maxLoad int, packages delivery.Packages) delivery.PackagesGroup {

	remainingPackages := delivery.NewPackages(packages...)

	// small piece of function helps to get the packages according to their weights
	getPossiblePackages := func(packages delivery.Packages) delivery.PackagesGroup {
		group := make(delivery.PackagesGroup, 0)

		for i := 0; i < len(packages); i++ {
			packageItems := delivery.Packages{packages[i]}
			for j := i; j < len(packages); j++ {
				if i == j {
					continue
				}
				if packages[j].Weight+packageItems.GetTotalWeight() <= maxLoad {
					packageItems = append(packageItems, packages[j])
				} else {
					group = append(group, packageItems)
					packageItems = delivery.Packages{packages[i]}

					if packages[j].Weight+packageItems.GetTotalWeight() <= maxLoad {
						packageItems = append(packageItems, packages[j])
					}
				}
			}

			group = append(group, packageItems)
		}

		return group
	}

	toBeDeliveredPackages := delivery.PackagesGroup{}
	for {
		if len(remainingPackages) == 0 {
			break
		}
		packagesGroup := getPossiblePackages(remainingPackages)

		// Creating Package Groups according to the package count & packages weight
		sort.Sort(packagesGroup)

		remainingPackages = delivery.Packages{}
		toBeDeliveredPackages = append(toBeDeliveredPackages, packagesGroup[0])
		for _, v := range packages {
			if !toBeDeliveredPackages.ConvertToPackages().ContainsPackage(v) {
				remainingPackages = append(remainingPackages, v)
			}
		}
	}

	return toBeDeliveredPackages
}

// 100 4
// P1 50 30 O1
// P2 110 30 O2
// P3 90 30 O3
// P4 20 30 O4
// 2 70 200

// 50, 110, 90, 20
// 110, 80, 50, 20

// 50, 75, 175, 110, 155
// 175, 155, 110, 75, 50
