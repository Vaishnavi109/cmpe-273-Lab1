package main
import "fmt"
type Shape interface {
	Perimeter() float64
}
type Rectangle struct {
	length , breadth float64
}
type Circle struct {
	radius float64
}

func (rec Rectangle) Perimeter() float64 {
   var recPerimeter float64 = float64(2 *(rec.length + rec.breadth))
   return recPerimeter
}
func (cir Circle) Perimeter() float64{
	var circlePerimeter float64 = 2 * 3.14 * cir.radius
	return circlePerimeter
}

func main() {
	var length1, breadth1, rad  float64
	var rectanglePerimeter,circlePerimeter  float64 
	var choice int
			fmt.Println("Enter Your Choice :")
			fmt.Println("1. Perimeter Of Rectangle")
			fmt.Println("2. Perimeter Of Circle")
			fmt.Println("Enter Your Choice:")
			fmt.Scanf("%d\n",&choice)
			switch choice{
				case 1:
					fmt.Println("Enter Length Of Rectangle")
					fmt.Scanf("%f\n", &length1)
					
					fmt.Println("Enter Breadth Of Rectangle")
					fmt.Scanf("%f\n",&breadth1)
					
					if length1 > 0 && breadth1 > 0 {
					rec := Rectangle{length:length1,breadth:breadth1}
					sRec := Shape(rec)
					rectanglePerimeter = sRec.Perimeter()
					fmt.Println("Perimeter Of Rectangle " ,rectanglePerimeter)
					}else {
					 fmt.Println("Invalid Length or Breadth")
					 }
				case 2:
					fmt.Println("Enter Radius for Circle")
					fmt.Scanf("%f\n", &rad)
					
					if rad > 0 {
					cir := Circle{radius:rad }
					sCir:= Shape(cir)
					circlePerimeter = sCir.Perimeter()
					fmt.Println("Perimeter Of Circle " ,circlePerimeter)
					}else{
					fmt.Println("Invalid Radius")}

				}

}
