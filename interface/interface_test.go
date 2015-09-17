package main
import "testing"

func Test_RectanglePerimeter(t *testing.T){
	var result float64
	rec := Rectangle{length:5,breadth:4}
	result = rec.Perimeter()
	if result != 18 {
	t.Error("Expected 18, got ", result)
	}
	
}
func Test_RectanglePerimeterForFloatValue(t *testing.T){
	var result float64
	rec := Rectangle{length:2.5,breadth:3.4}
	result = rec.Perimeter()
	if result != 11.8 {
	t.Error("Expected 18, got ", result)
	}
	
}
func Test_CirclePerimeter(t *testing.T){
	var result float64
	cir := Circle{radius:4 }
	result = cir.Perimeter()
	if result != 25.12 {
	t.Error("Expected 18, got ", result)
	}
	
}

func Test_CirclePerimeterForFloatValue(t *testing.T){
	var result float64
	cir := Circle{radius:4.5}
	result = cir.Perimeter()
	if result != 28.26 {
	t.Error("Expected 18, got ", result)
	}
	
}
