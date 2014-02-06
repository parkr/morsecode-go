package main

import (
  "fmt"
  "os"
)

func main() {
  args := os.Args[1:];
  fmt.Printf("%d\n", len(args));

  for i := 0; i<len(args); i++ {
    fmt.Printf("%s\n", args[i]);
  }
}
