package competition

type CompetitionClosed struct {
	CompetitionID string
	Winners       []string // sorted by points
}
