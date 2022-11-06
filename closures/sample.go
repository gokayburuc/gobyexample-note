package closures

func BaseFunc() func() int {
  i := 0
  return func() int{
    i++ // her calistirmada return degerinini 1 arttirir
    return i
  }
}
