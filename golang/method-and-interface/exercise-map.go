package main

import (
  "strings"
  "golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
  slice := strings.Fields(s)
  res := make(map[string]int)

  for _, v := range slice[:] {
    res[v]++
  }

  return res
}

func main() {
  wc.Test(WordCount)
}
