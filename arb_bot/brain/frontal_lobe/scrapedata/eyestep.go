package scrapedata

type eyeStep struct {
	url         string // a fully qualified url
	assignments []assignment
}

func (ps *pawStep) Url() {
	return ps.url
}
