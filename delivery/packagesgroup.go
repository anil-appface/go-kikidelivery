package delivery

type PackagesGroup []Packages

// Implementing sort.Sort interface
func (me PackagesGroup) Len() int      { return len(me) }
func (me PackagesGroup) Swap(i, j int) { me[i], me[j] = me[j], me[i] }
func (me PackagesGroup) Less(i, j int) bool {
	if len(me[i]) == len(me[j]) {
		return me[i].GetTotalWeight() > me[j].GetTotalWeight()
	}
	return len(me[i]) > len(me[j])
}

// ConvertToPackages
// Converts groups of packages to packages array
func (me PackagesGroup) ConvertToPackages() Packages {
	packages := Packages{}
	for _, v := range me {
		packages = append(packages, v...)
	}
	return packages
}
