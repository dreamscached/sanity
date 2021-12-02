package filename

import (
	"github.com/aquilax/truncate"
	"github.com/dreamscached/sanity"
)

var Unix = sanity.New(
	sanity.Replace("/", " "),
	sanity.StripRune(0x0),
	sanity.Truncate(255, truncate.DEFAULT_OMISSION, truncate.PositionEnd),
)
