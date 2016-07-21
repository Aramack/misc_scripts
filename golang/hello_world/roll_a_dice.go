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
  "bytes"
)

func main() {
  roll_regex, _ := regexp.Compile("roll ([0-9]+)d([0-9]+)")
  for {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("What do you want? ")
    input, _ := reader.ReadString('\n')

    if (roll_regex.MatchString(input)) {
      d := roll_regex.FindStringSubmatch(input)
      fmt.Println("Rolling " + d[1] +"d" + d[2])
      number_of_dice, _ := strconv.Atoi(d[1])
      sides_of_the_dice, _ := strconv.Atoi(d[2])
      var total = 0
      for i := 0; i < number_of_dice; i++ {
        roll := get_number(sides_of_the_dice)
        total += roll
		fmt.Println(roll)
      }
      var buffer bytes.Buffer
      buffer.WriteString("Total: ")
	  buffer.WriteString(strconv.Itoa(total))
      fmt.Println(buffer.String())
    } else if strings.Contains(input, "quit"){
      fmt.Println("bye...")
      os.Exit(0)
    } else {
      fmt.Println("I don't know how to do that")
    }
  }
}

/*
 Get a random number between [1, upper]
*/
func get_number(upper int) int {
  seed := rand.NewSource(time.Now().UnixNano())
  random := rand.New(seed)
  var number = random.Intn(upper) + 1
  return number
}