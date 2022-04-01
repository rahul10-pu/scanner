package utility

import (
	"os"

	"github.com/go-git/go-git/v5"
)

func Clone(repoName, repoURL string) (err error) {
	//username := strings.Split(repoURL, "/")[4]
	//repoName=username+"-"+repoName
	_, err = git.PlainClone(repoName, false, &git.CloneOptions{
		URL:      repoURL,
		Progress: os.Stdout,
	})
	return
}
func RepoExist(repoName string) (err error) {
	return
}
