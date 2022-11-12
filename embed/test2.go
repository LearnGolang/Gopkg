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

//go:embed fingerprint.txt
var ff embed.FS

func main() {
	data, _ := ff.ReadFile("fingerprint.txt")
	fmt.Println(string(data))
}
