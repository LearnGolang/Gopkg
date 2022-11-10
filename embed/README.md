# [embed](https://pkg.go.dev/embed)-打包静态资源

## 01-基础概念

embed包提供对嵌入在运行的 Go 程序中的文件的访问。这是Go 1.16中新增的特性。可以利用此包将静态资源文件打包到exe程序中。

导入“embed”的 Go 源文件可以使用 //go:embed 指令在编译时使用从包目录或子目录中读取的文件内容来初始化字符串、[]byte 或 FS 类型的变量。

如果要嵌入文件，变量的类型得是 string 或者 []byte。如果要嵌入一组文件，变量的类型得是embed.FS

## 02-使用例子

- https://github.com/carlmjohnson/exembed
- https://github.com/soulteary/awesome-golang-embed
- https://github.com/shadow1ng/fscan
- https://github.com/zan8in/afrog
- https://github.com/veo/vscan
- https://github.com/j5s/FscanX
- https://github.com/sairson/Yasso
- https://github.com/hktalent/scan4all
- https://github.com/404tk/lazyscan
- https://github.com/zyylhn/zscan-poc-check
- https://github.com/ExpLangcn/Aopo
- https://github.com/IceMoon1995/jray
- https://github.com/whitesource/spring4shell-detect

## 03-参考资源

- https://github.com/golang/go/issues/35950
- https://go.googlesource.com/proposal/+/master/design/draft-embed.md
- https://blog.carlmjohnson.net/post/2021/how-to-use-go-embed/
- https://colobu.com/2021/01/17/go-embed-tutorial/
- https://soulteary.com/2022/01/15/explain-the-golang-resource-embedding-solution-part-1.html
- https://github.red/go-embed