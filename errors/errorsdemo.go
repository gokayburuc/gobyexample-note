package errors

import (
	"errors"
	"fmt"
)

// f1 isimli bir fonksiyon tanımladık. arg isimli bir integer deger alir
// int ya da error döndürür
func f1(arg int) (int, error) {
	if arg == 42 {

		return -1, errors.New("can't work with 42")
	}

	//eger calisirsa arg degerine 3 ekle ve döndür, hata gelirse nil olarak döndür
	return arg + 3, nil
}

// arg ve prob ifadelerini iceren bir argError struct ifadesi tanımladık
type argError struct {
	arg  int
	prob string
}

// argError isimli struct ifadesine baglı bir method tanımladık. string ifade döndürür
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		//argError pointer degerini ekrana cagirdik. Bu sistemde tutulan argError degeridir
		return -1, &argError{arg, "can't work with it"}
	}

	return arg + 3, nil
}

func ErrorTester() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)

	//TODO: bu yazıma ait kısımları tekrar gozden gecir => if ae,ok := e.(*argError); ok{
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
