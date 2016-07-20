package main

import (
  "fmt"
  "bufio"
  "os"
  "time"
  "math/rand"
  "regexp"
  "strconv"
  "strings"
)

func main() {
  for {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("What do you want? ")
    input, _ := reader.ReadString('\n')
	
    r, _ := regexp.Compile("roll ([0-9]+)d([0-9]+)")
    if (r.MatchString(input)) {
	  d := r.FindStringSubmatch(input)
	  fmt.Println("Rolling " + d[1] +"d" + d[2])
	  number_of_dice, _ := strconv.Atoi(d[1])
	  sides_of_the_dice, _ := strconv.Atoi(d[2])
	  for i := 0; i < number_of_dice; i++ {
        fmt.Println(get_number(sides_of_the_dice))
	  }
    } else if strings.Contains(input, "quit"){
      fmt.Println("bye...")
	  os.Exit(0)
    } else {
	  fmt.Println("I don't know how to do that")
	}
  }
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