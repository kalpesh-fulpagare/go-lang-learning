package main

import "fmt"

func add(a int, b int) int {
  return a + b
}

func returnStatementDemo()(s string) {
  s = "Kalpesh will be returned from this function"
  return
}

func main() {
  fmt.Println(add(50, 40))
  fmt.Println(returnStatementDemo())
}

