## go 面试 1

### go 语言什么时候会发生内存逃逸

在 go 语言中，内存逃逸是指变量的内存分配从栈转移到堆上的现象。

内存逃逸会增加垃圾回收的负担，影响程序的性能。在编写 Go 代码时，应尽量避免不必要的内存逃逸，可以通过 go build -gcflags="-m"命令来分析代码中的内存逃逸情况。

**函数返回局部变量的指针**
当函数返回一个局部变量的指针时，该变量会发生内存逃逸。因为在函数执行结束后，栈上的局部变量会被销毁，
如果返回其指针，外部代码还需要使用这个变量，所以 go 编译器会将该变量分配到堆上。

```go
package main

func escapeExample() *int {
    num := 10
    return &num // 局部变量 num 发生内存逃逸
}

func main() {
    result := escapeExample()
    _ = result
}
```

**变量大小在编译时无法确定**
如果变量的大小在编译时无法确定，go 编译器会将其分配到堆上。例如，使用切片时，如果切片的长度或容量在编译时无法确定，
就可能发生内存逃逸。

```go
package main

func dynamicSlice(size int) []int {
    // 切片的大小在编译时无法确定，会发生内存逃逸
    slice := make([]int, size)
    return slice
}

func main() {
    result := dynamicSlice(10)
    _ = result
}
```

**栈空间不足**
如果变量所需的栈空间超过了栈的最大限制，go 编译器会将其分配到堆上。不过，这种情况相对较少，因为
go 的栈空间会根据需要动态增长。

**闭包引用外部函数的局部变量**
当闭包引用了外部函数的局部变量时，这些局部变量会发生内存逃逸。一 in 为 i 闭包可能会在外部函数执行结束后
继续存在并访问这些变量，所以需要将这些变量分配到堆上。

```go
package main

func closureExample() func() int {
    num := 10
    // 闭包引用了外部函数的局部变量 num，num 发生内存逃逸
    return func() int {
        return num
    }
}

func main() {
    closure := closureExample()
    _ = closure()
}
```

**向接口类型变量赋值**
将具体类型的值赋给接口类型的变量时，也可能会发生内存逃逸。
因为接口是动态类型，需要在运行时确定具体的值，所以 go 编译器会将值分配到堆上。

```go
package main

type MyInterface interface {
    Print()
}

type MyStruct struct{}

func (m MyStruct) Print() {}

func interfaceExample() {
    s := MyStruct{}
    var i MyInterface = s // 变量 s 发生内存逃逸
    _ = i
}

func main() {
    interfaceExample()
}
```
