package main

import "fmt"
import m "WorkSpace/Fibonacci/math"

func main(){
	var input int
	fmt.Print("Enter the value:")
	fmt.Scanf("%d", &input)
	if input > 0{
		output := m.Fibonacci(input) 
		fmt.Print(output)
	}else{
		fmt.Print("Invalid Input")}
}