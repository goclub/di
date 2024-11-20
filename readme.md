# 依赖注入

## 循环依赖&猫狗爱情


[代码实现](./internal/dep/main.go)

猫和狗互相依赖, 猫会发嗲 `cat.Meow()`,狗会发财 `dog.Wang()`

猫最爱的是狗叫 `cat.Favorite()`  它内部调用了 `dog.Wang()`

狗最高的是猫叫 `dog.Favorite()`  它内部调用了 `cat.Meow()`

猫狗事实上循环依赖了

但是 `go` 不允许循环依赖, **go比狗还狗**

猫狗利用*接口*描述了各自的*功能*

```go
// ICat.Interface
// 猫叫和最爱
type Interface interface {
    Meow() string
    Favorite() string
}
// IDog.Interface
// 狗叫和最爱
type Interface interface {
    Wang() string
    Favorite() string
}
// 并在内部功能调各自的接口
type Cat struct {
    Dog func() IDog.Interface
}
type Dog struct {
    Cat func() ICat.Interface
}
```
突破了 go 的限制,一辈子在一起❤️🎉🎉🎉❤️