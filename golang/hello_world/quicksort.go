package main

import (
  "time"
  "math/rand"
  "fmt"
)

func populate_array(size int, upper int) []int{
  seed := rand.NewSource(time.Now().UnixNano())
  random := rand.New(seed)
  //unsorted_array []int := buffer[0:upper]
  unsorted_array := make([]int, size)
  for i := 0; i < size; i++ {
    unsorted_array[i] = random.Intn(upper)
  }
  return unsorted_array
}

func main(){
  unsorted_array := populate_array(10, 10)
  fmt.Println(unsorted_array);
  fmt.Println(pivot(unsorted_array))
}

func pivot(unsorted_array []int) []int{
  if (len(unsorted_array) <= 1) {
    return unsorted_array
  }
  //return append(unsorted_array, 1)
  
  pivot_number := unsorted_array[len(unsorted_array)-1]
  var less_than_pivot = []int{}
  var greater_than_pivot = []int{}
  for i := 0; i < (len(unsorted_array) - 1); i++ {
    if (unsorted_array[i] <= pivot_number) {
      less_than_pivot = append(less_than_pivot, unsorted_array[i])
    } else {
      greater_than_pivot = append(greater_than_pivot, unsorted_array[i])
    }  
  }
  sorted_less_than := pivot(less_than_pivot)
  sorted_greater_than := pivot(greater_than_pivot)
  sorted_slice := append(sorted_less_than, pivot_number)
  for i := 0; i < len(sorted_greater_than); i++ {
    sorted_slice = append(sorted_slice, sorted_greater_than[i])
  }
  return sorted_slice
 }