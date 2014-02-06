package main

import (
  "fmt"
  "os"
)

func rosettaStone() map[string]string {
  return map[string]string {
    "a": ". _",
    "b": "_ . . .",
    "c": "_ . _ .",
    "d": "_ . .",
    "e": ".",
    "f": ". . _ .",
    "g": "_ _ .",
    "h": ". . . .",
    "i": ". .",
    "j": ". _ _ _",
    "k": "_ . _",
    "l": ". _ . .",
    "m": "_ _",
    "n": "_ .",
    "o": "_ _ _",
    "p": ". _ _ .",
    "q": "_ _ . _",
    "r": ". _ .",
    "s": ". . .",
    "t": "_",
    "u": ". . _",
    "v": ". . . _",
    "w": ". _ _",
    "x": "_ . . _",
    "y": "_ . _ _",
    "z": "_ _ . .",
    "0": "_ _ _ _ _",
    "1": ". _ _ _ _",
    "2": ". . _ _ _",
    "3": ". . . _ _",
    "4": ". . . . _",
    "5": ". . . . .",
    "6": "_ . . . .",
    "7": "_ _ . . .",
    "8": "_ _ _ . .",
    "9": "_ _ _ _ ."}
}

func Convert(input string) string {
  morsecode := []byte
  stone     := rosettaStone()
  for _, c  := range x {
    append(morsecode, stone[c])
  }
  return string(morsecode[:])
}

func main() {
  args := os.Args[1:]
  fmt.Printf("%d\n", len(args))

  for i := 0; i<len(args); i++ {
    fmt.Printf("%s\n", args[i])
    fmt.Printf("%s\n", Convert(args[i]))
  }
}
