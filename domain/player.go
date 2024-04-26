package domain

type Player struct {
	Person

	Teams       []Team
	CurrentTeam *Team
}

type Coach struct {
	Person

	Teams       []Team
	CurrentTeam *Team
}

type Staff struct {
	Person

	Teams       []Team
	CurrentTeam *Team
}
