package slices

import (
  "fmt"
)

func GetSlice()  {
  mySlice := make([]string, 3)
  fmt.Println("empty slice : ", mySlice) // empty slice :  [  ]

  // atamalar
  mySlice[0] = "Gokay"
  mySlice[1] = "Esin"
  mySlice[2] = "Berk"
  // mySlice[3] = "Ahmet"

  fmt.Println(mySlice)
  fmt.Println("mySlice[1]:", mySlice[1])
  fmt.Println(len(mySlice))

  //append
  mySlice = append(mySlice, "Didem")
  fmt.Println(mySlice)

  //copy
  sehirler := make([]string, 3)
  sehirler[0] = "Edirne"
  sehirler[1] = "Bursa"
  sehirler[2] = "Denizli"

  fmt.Println(sehirler)
  yeni := make([]string, 3)
  copy(yeni, sehirler ) // kopya olusturulacak yer , kopyalanan deger
  fmt.Println(yeni)


  // slice
  ilkiki := sehirler[:2]
  fmt.Println(ilkiki)

  deneme := sehirler[1:2]
  fmt.Println(deneme)

  // deneme2 := sehirler[-1]
  // fmt.Println(deneme2)

  // 2d slice -array
  twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
  }
