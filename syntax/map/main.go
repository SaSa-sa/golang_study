package main

import (
	"fmt"
	"sort"
)


func main() {
	// mapが参照型である事の確認
	m := make(map[string]string)
	// main関数内でmは変更していないが参照渡しされるので変更される。
	addMap(m)
	fmt.Println(m)
	// copyしているので、mに影響はない
	addMapWithCopy(m)
	fmt.Println(m)

	// ---------
	// mapに対するfor  ==> 順序性が担保されない(入れた順に印刷されない)
	// 学籍番号と学生名のMap
	studentIDMap := map[int]string{
		3: "田中",
		1: "伊藤",
		2: "佐藤",
		4: "佐々木",
	}

	sum := Students{}
	for k, v := range studentIDMap {
		s := Student{k,v} //vにidとnameどっちも入ってる。変更必要
		sum = append(sum,s)
		//// fmt.Printfでフォーマットに従った文字列を標準出力に出せる
		//fmt.Printf("Name of StudentID:%d is %s\n", k, v)
	}
	sort.Slice(sum,func(i,j int) bool { return sum[i].id < sum[j].id })
	for k, v := range sum {
		//// fmt.Printfでフォーマットに従った文字列を標準出力に出せる
		fmt.Printf("Name of StudentID:%d is %s\n", k, v)
	}
}

func addMap(m map[string]string) {
	m["a"] = "あ"
}
func addMapWithCopy(m map[string]string) {
	copied := make(map[string]string)
	for k, v := range m {
		copied[k] = v
	}
	copied["i"] = "い"
}

type Student struct {
	id    int
	name  string
}
type Students []Student