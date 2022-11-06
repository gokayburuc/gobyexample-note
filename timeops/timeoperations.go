package timeops

import (
	"fmt"
	"time"
)

func TimeOps() {
	now := time.Now()
	fmt.Println(now)
	month := time.Now().Month()
	fmt.Println("month : ", month)
	twentyfive := now.AddDate(0, 0, 25)
	fmt.Println(twentyfive)

	//time.Sleep(3)
}
