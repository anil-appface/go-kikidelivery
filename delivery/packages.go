package delivery

// Packages
// Array of Package.
type Packages []*Package

// Implementing sort.Sort interface
func (me Packages) Len() int           { return len(me) }
func (me Packages) Swap(i, j int)      { me[i], me[j] = me[j], me[i] }
func (me Packages) Less(i, j int) bool { return me[i].Weight < me[j].Weight }

func NewPackages(pkgs ...*Package) Packages {
	packages := Packages{}
	for _, v := range pkgs {
		packages = append(packages, v)
	}
	return packages
}

// GetTotalWeight
// Iterates through the package and calculates the weight of packages.
// @returns total weight of the given packages
func (me Packages) GetTotalWeight() int {
	weight := 0
	for _, p := range me {
		weight += p.Weight
	}
	return weight
}

/*
 SetDeliveryTimeForPackages
 Calculates and sets the total delivery time for the packages in the vehicle and assign the delivery time to packages.
 @input
 maxSpeed: maxSpeed of the vehicle in which the package carrying.
 additionalTime: Any additional Time.
*/
// TODO:: This function can be improved
func (me Packages) SetDeliveryTimeForPackages(maxSpeed int, additionalTime float32) {
	for _, p := range me {
		p.DeliveryTime = roundoff(additionalTime + float32(p.Distance)/float32(maxSpeed))
	}
}

// GetTotalDeliveryTime
// Calculates the total delivery time of the packages to be shipped.
// @return float32
func (me Packages) GetTotalDeliveryTime() float32 {
	var deliveryTime float32
	for _, p := range me {
		if deliveryTime < p.DeliveryTime {
			deliveryTime = p.DeliveryTime
		}
	}
	return deliveryTime
}

// HasPackageWithID
// Checks whether the package exists or not
// @return bool
func (me Packages) ContainsPackage(pack *Package) bool {
	for _, p := range me {
		if p.Id == pack.Id {
			return true
		}
	}
	return false
}

// CalculateDeliveryCost
// Utiliy method to calculate the delivery cost of the packages.
// It also sets the totalCost variable for each Package.
func (me Packages) CalculateDeliveryCost() {
	for _, p := range me {
		p.CalculateDeliveryCost()
	}
}
