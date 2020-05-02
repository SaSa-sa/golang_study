package main

import (
	"errors"
	"fmt"
	"strconv"

	gerrors "github.com/pkg/errors"
)

func main() {
	// エラーハンドリングの例
	// (error以外の型, error)が返ってくるので、両方を受け取りifでハンドリングする。
	iStr := "1"
	if i, err := strconv.Atoi(iStr); err != nil {
		fmt.Printf("Atoi of i failed. %v\n", err)
	} else {
		fmt.Printf("i is %d\n", i)
	}

	jStr := "a"
	if j, err := strconv.Atoi(jStr); err != nil {
		fmt.Printf("Atoi of j failed. %v\n", err)
	} else {
		fmt.Printf("j is %d\n", j)
	}

	// Goの標準パッケージ。↑のerrを変えたい時
	err := errors.New("error occurred")
	fmt.Printf("%+v\n\n", err)

	// github.com/pkg/errors
	// 以下の様なスタックトレースが出せる(どこでエラーが起きたか分かる)ので、こちらをよく使う。
	// main.main
	// 	/Users/sakura.yuto/go/src/github.com/mf-sakura/golang_study/exception/error/main.go:15
	// runtime.main
	// 	/Users/sakura.yuto/.goenv/versions/1.14.0/src/runtime/proc.go:203
	// runtime.goexit
	// 	/Users/sakura.yuto/.goenv/versions/1.14.0/src/runtime/asm_amd64.s:1373
	gErr := gerrors.New("error occurred")
	// スタックトレースを出すには、`%+v`を使う
	fmt.Printf("%+v\n\n", gErr) //-> 'error occurred"'
	// WithStackでStackを追加出来る
	stackErr := gerrors.WithStack(gErr) // .Withstach => ファイル元情報追加
	fmt.Printf("%+v\n\n", stackErr)
	// Wrap, WrapfでStackを追加しながら、エラーメッセージの追加が出来る
	wrappedErr := gerrors.Wrap(stackErr, "wrap") // .Wrap => 任意のエラー文を追加
	fmt.Printf("%+v\n\n", wrappedErr)            // => %+vを指定するとファイル場所ができる
	fmt.Println(wrappedErr)                      // => これだとエラー文だけ

	// 自分で定義したエラー
	// 特定の条件で使うエラーを定義しておき、そのエラーの場合にのみ処理する事が可能になる。

	var myErr error
	myErr = MyError{}                 // => 構造体を代入(MyError{}がerror型として正しくないとここでエラー)
	if _, ok := myErr.(MyError); ok { //キャストしている(myErrが、MyError型として正しいか判定)
		fmt.Println("Variable is MyError")
	} else {
		fmt.Println("Variable isn't MyError")
	}

	if _, ok := gErr.(MyError); ok { // => これはfalseになる
		fmt.Println("Type of gErr is MyError")
	} else {
		fmt.Println("Type of gErr isn't MyError")
	}
}

type MyError struct{}

// Error()メソッドがあるので、MyErrorはerrorになる。
func (e MyError) Error() string {
	return "My Error Occurred"
}
