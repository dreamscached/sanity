# Sanity

Sanity is a fast and easily extensible file name (and in fact any other string) sanitizer.

## Usage

### Built-in rule set

Sanity provides a sensible default rule set for Windows and Linux file names that will be enough in most cases:

```go
package main

import (
	"fmt"

	"github.com/dreamscached/sanity"
)

func main() {
	// Prints 'con_.txt' if on Windows
	fmt.Println(sanity.Default.Sanitize("con.txt"))

	// Prints 'foobar' if on Linux/Darwin
	fmt.Println(sanity.Default.Sanitize("foo\x00bar.txt"))
}
```

### Creating custom rule sets

Sanity provides a clean and obvious way to create custom rule sets with Rule factory functions. For example, here's
Linux default file name rule set.

```go
package main

import (
	"github.com/dreamscached/sanity"
	"github.com/aquilax/truncate"
)

var Default = sanity.New(
	sanity.Replace("/", " "),
	sanity.StripRune(0x0),
	sanity.Truncate(255, truncate.DEFAULT_OMISSION, truncate.PositionEnd),
)
```