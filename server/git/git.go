package git_internal

import "github.com/go-git/go-git/v5"

type GitWrapper interface {
	CloneHistory(dname string, options *git.CloneOptions)
}

type GitHandler struct{}

func (g *GitHandler) CloneHistory(dname string, options *git.CloneOptions) {
	git.PlainClone(dname, true, options)
}

func New() *GitHandler {
	return &GitHandler{}
}
