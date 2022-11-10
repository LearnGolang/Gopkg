/*
0 error(s),0 warning(s)
Team:0e0w Security Team
Author:0e0wTeam[at]gmail.com
Datetime:2022/11/9 20:43
*/

package main

import (
	"embed"
	"log"
	"net/http"
)

// 将README.md文件打包到可执行程序中

//go:embed README.md
var content embed.FS

func main() {
	mutex := http.NewServeMux()
	mutex.Handle("/", http.FileServer(http.FS(content)))
	err := http.ListenAndServe(":8080", mutex)
	if err != nil {
		log.Fatal(err)
	}
}
