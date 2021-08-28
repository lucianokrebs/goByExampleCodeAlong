package goSwitch

import (
	"fmt"
	"github.com/lucianokrebs/goByExampleCodeAlong/pkg/goPrint"
	"time"
)

func GoSwitch() {
	i := 2

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Don't know")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a int")
		default:
			fmt.Printf("Don't know type: %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("Hey!")

	goPrint.PrintDashes()
}
