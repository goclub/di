package main

import (
	Cat "github.com/goclub/di/internal/dep/cat"
	Dog "github.com/goclub/di/internal/dep/dog"
	ICat "github.com/goclub/di/internal/dep/icat"
	IDog "github.com/goclub/di/internal/dep/idog"
	"log"
)

func main() {
	// 声明区
	var cat ICat.Interface
	var dog IDog.Interface

	// 创建区
	cat = Cat.Cat{
		func() IDog.Interface {
			return dog
		},
	}
	dog = Dog.Dog{
		func() ICat.Interface {
			return cat
		},
	}
	log.Println("cat.Meow()", cat.Meow())         // meow~
	log.Println("cat.Favorite()", cat.Favorite()) // dog:wang!
	log.Println("dog.Wang()", dog.Wang())         // wang!
	log.Println("dog.Favorite()", dog.Favorite()) // cat:meow~
}

/*
猫和狗互相依赖, 猫会发嗲 `cat.Meow()`,狗会发财 `dog.Wang()`

猫的最爱是狗叫 `cat.Favorite()`  它内部调用了 `dog.Wang()`

狗的最爱是猫叫 `dog.Favorite()`  它内部调用了 `cat.Meow()`

猫狗事实上循环依赖了

但是 `go` 不允许循环依赖, **go比狗还狗**

猫狗利用*接口*描述了各自的*功能*

突破了 go 的限制,一辈子在一起❤️🎉🎉🎉❤️
*/
