package repository

import (
	"log"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

//type config struct {
//	userToken string
//	owner     string
//	repo      string
//}

type client struct {
	config *ServiceConfig
	github *github.Client
}

func New(cfg *ServiceConfig) *client {
	c := &client{
		config: cfg,
	}
	c.authenticateWithToken()
	return c
}

func (c *client) Connect() error {
	_, _, err := c.github.Zen()
	return err
}

func (g *client) authenticateWithToken() *client {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: g.config.Token,
		},
	)

	client := oauth2.NewClient(oauth2.NoContext, src)
	g.github = github.NewClient(client)

	return g
}

func (g *client) GetPullRequest(id PullRequestID) *PullRequest {

	res := g.fetchPullRequest(id)

	return &PullRequest{
		Id:          PullRequestID(*res.Number),
		User:        res.User.Login,
		Title:       res.Title,
		Description: res.Body,
		Url:         res.URL,
	}
}

func (g *client) GetPullRequests(ids []PullRequestID) []*PullRequest {

	prs := []*PullRequest{}
	for _, id := range ids {
		res := g.fetchPullRequest(id)
		pr := &PullRequest{
			Id:          PullRequestID(*res.Number),
			User:        res.User.Login,
			Title:       res.Title,
			Description: res.Body,
			Url:         res.URL,
		}

		prs = append(prs, pr)
	}

	return prs
}

func (g *client) fetchPullRequest(i PullRequestID) *github.PullRequest {

	//id, err := strconv.Atoi(i)
	//if err != nil {
	//	log.Fatalf("Something went wrong converting to string: %+v", err)
	//}

	pr, _, err := g.github.PullRequests.Get(g.config.Owner, g.config.Repo, id)
	if err != nil {
		log.Fatalf("Something went wrong talking to GH: %+v", err)
	}

	return pr
}
