package publisher

import (
	"os"

	"app/db"
	git_internal "app/git"
	"app/site"

	"github.com/go-git/go-git/v5"
)

type Publisher interface {
	PublishSite(siteID string, site *db.Site, output *site.SiteOutput) error
}

type PublisherService struct {
	git git_internal.GitWrapper
}

func commitGitFiles() {
}

func removeTempFiles(list []string) {
	for _, file := range list {
		os.Remove(file)
	}
}

func (s *PublisherService) writeFilesToDir(dname string, page *site.PageOutput) ([]string, error) {
	fileList := make([]string, 0)
	f, err := os.CreateTemp(dname, page.Slug+".html")
	if err != nil {
		return nil, err
	}

	fileList = append(fileList, f.Name())

	if page.SubPages != nil && len(page.SubPages) > 0 {
		for _, subPage := range page.SubPages {
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

func (s *PublisherService) PublishSite(siteID string, site *site.Site, output *site.SiteOutput) error {
	dname, err := os.MkdirTemp("", siteID)
	fileList := make([]string, 0)
	if err != nil {
		return err
	}

	defer os.Remove(dname)
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

	defer removeTempFiles(fileList)

	return nil
}

func New(gitInternal git_internal.GitWrapper) *PublisherService {
	return &PublisherService{
		git: gitInternal,
	}
}
