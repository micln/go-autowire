// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package init_zoo

import (
	"github.com/Just-maple/go-autowire/example_zoo"
)

// Injectors from wire.go:

func InitZoo() example_zoo.Zoo {
	cat := example_zoo.Cat{}
	dog := example_zoo.Dog{}
	bird := &example_zoo.Bird{}
	zoo := example_zoo.Zoo{
		Cat:       cat,
		Dog:       dog,
		FlyAnimal: bird,
	}
	return zoo
}