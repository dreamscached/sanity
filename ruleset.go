package sanity

// Ruleset represents string sanitization rule set.
type Ruleset []Rule

// New constructs new Ruleset from the provided rules.
func New(rules ...Rule) Ruleset {
	return rules
}

// Extend adds more rules to the existing Ruleset
func (r Ruleset) Extend(rules ...Rule) Ruleset {
	return append(r, rules...)
}

// Copy creates a copy of existing Ruleset.
func (r Ruleset) Copy() Ruleset {
	c := make(Ruleset, len(r))
	copy(c, r)
	return c
}

// Sanitize applies rules in the Ruleset to the input string and returns sanitized string with all rules applied.
func (r Ruleset) Sanitize(in string) string {
	for _, rule := range r {
		in = rule(in)
	}

	return in
}
