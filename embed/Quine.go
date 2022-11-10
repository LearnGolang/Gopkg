/*
0 error(s),0 warning(s)
Team:0e0w Security Team
Author:0e0wTeam[at]gmail.com
Datetime:2022/11/10 7:48
*/

package main

import (
	_ "embed"
	"fmt"
)

//go:embed Quine.go
var src string

func main() {
	fmt.Print(src)
}
