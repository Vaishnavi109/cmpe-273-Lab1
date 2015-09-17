package main
import "testing"
import "time"
func Test_Sleep(t *testing.T){
	timeForSleep  := 3
	timeToCheck :=time.Now()
	Sleep(timeForSleep)
	timeAfterSleep:=time.Now()
	sleepDifference := timeAfterSleep.Sub(timeToCheck)
	
	if sleepDifference > time.Second*time.Duration(4) {
	t.Error("Sleep time exceeded 3 and got ",sleepDifference)
	}
	
}