package testing

import (
	"fmt"
	"testing"
)

func TestAbc(t *testing.T) {

	fmt.Printf("invoke testAbcde")

	// 打印日志到错误日志中
	t.Log("this is error log in test file")
	// 通知函数测试失败， 然后继续执行其他测试
	t.Failed()
	// 通知函数测试失败，文件中其他测试也被跳过， 继续执行其他测试文件
	t.FailNow()
	// 先打印日志，然后失败， 继续执行其他测试文件
	t.Fatal("this is fatal function")
}
