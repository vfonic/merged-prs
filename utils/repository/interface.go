package repository

type Service interface {
	Connect() error
	GetPullRequest(PullRequestID) *PullRequest
	GetPullRequests([]PullRequestID) []*PullRequest
}

type ServiceConfig struct {
	Owner string
	Repo  string
	Token string
}

type PullRequestID int

type PullRequest struct {
	Id          *int
	User        *string
	Title       *string
	Description *string
	Url         *string
}
