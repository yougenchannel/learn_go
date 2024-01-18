package main

import "fmt"

func main() {
	stringMap := map[string]interface{}{}
	makeM := make(map[int]int)

	makeM[1] = 1

	stringMap["str1"] = 1
	stringMap["str2"] = 2

	value, isPresent := stringMap["str1"]
	if isPresent {
		fmt.Println(value)
	}

	// 删除元素
	delete(stringMap, "str1")

	// 常用技巧, 和if 一同使用

	if value, isPresent := stringMap["str1"]; isPresent {
		fmt.Println(value)
	}

	// 遍历Map
	for key, value := range stringMap {
		fmt.Println("key = ", key, "value = ", value)
	}
}
