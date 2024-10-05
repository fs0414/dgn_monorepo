package main

import (
	"fmt"
	"runtime"
)

func printStack() {
  buf := make([]byte, 1024)
  n := runtime.Stack(buf, false)
  //fmt.Printf("Stack trace:\n%s\n", buf[:n])
  fmt.Printf(string(buf[:n]))
}

func quickSort(arr []int) []int {
  //printStack()

  int con := 1

  if len(arr) < 2 {
    return arr
  }

  pivot := arr[0]

  var less []int
  var greater []int

  for _, v := range arr[1:] {
    fmt.Println("pivot:", pivot)
    fmt.Println("v:", v)
    if v <= pivot {
      less = append(less, v)
    } else {
      greater = append(greater, v)
    }
  }
  fmt.Println("less:", less)
  fmt.Println("geater:", greater)
  return append(append(quickSort(less), pivot), quickSort(greater)...)
}

func callStackRun(arg int) int {
  if arg < 2 {
    return 1
  }
  fmt.Println("arg:", arg)
  return callStackRun(arg - 1)
}

func main() {
  arr := []int{33, 10, 55, 7, 89, 21, 45}
  fmt.Println("before:", arr)
  result := quickSort(arr)
  //result := callStackRun(5)
  fmt.Println("after:", result)
}
