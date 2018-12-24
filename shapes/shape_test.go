package shape

import (
	"testing"
)

func TestTriangleGetArea(t *testing.T) {
	tri := triangle{height: 10, width: 10}

	area := tri.getArea()
	if area != 50 {
		t.Errorf("Expected area 50, but got %v", area)
	}
}

func TestRectangleGetArea(t *testing.T) {
	r := rectangle{height: 10, width: 10}

	area := r.getArea()
	if area != 100 {
		t.Errorf("Expected area 100, but got %v", area)
	}
}
