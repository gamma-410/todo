package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	arg := flag.Arg(0)
	switch arg {
	case "init":
		initTodo()
	case "add":
		arg1 := flag.Arg(1)
		add(arg1)
	case "list":
		list()
	case "delete":
		deleted()
	default:
		fmt.Println("error: received command was not found.")
	}
}

func initTodo() {
	file, err := os.Create(".todo")
	// 学び: "" で囲むとstring, '' で囲むとruneになる

	if err != nil {
		panic(err)
	}
	file.Close()
	fmt.Println("success: successfully initialized todo.")
}

func add(arg string) {
	file, err := os.OpenFile(".todo", os.O_WRONLY|os.O_APPEND, 0666)
	// 引数: ファイルのパス, フラグ, パーミッション

	if err != nil {
		panic(err)
	}

	defer file.Close()
	// 漢文で言うレ点? 下の処理が終わったら実行だと思う

	fmt.Fprintln(file, arg)
	// "F" がついてるものは書き込み先を明示的に指定できる fmt.Fprintln(パス, 内容)
}

func list() {
	file, err := os.Open(".todo")

	if err != nil {
		panic(err)
	}

	data := make([]byte, 1024)
	// UTF-8 - English: 1byte, Japanese: 3byte
	// data - ファイルに入っている文字列を格納するスライス

	count, err := file.Read(data)
	// f.Read(data) - おそらくこの時に読み込んだデータをdataスライスに渡してる...?
	// count - Read()をした結果、何byte格納されているかがint型で入っている

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data[:count]))
	// この形は Slice[:end] - data の先頭から end -1 までを取得
}

func deleted() {
	fmt.Println("削除機能は実装予定です。")
}
