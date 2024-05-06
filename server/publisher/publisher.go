package publisher

import (
	"os"

	git_internal "github.com/DogAndHerDude/web-builder/git"
	"github.com/DogAndHerDude/web-builder/internal/app/db"
	"github.com/DogAndHerDude/web-builder/internal/pkg/builder"

	"github.com/go-git/go-git/v5"
)

type Publisher interface {
	PublishSite(siteID string, site *db.Site, output *builder.BuildResult) error
}

type PublisherService struct {
	git git_internal.GitWrapper
}

func commitGitFiles() {
}

func (s *PublisherService) writeFilesToDir(dname string, page *builder.PageBuildResult) ([]string, error) {
	fileList := make([]string, 0)
	f, err := os.CreateTemp(dname, page.Slug+".html")
	if err != nil {
		return nil, err
	}

	fileList = append(fileList, f.Name())

	if page.Pages != nil && len(page.Pages) > 0 {
		for _, subPage := range page.Pages {
			subPageDname, err := os.MkdirTemp(dname, page.Slug)
			if err != nil {
				return nil, err
			}

			defer os.Remove(subPageDname)

			result, err := s.writeFilesToDir(subPageDname, subPage)
			if err != nil {
				return nil, err
			}

			fileList = append(fileList, result...)
			// map additional files
		}
	}

	return fileList, nil
}

func (s *PublisherService) PublishSite(siteID string, site *db.Site, output *builder.BuildResult) error {
	dname, err := os.MkdirTemp("", siteID)
	fileList := make([]string, 0)
	if err != nil {
		return err
	}

	defer os.RemoveAll(dname)
	s.git.CloneHistory(dname, &git.CloneOptions{
		URL: site.Repository,
	})

	for _, page := range output.Pages {
		l, err := s.writeFilesToDir(dname, page)
		if err != nil {
			return err
		}

		fileList = append(fileList, l...)
	}

	return nil
}

func New(gitInternal git_internal.GitWrapper) *PublisherService {
	return &PublisherService{
		git: gitInternal,
	}
}
