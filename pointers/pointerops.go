package pointers

import "fmt"


func ZeroValue(x int) {
	x = 0
}

func ZeroValuePointer(x *int){
	*x = 0
}


func PointerTest(){
	x := 1
	fmt.Println("value\t:",x)

	ZeroValue(x)
	fmt.Println("zerovalue\t:", x)

	ZeroValuePointer(&x)
	//pointer ifadesi ile tanımlanan degeri getirir. 0 döndürür
	fmt.Println("zeropointer\t:",x)

	// pointer ifadesinin hafizada tutuldugu yeri getirir
	fmt.Println("pointer\t:", &x)
}