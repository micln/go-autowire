// Code generated by go-autowire. DO NOT EDIT.

package example_zoo

import (
	"github.com/google/wire"
)

var ZooSet = wire.NewSet(
	wire.Struct(new(Zoo), "*"),
)