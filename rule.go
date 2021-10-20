package sanity

import (
	"regexp"
	"strings"

	"github.com/aquilax/truncate"
)

// Rule represents a file sanitization rule, a function that turns (potentially)
// invalid input string into valid output.
type Rule func(in string) (out string)

// Truncate truncates string to the specified length.
func Truncate(length int, omission string, pos truncate.TruncatePosition) Rule {
	return func(in string) string {
		return truncate.Truncate(in, length, omission, pos)
	}
}

// Replace replaces met substrings (all varargs except the last one) with a replacement string (last vararg.)
// It panics if there were less than two arguments provided.
func Replace(substringsAndRepl ...string) Rule {
	if len(substringsAndRepl) < 2 {
		panic("rule.Replace: less than two arguments")
	}

	lm1 := len(substringsAndRepl) - 1
	repl := substringsAndRepl[lm1]
	str := substringsAndRepl[:lm1]

	return func(in string) string {
		for _, s := range str {
			in = strings.ReplaceAll(in, s, repl)
		}

		return in
	}
}

// Strip is a convenience call to Replace with last argument being an empty string.
// It panics if there was less than one argument provided.
func Strip(substrings ...string) Rule {
	return Replace(append(substrings, "")...)
}

// ReplaceRegexp replaces matched patterns (all varargs except the last one) with a replacement string (last vararg.)
// It panics if there were less than two arguments provided.
func ReplaceRegexp(patternsAndRepl ...string) Rule {
	if len(patternsAndRepl) < 2 {
		panic("rule.ReplaceRegexp: less than two arguments")
	}

	lm1 := len(patternsAndRepl) - 1
	repl := patternsAndRepl[lm1]

	pats := make([]*regexp.Regexp, lm1)
	for i := 0; i < lm1; i++ {
		var err error
		pats[i], err = regexp.Compile(patternsAndRepl[i])
		if err != nil {
			panic("rule.ReplaceRegexp: invalid pattern: " + err.Error())
		}
	}

	return func(in string) string {
		for _, re := range pats {
			in = re.ReplaceAllString(in, repl)
		}

		return in
	}
}

// StripRegexp is a convenience call to ReplaceRegexp with last argument being an empty string.
func StripRegexp(patterns ...string) Rule {
	return ReplaceRegexp(append(patterns, "")...)
}

type nRange struct {
	from, to int32
}

func (r nRange) includes(n int32) bool { return n >= r.from && n <= r.to }

type nRangeSlice []nRange

func (s nRangeSlice) includes(n int32) bool {
	for _, r := range s {
		if r.includes(n) {
			return true
		}
	}

	return false
}

// StripRange removes runes in the provided ranges (defined as paired varargs min-max) from the input string.
// It panics if odd amount of arguments was provided.
func StripRange(ranges ...int32) Rule {
	l := len(ranges)
	if l%2 > 0 {
		panic("rule.StripRange: odd amount of arguments")
	}

	ld2 := l / 2
	nrs := make(nRangeSlice, ld2)
	for i := 0; i < ld2; i++ {
		j := i * 2
		nrs[i] = nRange{ranges[j], ranges[j+1]}
	}

	return func(in string) string {
		b := &strings.Builder{}
		for _, c := range in {
			if !nrs.includes(c) {
				b.WriteRune(c)
			}
		}

		return b.String()
	}
}

// StripRune removes the provided runes from the input string.
func StripRune(runes ...rune) Rule {
	rm := make(map[rune]struct{}, len(runes))
	for _, r := range runes {
		rm[r] = struct{}{}
	}

	return func(in string) string {
		b := &strings.Builder{}
		for _, c := range in {
			if _, ok := rm[c]; !ok {
				b.WriteRune(c)
			}
		}

		return b.String()
	}
}
