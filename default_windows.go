package main

import (
	"github.com/aquilax/truncate"
)

var Default = New(
	Replace("/", "?", "<", ">", `\`, ":", "*", "|", `"`, " "),
	StripRange(0x0, 0x1f, 0x80, 0x9f),
	ReplaceRegexp("[. ]+$", "$0_"),
	ReplaceRegexp(`^\.+$`, "$0_"),
	ReplaceRegexp(`(?i)^(con|prn|aux|nul|com[0-9]|lpt[0-9])(\..*)?$`, "$0_"),
	Truncate(255, truncate.DEFAULT_OMISSION, truncate.PositionEnd),
)
