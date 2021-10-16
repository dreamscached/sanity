package main

type Sanitizer []Rule

func New(rules ...Rule) Sanitizer {
	return rules
}

func (s *Sanitizer) Extend(rules ...Rule) {
	*s = append(*s, rules...)
}

func (s Sanitizer) Sanitize(in string) string {
	for _, rule := range s {
		in = rule(in)
	}

	return in
}
