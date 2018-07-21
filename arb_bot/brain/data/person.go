package data

type Person struct {
	winnings int               // total tally of money won or lost by that person
	details  map[string]string // the details of the given person for logging in
}
