package brain

type snapshot struct {
	state map[string]map[string]map[string]odd
	// event => outcome => oddProvider => odd
	// liverpool vs arsenal 02/03/2019 => win => ladbrokes => 3.24
}

// fileReport adds the report data (which should always be an odd) to the current snapshot state if the report's name is compatible with the snapshot.
func (s *snapshot) fileReport(report) {

}
