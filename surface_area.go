package geoshp

// SphereArea is to calculate the total suraface area of sphere
func SphereArea(radius float64) float64 {
	return 4 * Pi * radius * radius
}

// CylinderArea is to calculate the total suraface area of cylinder
func CylinderArea(radius float64, height float64) float64 {
	return 2 * Pi * radius * (radius + height)
}

// ConeArea is to calculate the total suraface area of of cone
func ConeArea(radius float64, length float64) float64 {
	return Pi * radius * (radius + length)
}

// HemisphereArea is to calculate the Areaume of Hemisphere
func HemisphereArea(radius float64) float64 {
	return 3 * Pi * radius * radius
}
