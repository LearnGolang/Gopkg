# [reflect](https://pkg.go.dev/reflect)-反射

reflect包实现运行时反射，允许程序操作任意类型的对象。典型的用法是取一个静态类型 interface{} 的值，通过调用 TypeOf 提取其动态类型信息，返回一个 Type。