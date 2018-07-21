package data

type stepInterface interface {
	Url() string
	GenerateUrl(interface{})
}
