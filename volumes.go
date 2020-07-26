package geoshp

// SphereVol is to calculate the volume of sphere
func SphereVol(radius float64) float64 {
	return (4 / 3) * Pi * radius * radius * radius
}

// CyclinderVol is to calculate the volume of cyclinder
func CyclinderVol(radius float64, height float64) float64 {
	return Pi * radius * radius * height
}

// ConeVol is to calculate the volume of cone
func ConeVol(radius float64, height float64) float64 {
	return (1 / 3) * Pi * radius * radius * height
}

// HemisphereVol is to calculate the volume of Hemisphere
func HemisphereVol(radius float64) float64 {
	return (2 / 3) * Pi * radius * radius * radius
}
