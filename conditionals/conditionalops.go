package conditionals

import (
  "fmt"
)

func  Test(){
  if 7%2 == 0 {
    fmt.Println("7 is odd")
  }else{
    fmt.Println("7 is even")
  }

  if num:= 9 ; num <0{
    fmt.Println(num, "is negative")
  } else if num < 10{
    fmt.Println(num, "has 1 digit") //
  }else {
    fmt.Println(num, "has multiple digits")
  }
}
