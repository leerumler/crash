package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Reflection do
type Reflection struct {
	AbsPath, RelPath string
	Mode             os.FileMode
	Dir              bool
	crawled          bool
	// Content []byte
}

func crawl(ref *Reflection) []Reflection {

	ref.crawled = true

	// Get a list of all the files in the directory.
	files, err := ioutil.ReadDir(ref.AbsPath)
	if err != nil {
		return nil
	}

	//
	var reflections []Reflection
	for _, file := range files {

		var reflection Reflection
		reflection.AbsPath = ref.AbsPath + "/" + file.Name()
		reflection.Mode = file.Mode()
		reflection.Dir = file.IsDir()

		if !reflection.Dir {
			reflection.crawled = true
		}

		reflections = append(reflections, reflection)

	}

	return reflections

}

func deepCrawl(dir string) []Reflection {

	var rootRef Reflection
	rootRef.AbsPath = dir
	if file, err := os.Stat(dir); err != nil {
		rootRef.Mode = file.Mode()
		rootRef.Dir = file.IsDir()
	}

	var refs []Reflection
	refs = append(refs, rootRef)
	newRefs := crawl(&rootRef)
	refs = append(refs, newRefs...)

	for again := true; again; {
		for _, ref := range newRefs {
			if !ref.crawled {
				newRefs = crawl(&ref)
				refs = append(refs, newRefs...)
				again = true
			} else {
				again = false
			}
		}
	}

	return refs
}

const crashPath = "/home/tbg/crash"

func main() {

	refs := deepCrawl(crashPath)
	for _, ref := range refs {
		if ref.Dir {
			fmt.Println("Directory:", ref.AbsPath)
		} else {
			fmt.Println("File Path:", ref.AbsPath)
		}

		// fmt.Println()
	}
}
