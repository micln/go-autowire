// Code generated by go-autowire. DO NOT EDIT.

// +build wireinject

package autowire_gen

import (
	"github.com/Just-maple/go-autowire/example/dependencies"
	"github.com/google/wire"
)

func InitializeTest() (*dependencies.Test, func(), error) {
	panic(wire.Build(Sets))
}