package rpci

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewGreeterServer,
)
