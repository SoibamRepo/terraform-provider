package main

import (
      "fmt"
       "github.com/SoibamRepo/terraform-provider/morestrings"
       "github.com/google/go-cmp/cmp"
      
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,oooooolleH"))
        fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}