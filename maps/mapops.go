package maps

import (
  "fmt"
)

func MapOperations()  {
  // make ile map tanimlama
  myMap := make(map[string]int)

  myMap["key1"] = 18
  myMap["key2"] = 22

  fmt.Println("map:" , myMap)

  //direkt olarak map tanimlama
  newMap := map[string]int{
    "edirne":22,
    "bursa":16,
  }
  
  fmt.Println(newMap)
}
