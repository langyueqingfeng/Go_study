package main

import (
	"fmt"
	"github.com/agoussia/godes"
)

/*
* @Author:hanyajun
* @Date:2019/6/18 14:56
* @Name:sim
* @Function:
 */

// the arrival and service are two random number generators for the uniform  distribution
var arrival *godes.UniformDistr = godes.NewUniformDistr(true)

// the Visitor is a Runner
// any type of the Runner should be defined as struct
// with the *godes.Runner as anonimous field
type Visitor struct {
	*godes.Runner
	number int
}

var visitorsCount int = 0

func (vst *Visitor) Run() { // Any runner should have the Run method
	fmt.Printf(" %-6.3f \t Visitor # %v arrives \n", godes.GetSystemTime(), vst.number)
}
func main() {
	var shutdown_time float64 = 8 * 60
	godes.Run()
	for {
		//godes.Stime is the current simulation time
		if godes.GetSystemTime() < shutdown_time {
			//the function acivates the Runner
			godes.AddRunner(&Visitor{&godes.Runner{}, visitorsCount})
			//this advance the system time
			godes.Advance(arrival.Get(0, 70))
			visitorsCount++
		} else {
			break
		}
	}
	// waits for all the runners to finish the Run()
	godes.WaitUntilDone()
}
