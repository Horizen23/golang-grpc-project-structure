package services

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewGreeterService,
)
