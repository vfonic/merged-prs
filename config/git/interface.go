package git

type Client interface {
	Available() bool
	IsRepo(string) bool
	IsRef(string) bool
	GetMergeCommits(string, string) []*MergeCommit
}

type MergeCommit struct {
	Ref     string
	Message string
}
