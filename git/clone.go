package git

import (
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/src-d/go-billy.v4/memfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

// Clone git clone
func Clone() {

	url := "https://github.com/gogs/docs-api.git"
	directory := "/Users/chenfeng/work/opensource/go-git-demo2"

	// Clone the given repository to the given directory
	//Info("git clone %s %s --recursive", url, directory)

	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	checkIfError(err)

	// ... retrieving the branch being pointed by HEAD
	ref, err := r.Head()
	checkIfError(err)
	// ... retrieving the commit object
	commit, err := r.CommitObject(ref.Hash())
	checkIfError(err)

	fmt.Println(commit)
}

// CloneInMem git clone in mmemory
func CloneInMem() {
	// Filesystem abstraction based on memory
	fs := memfs.New()
	// Git objects storer based on memory
	storer := memory.NewStorage()

	// Clones the repository into the worktree (fs) and storer all the .git
	// content into the storer
	r, err := git.Clone(storer, fs, &git.CloneOptions{
		URL: "https://github.com/git-fixtures/basic.git",
	})
	if err != nil {
		log.Fatal(err)
	}

	head,err := r.Head()

	commit,err := r.CommitObject(head.Hash())

	fmt.Println(commit.Message,commit.Hash,commit.Author.Email)

	// Prints the content of the CHANGELOG file from the cloned repository
	changelog, err := fs.Open("CHANGELOG")
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, changelog)
}

func checkIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}
