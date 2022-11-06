package structs

import "fmt"


type person struct{
	name string
	age int32
}

func newPerson(nameinput string) *person {
	p:= person{name: nameinput}
	p.age = 42
	/* sistemdeki pointer degerini getir & ve bu degeri *ile ifade seklinde goruntule
	*/
	return &p
}

func PersonTest(){
	// sistemde bir kisi tanimla/
	secondperson := newPerson("Esin")
	fmt.Println(secondperson)
	fmt.Println(secondperson.name)
	fmt.Println(secondperson.age)

/*sistemde ikinci bir kisi tanımla ve degerleri ekrana yazdir*/

	firstperson :=  person{name :"Gokay", age: 35}
	fmt.Println("FIRST PERSON : ", firstperson)
	fmt.Println("firstperson.age : " , firstperson.age)
	fmt.Println("firstperson.name:", firstperson.name)

}




