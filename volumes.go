package geoshp

// SphereVol is to calculate the volume of sphere
func SphereVol(radius float64) float64 {
	return (4 / 3) * Pi * radius * radius * radius
}

// CylinderVol is to calculate the volume of cylinder
func CylinderVol(radius float64, height float64) float64 {
	return Pi * radius * radius * height
}

// ConeVol is to calculate the volume of cone
func ConeVol(radius float64, height float64) float64 {
	return (1.0 / 3.0) * Pi * radius * radius * height
}

// HemisphereVol is to calculate the volume of Hemisphere
func HemisphereVol(radius float64) float64 {
	return (2.0 / 3.0) * Pi * radius * radius * radius
}
