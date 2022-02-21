package main

import (
  "flag"
	"fmt"
  "strings"
)

func main() {
	
  /* User input text to parse */
  text := flag.String("text", "", "Text to parse.")
  flag.Parse()

  /* calculate count */
  count := strings.Count(*text, "")
  fmt.Printf("Your string has %d chars", count)

  
}
