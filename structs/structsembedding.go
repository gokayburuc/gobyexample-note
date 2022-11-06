package structs

import (
	"fmt"
)

type base struct {
	num int
}

// describe() isimli bir method tanımladık base struct u altında
func (b base) describe() string {
	return fmt.Sprintf("Base with num = %v", b.num)
}

// burada container isimli bir struct tanımladık
// icine de ifademize ait base struct ifadesini dahil ettik
type container struct {
	base
	str string
}

func EmbedTest() {
	myContainer := container{
		base: base{
			num: 1,
		},
		str: "Test Name",
	}
	fmt.Printf("co = {num: %v, str : %v}\n", myContainer.num, myContainer.str)

	fmt.Println("also num:", myContainer.num)

	fmt.Println("describe:", myContainer.describe())

	type describer interface {
		describe() string
	}

	var d describer = myContainer
	fmt.Println("describer : ", d.describe())

}
