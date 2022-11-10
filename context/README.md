# [context](https://pkg.go.dev/context)-并发控制

## 01-基础概念

context是goroutine 的上下文，包含 goroutine 的运行状态、环境、现场等信息。context 主要用来在 goroutine 之间传递上下文信息，包括：取消信号、超时时间、截止时间、k-v 等。Go 1.7 标准库引入context。

## 02-使用例子

- https://github.com/jamestrandung/go-context
- https://github.com/aarie33/go-context

## 03-参考资源

- https://zhuanlan.zhihu.com/p/68792989
- https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-context/