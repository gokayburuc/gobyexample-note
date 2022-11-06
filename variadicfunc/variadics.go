package variadicfunc

import (
  "fmt"
)

func VariadicTest(nums...int){
  fmt.Print(nums," ") // nums degerlerini goster

  total := 0 // toplam degeri sıfır
  // var total = 0
  for _,num := range nums{
    total += num
  }

  fmt.Println("Total :",total)
}
