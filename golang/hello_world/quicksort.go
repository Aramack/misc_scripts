package main

import (
  "time"
  "math/rand"
  "fmt"
  "strconv"
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
  //Generate a unsorted array
  unsorted_array := populate_array(10000, 10000)
  
  //Single threaded recursice quick sort
  start_single_thread := time.Now()
  pivot(unsorted_array)
  fmt.Println("Single Thread: " + time.Since(start_single_thread).String())
  
  //Multi threaded go routine quick sort  
  for i := 1; i < 5; i++  {
    start_go_routine := time.Now()
    result_chan := make(chan []int)
    go pivot_go_routine(unsorted_array, result_chan, i)
    <-result_chan
    //sorted := <-result_chan
    //fmt.Println(sorted)
    fmt.Println("Goroutine (Max depth " + strconv.Itoa(i) + "): " + time.Since(start_go_routine).String())
  }  
}

func pivot(unsorted_array []int) []int{
  if (len(unsorted_array) <= 1) {
    return unsorted_array
  }
  
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
 
 func pivot_go_routine(unsorted_array []int, return_chan chan []int, max_goroutine_depth int) {
  max_goroutine_depth = max_goroutine_depth - 1
  if (len(unsorted_array) <= 1) {
    return_chan <- unsorted_array
  }
  
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
  var sorted_less_than = []int{}
  var sorted_greater_than = []int{}
  if (max_goroutine_depth <= 0) {
    sorted_less_than = pivot(less_than_pivot)
    sorted_greater_than = pivot(greater_than_pivot)
  } else {
    less_than_chan := make(chan []int)
	greater_than_chan := make(chan []int)
    go pivot_go_routine(less_than_pivot, less_than_chan, max_goroutine_depth)
	go pivot_go_routine(greater_than_pivot, greater_than_chan, max_goroutine_depth)
    sorted_less_than = <-less_than_chan
	sorted_greater_than = <-greater_than_chan
  }
  sorted_slice := append(sorted_less_than, pivot_number)
  for i := 0; i < len(sorted_greater_than); i++ {
    sorted_slice = append(sorted_slice, sorted_greater_than[i])
  }
  return_chan <- sorted_slice
 }