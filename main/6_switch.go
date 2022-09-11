package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	//use switch case instead of if - else if - else structure. variable in switch part is controlled in case part. one case is executed others doens work. cases dont need break keyword. after a case works, switch case block is automatically done

	//in switch part, a variable can be assigned like in if statement
	switch os := runtime.GOOS; os {

	case "darwin":
		fmt.Println("Darwin")
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Println("Unkown system")
	}

	//there can be condition in any case part. switch part can be blank. this means that switch true.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

}
