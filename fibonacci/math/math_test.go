package math
import "testing"

func TestFiboForInterger(t *testing.T){
	var result int
	result = Fibonacci(10)
	if result != 55 {
	t.Error("Expected 1.5, got ", result)
	}
	
}
func TestFiboForZero(t *testing.T){
	var result int
	result = Fibonacci(0)
	if result != 0 {
	t.Error("Expected 0 , got ",result)

}
}
func TestFiboForOne(t *testing.T){
	var result int
	result = Fibonacci(1)
	if result != 1 {
	t.Error("Expected 0 , got ",result)

}
}
/*
func TestFiboForNegavtiveValues(t *testing.T){
	var result int
	result = Fibonacci(-1)
	if result == 0{
	t.Error("Expected 0 , got ",result)

}
}
*/