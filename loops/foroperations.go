package loops

import (
  "fmt"
)

func CheckLoops(){
  for {
    fmt.Println("Hello World!")
    fmt.Println("Bye Bye World")
    break
    fmt.Println("Broken!") // döngü kırıldığı için bu ifade ekranda gözükmeyecektir
  }

  for i := 5 ; i <= 100 ; i++{
    fmt.Println(float64(i) / 100)
  }

  for n := 0; n <= 5; n++ {
       if n%2 == 0 {
         fmt.Printf("2 ile tam bolunen %v degeri atlandi\n",n) // opsiyonel olarak yazilmayabilir
           continue // buradaki islemi es gec ve islemine devam et
       }
       fmt.Println(n)
   }
}
