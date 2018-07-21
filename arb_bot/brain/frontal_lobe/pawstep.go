package frontal_lobe

struct pawStep type {
	url string // a fully qualified url
	expectations []expectation // a list of exepctations, all of which must be fulfilled before the step can be considered successful
}

func (ps *pawStep)Url() {
	return ps.url
}