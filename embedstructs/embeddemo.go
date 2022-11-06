package embedstructs

import "fmt"


type base struct{
	num int
}


func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}


func EmbedTest(){
	cont := container{
		base: base{
			num :1,
		},
		str : "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", cont.num, cont.str)

	fmt.Println("also num :", cont.base.num)

	fmt.Println("describe :", cont.describe())


	type describer interface {
		describe() string
	}

	var d describer = cont

	fmt.Println("describer :", d.describe())
}

