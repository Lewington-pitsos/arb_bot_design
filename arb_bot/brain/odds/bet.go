package odds

type bet struct {
	event    string
	outcome  string
	provider string
	oddUrl   string // the fully qualified url where the bet odd is sitting (we will assume that it is possible to conststantly convert this to a bet-making url if everything else in bet is defined properly)
	amount   int
	person   person // the details of the user who will be making this bet
}
