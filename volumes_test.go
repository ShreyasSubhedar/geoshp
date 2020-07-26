package geoshp

import "testing"

func TestCylinderVol(t *testing.T) {
	ans := CylinderVol(4, 9)
	if ans != 452.16 {
		t.Error("Expected :  452.16 , but found: ", ans)
	}

}

func TestSphereVol(t *testing.T) {
	ans := SphereVol(4)
	if ans != 200.96 {
		t.Error("Expected :  200.96 , but found: ", ans)
	}
}

func TestConeVol(t *testing.T) {
	ans := ConeVol(4, 3)
	if ans != 50.239999999999995 {
		t.Error("Expected :  50.239999999999995 , but found: ", ans)
	}
}

func TestHemisphereVol(t *testing.T) {
	ans := HemisphereVol(4)
	if ans != 133.97333333333333 {
		t.Error("Expected :  133.97333333333333 , but found: ", ans)
	}
}
