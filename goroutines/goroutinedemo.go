package goroutines

import (
	"fmt"
	"time"
)

func f(fromSource string) {
	for i := 0; i < 3; i++ {
		fmt.Println(fromSource, ":", i)
	}
}

func Operations() {

	f("Direct")
	// routine tanimladik
	go f("Goroutine")

	// isimsiz bir fonksiyon tanimladik
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
