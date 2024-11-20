package Dog

import ICat "github.com/goclub/di/internal/dep/icat"

type Dog struct {
	Cat func() ICat.Interface
}

func (dep Dog) Wang() string {
	return "wang!"
}
func (dep Dog) Favorite() string {
	return "cat:" + dep.Cat().Meow()
}
