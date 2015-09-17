package main
import ("fmt"
		"time")
/*		
func main(){
	var sleepTime int
	fmt.Print("Enter Sleep Duration in Seconds")
	fmt.Scanf("%d" ,&sleepTime)
	Sleep(sleepTime)
}
*/
func Sleep(sleepTime int){
	<-time.After(time.Second * time.Duration(sleepTime))
	fmt.Print("Sleep Over")
}