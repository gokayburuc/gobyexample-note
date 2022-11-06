package methods

import "fmt"


type rectangle struct{
	width, height int
}

func (r *rectangle) area() int{
	/*
	rectangle isimli struct ifadesine bağlı
	bir area isimli method tanımladik
	tanımnladıgımız degerden r isimli
	bir degiskene atama yaptık
	sonucu ekrana dondurduk
	,*/
	return r.width * r.height
}


func (r *rectangle) perim() int{
	// cevre toplami veren bir method yazdik
	return 2*r.width + 2*r.height
}


func MethodTest(){
	r := rectangle{width: 10, height: 25}

	fmt.Println("area :" , r.area())
	fmt.Println("perim :", r.perim())

	// go otomatik olarak pointer ve 
	// deger ifadelerini cevirir
	// asagida pointer vermemize ragmen islemi
	// normal bir sekilde dondurdu
	rp := &r
	fmt.Println("area", rp.area())
	fmt.Println("perim", rp.perim())

}