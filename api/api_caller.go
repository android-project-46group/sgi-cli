package api

// Sakamichi API caller
type ApiCaller interface {
	ListGroups() ([]Group, error)
	ListMembers(gn string) ([]Member, error)
}
