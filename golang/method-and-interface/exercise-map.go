package main

import (
  "strings"
  "golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
  slice := strings.Split(s, " ")
  res := make(map[string]int)

  for _, v := range slice[:] {
    c, ok := res[v]
    if ok {
      res[v] = c + 1
    } else {
      res[v] = 1
    }
  }

  return res
}

func main() {
  wc.Test(WordCount)
}
