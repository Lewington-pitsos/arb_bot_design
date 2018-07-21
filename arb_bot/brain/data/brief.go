package frontal_lobe

type brief struct {
	Success: bool // indivates whether the scrape or step in question was successful
	reports []report // all the information returned by the scrape
}

func (b *brief)AddReport(report) {
	b.reports = append(b.reports, report)
}

func (b *brief)Reports() []report {
	return b.reports
}

func (b *brief)MergeBrief(otherBrief *brief) {
	b.success = b.success && otherBrief.success
	b.reports = append(b.reports, otherBrief.Reports()...)
}
