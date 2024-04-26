package domain

type Team struct {
	Name    string
	TLA     string //alias
	Address string
	Venue   string
	WebSite string

	Coach *Coach
	Squad []Player
	Staff []Staff

	Competitions []Competition
}
