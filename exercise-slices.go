package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
  a := make([][]uint8, dy)
  
  for x := range a {
  	a[x] = make([]uint8, dx)
	for y := range a[x] {
		a[x][y] = uint8((x+y)/2)
	}
  }
  return a
}


func main() {
	pic.Show(Pic)
}