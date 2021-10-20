//go:build windows
// +build windows

package filename

import (
	"strings"
	"unicode"

	"github.com/aquilax/truncate"
	"github.com/dreamscached/sanity"
)

var ruleset = sanity.New(
	sanity.Replace("/", "?", "<", ">", `\`, ":", "*", "|", `"`, " "),
	sanity.StripRange(0x0, 0x1f, 0x80, 0x9f),
	sanity.ReplaceRegexp("[. ]+$", "$0_"),
	sanity.ReplaceRegexp(`^\.+$`, "$0_"),
	replaceDevices(),
	sanity.Truncate(255, truncate.DEFAULT_OMISSION, truncate.PositionEnd),
)

func replaceDevices() sanity.Rule {
	return func(in string) string {
		parts := strings.SplitN(in, ".", 2)
		name := strings.ToLower(parts[0])

		l := len(name)
		if l < 3 || l > 4 {
			return in
		}

		if l == 3 {
			switch name {
			case "con", "prn", "aux", "nul":
			default:
				return in
			}
		} else /* if l == 4 */ {
			switch name[:3] {
			case "com", "lpt":
				if !unicode.IsNumber(rune(name[3])) {
					return in
				}
			default:
				return in
			}
		}

		if len(parts) == 1 {
			return parts[0] + "_"
		} else /* if len(parts) == 2 */ {
			return parts[0] + "_." + parts[1]
		}
	}
}
