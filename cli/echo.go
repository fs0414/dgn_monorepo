package main

import (
	"fmt"
	"os"
)

func main() {
  arg := os.Args[1:]
  if len(arg) < 0 {
    return
  }
  fmt.Println(arg)
}
