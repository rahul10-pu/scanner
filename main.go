package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
)

func main() {

	// now := time.Now() //doHashWalk("/Users/rahulkumar/CRUD-with-GRPC-Golang")
	// id := uuid.New()
	// fmt.Println(id.String())
	name := "https://github.com/go-git/go-git"
	re := strings.Split(name, "/")
	log.Println(re[4])
	a, err := git.PlainClone("hehe", false, &git.CloneOptions{
		URL:      "https://github.com/go-git/go-git",
		Progress: os.Stdout,
	})
	if err != nil {
		log.Panic(err)
	}
	log.Print(a)

}
func doHashWalk(dirPath string) error {
	fullPath, err := filepath.Abs(dirPath)
	log.Println(fullPath)
	if err != nil {
		return err
	}
	return filepath.Walk(fullPath, walkFn)

}
func hashFile(root string, path string, fi os.FileInfo, err error) error {
	if fi.IsDir() {
		return nil
	}
	log.Println("hashfile")
	log.Println(root, path)
	rel, err := filepath.Rel(root, path)

	if err != nil {
		return err
	}

	log.Println("hash rel:", rel, "abs:", path)

	return nil
}
func walkFn(path string, fi os.FileInfo, err error) (e error) {
	if !fi.IsDir() {
		bytes, err := ioutil.ReadFile(path) // path is the path to the file.
		if err != nil {
			fmt.Println("Fail")
		}
		log.Println(string(bytes))
	}
	return nil
}
