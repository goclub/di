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
