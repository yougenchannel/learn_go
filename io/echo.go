package main

import (
	"flag"
	"os"
)

var NewLine = flag.Bool("n", false, "print newline")

const (
	Space   = " "
	Newline = "\n"
)

func main() {

	flag.PrintDefaults() // 打印flag使用帮助
	flag.Parse()         // 读取并解析命令行参数
	var s string = ""

	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
		}
		if *NewLine {
			s += Newline
		}

		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)

}
