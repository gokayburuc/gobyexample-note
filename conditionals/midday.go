package conditionals

import (
  "time"
  "fmt"
)

func Midday(){
  t := time.Now() // sistem zamanını getir
  fmt.Println()
  switch {
  case t.Hour() < 12 :
    fmt.Println("Morning")
  case t.Hour() > 12 :
    fmt.Println("its ", t.Hour())
    fmt.Println("Afternoon")
  }
}
