package Cat

import (
	IDog "github.com/goclub/di/internal/dep/idog"
)

type Cat struct {
	Dog func() IDog.Interface
}

func (dep Cat) Meow() string {
	return "meow~"
}
func (dep Cat) Favorite() string {
	return "dog:" + dep.Dog().Wang()
}
