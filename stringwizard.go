package main

import (
  "flag"
	"fmt"
)

func main() {
	
  /* User input text to parse */
  textPtr := flag.String("text", "", "Text to parse.")

  /* User input metric Ex. 'chars' in the input text*/
  metricPtr := flag.String("metric", "chars", "Metric {chars|words|lines};.")

  /* */
    uniquePtr := flag.Bool("unique", false, "Measure unique values of a metric.")

  /* Make sense of complete user input */
    flag.Parse()

  /* Output */
    fmt.Printf("textPtr: %s, metricPtr: %s, uniquePtr: %t\n", *textPtr, *metricPtr, *uniquePtr)

  
}
