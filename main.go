// todo | @gamma-410

package main

import (
  "fmt"
  "os"
  "flag"
)

func main() {
  flag.Parse() // 引数を解析
  str := flag.Arg(0) 

  f, err := os.Open(str)

  if err != nil {
    fmt.Println("えらー")
  }

  if str == "" {
    fmt.Println("Error: cannot find filename.")
  } else {
    data := make([]byte, 1024)
    count, err := f.Read(data)
    
    if err != nil {
      fmt.Println("えらー")
    }
    fmt.Println(string(data[:count]))
  }

}
