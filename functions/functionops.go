package functions

import ("fmt")

func Plus(a int, b int) int{
  return a + b
}

func plusPlus(a,b,c int ) int{
  return a + b + c
}


func FunctionTransactions(){
  res := Plus(25,35)
  fmt.Println(res)

  newres := plusPlus(12,13,14)
  fmt.Println(newres)
}
