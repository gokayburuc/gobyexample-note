package conditionals

import (
  "fmt"
)

func SwitchExpressions(){
  i := 1453
  switch i {
  case 1071:
      fmt.Println("Malazgirt Pitched Battle")
  case 1041:
    fmt.Println("Dandanakan Pitched Battle")
  case 1453:
    fmt.Println("Conquer of the Constantinople")
  default:
    fmt.Println("Not a war!")
  }


}
