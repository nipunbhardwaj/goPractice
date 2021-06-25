package main

import "fmt"
import "time"

func main () {

	christmas := time.Date(time.Now().Year(), 12, 25, 0, 0, 0, 0, time.Now().Location());

	timeToChristmas := time.Until(christmas);

	fmt.Println("Milliseconds to christmas: ", int(timeToChristmas.Milliseconds()));
	fmt.Println("Seconds to christmas: ", int(timeToChristmas.Seconds()));
	fmt.Println("Minutes to christmas: ", int(timeToChristmas.Minutes()));
	fmt.Println("Hours to christmas: ", int(timeToChristmas.Hours()));
	fmt.Println("Days to christmas: ", int(timeToChristmas.Hours()/24));
}