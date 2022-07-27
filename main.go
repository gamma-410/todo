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
    panic(err)
  } // エラー処理

  switch str {
    case "list":
      data := make([]byte, 1024) 
      // UTF-8 - English: 1byte, Japanese: 3byte
      // data - ファイルに入っている文字列を格納するスライス

      count, err := f.Read(data)
      // f.Read(data) - おそらくこの時に読み込んだデータをdataスライスに渡してる...?
      // count - Read()をした結果、何byte格納されているかがint型で入っている

      if err != nil {
        panic(err)
      } // エラー処理

      fmt.Println(string(data[:count]))
      // この形は Slice[:end] - data の先頭から end -1 までを取得
      // stirng()で囲って変換してやる

    case "add":
      return

    case "delete":
      return

    default:
      fmt.Println("Error: Command not found.")
      return

  }
  
}
