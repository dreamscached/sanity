package main

type Ruleset []Rule

func New(rules ...Rule) Ruleset {
	return rules
}

func (r *Ruleset) Extend(rules ...Rule) {
	*r = append(*r, rules...)
}

func (r Ruleset) Copy() Ruleset {
	c := make(Ruleset, len(r))
	copy(c, r)
	return c
}

func (r Ruleset) Sanitize(in string) string {
	for _, rule := range r {
		in = rule(in)
	}

	return in
}
