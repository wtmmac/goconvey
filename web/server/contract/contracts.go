package contract

import (
	"net/http"
	"path/filepath"
)

type (
	Server interface {
		ReceiveUpdate(*CompleteOutput)
		Watch(writer http.ResponseWriter, request *http.Request)
		Ignore(writer http.ResponseWriter, request *http.Request)
		Reinstate(writer http.ResponseWriter, request *http.Request)
		Status(writer http.ResponseWriter, request *http.Request)
		Results(writer http.ResponseWriter, request *http.Request)
		Execute(writer http.ResponseWriter, request *http.Request)
	}

	Executor interface {
		ExecuteTests([]*Package) *CompleteOutput
		Status() string
	}

	Scanner interface {
		Scan() (changed bool)
	}

	Watcher interface {
		Root() string
		Adjust(root string) error

		Deletion(folder string)
		Creation(folder string)

		Ignore(folder string)
		Reinstate(folder string)

		WatchedFolders() []*Package
		IsWatched(folder string) bool
		IsIgnored(folder string) bool
	}

	FileSystem interface {
		Walk(root string, step filepath.WalkFunc)
		Exists(directory string) bool
	}

	Shell interface {
		Execute(name string, args ...string) (output string, err error)
		Getenv(key string) string
		Setenv(key, value string) error
	}
)
