package geoshp

import "testing"

func TestCylinderArea(t *testing.T) {
	ans := CylinderArea(7, 12)
	if ans != 835.24 {
		t.Error("Expected :  835.24 , but found: ", ans)
	}

}

func TestSphereArea(t *testing.T) {
	ans := SphereArea(8)
	if ans != 803.84 {
		t.Error("Expected :  803.84 , but found: ", ans)
	}
}

func TestConeArea(t *testing.T) {
	ans := ConeArea(7, 7)
	if ans != 307.72 {
		t.Error("Expected :  307.72 , but found: ", ans)
	}
}

func TestHemisphereArea(t *testing.T) {
	ans := HemisphereArea(11)
	if ans != 1139.8200000000002 {
		t.Error("Expected :  1139.8200000000002 , but found: ", ans)
	}
}
