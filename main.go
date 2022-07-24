// todo | @gamma-410

package main

import (
  "fmt"
  "os"
  "flag"
)

func main() {
  flag.Parse() 
  str := flag.Arg(0) 
  
  f, err := os.Open("todo.txt")

  if err != nil {
    fmt.Println("Error: Cannot find filename.")
  }

  switch str {
    case "list":
      data := make([]byte, 1024)
      count, err := f.Read(data)
    
      if err != nil {
        fmt.Println("Error: Failed to read the file.")
      }
      fmt.Println(string(data[:count]))

    case "add":
      return

    case "delete":
      return

    default:
      fmt.Println("Error: No arguments are specified.")
      return

  }

}
