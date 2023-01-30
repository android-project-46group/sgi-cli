package api

// Sakamichi API caller
type ApiCaller interface {
	ListMembers(gn string) ([]Member, error)
}
