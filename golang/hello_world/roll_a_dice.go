package main

import "fmt"
import "time"
import "math/rand"

func main() {
  var sides_of_the_dice = 10
  fmt.Println(get_number(sides_of_the_dice))
}

/*
 Get a random number between [0, x)
*/
func get_number(upper int) int {
  seed := rand.NewSource(time.Now().UnixNano())
  random := rand.New(seed)
  var number = random.Intn(upper)
  return number
}