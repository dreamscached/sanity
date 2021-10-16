package main

import (
	"github.com/aquilax/truncate"
)

var Default = New(
	Replace("/", " "),
	StripRune(0x0),
	Truncate(255, truncate.DEFAULT_OMISSION, truncate.PositionEnd),
)
