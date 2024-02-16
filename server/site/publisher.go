package site

import (
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/google/uuid"
)

func cloneGitHistory(dname string, repository string, accessToken string) (*git.Repository, error) {
	return git.PlainClone(dname, true, &git.CloneOptions{
		Auth: &http.TokenAuth{
			Token: accessToken,
		},
	})
}

func removeTempFiles(list []string) {
	for _, file := range list {
		os.Remove(file)
	}
}

func writeFilesToDir(dname string, page *PageOutput) ([]string, error) {
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

			result, err := writeFilesToDir(subPageDname, subPage)
			if err != nil {
				return nil, err
			}

			fileList = append(fileList, result...)
			// map additional files
		}
	}

	return fileList, nil
}

func publishSite(siteID string, site *Site, output *SiteOutput) error {
	dname, err := os.MkdirTemp("", siteID)
	fileList := make([]string, 0)
	if err != nil {
		return err
	}

	defer os.Remove(dname)
	repo, err := cloneGitHistory(dname, site.Repository, site.Credentials.AccessToken)
	if err != nil {
		return err
	}

	for _, page := range output.Pages {
		l, err := writeFilesToDir(dname, page)
		if err != nil {
			return err
		}

		fileList = append(fileList, l...)
	}

	defer removeTempFiles(fileList)

	repoErr := repo.CreateBranch(&config.Branch{
		Name: uuid.NewString(),
	})
	if repoErr != nil {
		return repoErr
	}

	// Add to created branch
	// Push changes
	// Merge
	// Remove branch
	// Commit merge

	return nil
}

