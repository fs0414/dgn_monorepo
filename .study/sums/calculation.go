package main

import (
	"fmt"
	"strconv"
)

func stringCast(res int) string{
  return strconv.Itoa(res)
}

func main() {
 result := 1 + 1
  stringPutRes := stringCast(result)
  fmt.Println(stringPutRes)
}
