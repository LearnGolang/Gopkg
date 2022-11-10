/*
0 error(s),0 warning(s)
Team:0e0w Security Team
Author:0e0wTeam[at]gmail.com
Datetime:2022/11/10 7:17
*/

package main

import (
	"embed"
)

//go:embed hello.txt
var s string
var b []byte
var f embed.FS

func main() {
	E1()
	E2()
	E3()
}

// E1 将一个文件嵌入到字符串中：
func E1() {
	print(s)
}

// E2 将一个文件嵌入到一片字节中：
func E2() {
	print(string(b))
}

// E3 将一个或多个文件嵌入到文件系统中
func E3() {
	data, _ := f.ReadFile("hello.txt")
	print(string(data))
}
