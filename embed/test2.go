/*
0 error(s),0 warning(s)
Team:0e0w Security Team
Author:0e0wTeam[at]gmail.com
Datetime:2022/11/10 7:40
*/

package main

import (
	"embed"
	"fmt"
)

//go:embed pocs
var ftest2 embed.FS

func main() {
	data, _ := ftest2.ReadFile("pocs/test1.yaml")
	fmt.Println(string(data))
	data, _ = ftest2.ReadFile("pocs/test2.yaml")
	fmt.Println(string(data))
}
