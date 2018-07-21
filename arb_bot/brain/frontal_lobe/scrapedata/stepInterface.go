package scrapedata

type stepInterface interface {
	Url() string
	GenerateUrl(interface{})
}
