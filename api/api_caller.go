package api

// Interface for Sakamichi API call.
type ApiCaller interface {
	ListGroups() ([]Group, error)
	ListMembers(gn string) ([]Member, error)
}
