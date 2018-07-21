package scrapedata

type expectation struct {
	selector string // an xpath selector
	value    string // the value we expect to be returned by the selector
}
