package data

type assignment struct {
	name     string // name of the field whose data we've been assigned to get
	selector string // an xpath selector
}

// generateReport returns a report form the selector and passed in xml
func (a *assignment) GenerateReport( /*Xpath object*/ ) {

}
