package variables

import "fmt"

func GetVariables() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	var g, h bool = true, false
	fmt.Println(g, h)
	fmt.Println("true and false" , g && h)
	fmt.Println("false and false", false && h)
	fmt.Println("true or false" , g || h)

}
