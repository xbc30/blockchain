## GO

### 为并发而生
```text
Goroutine 非常轻量，主要体现在以下两个方面：

上下文切换代价小： Goroutine 上下文切换只涉及到三个寄存器（PC / SP / DX）的值修改；而对比线程的上下文切换则需要涉及模式切换（从用户态切换到内核态）、以及 16 个寄存器、PC、SP…等寄存器的刷新；

内存占用少：线程栈空间通常是 2M，Goroutine 栈空间最小 2K；

Golang 程序中可以轻松支持10w 级别的 Goroutine 运行，而线程数量达到 1k 时，内存占用就已经达到 2G。
```

### code
> go实战部分

* consensus go实现各类共识算法
* internal 官方内置包
* micro go微服务相关工具实战
* type go数据结构类型
* web gin/beego

### frame
> go实用工具和依赖库

### list
> go常见知识点

### point
> 容易混淆的点