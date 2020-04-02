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
	// mapに対するfor
	// 学籍番号と学生名のMap
	studentIDMap := map[int]string{
		3: "田中",
		1: "伊藤",
		2: "佐藤",
		4: "佐々木",
	}

	// mapのキーだけ取り出す
	var keys []int
	for k := range studentIDMap {
		keys = append(keys, k)
	}

	//keyを昇順でソート
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	// ※sort.Ints(keys)のが早い。けど、降順にする時はsort.Slice

	for _, k := range keys {
		fmt.Printf("Name of StudentID:%d is %s\n", k, studentIDMap[k])
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
