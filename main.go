package main

import (
	"flag"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	folderFlag = flag.String("folder", "", "The folder to delete")
	simulate = flag.Bool("simulate", false, "IF the delete should only be simulated")
)

func main() {
	flag.Parse()

	folder := *folderFlag

	if folder == "" {
		log.Fatalln("Folder not specified")
	}

	foldersToRemove := make([]string, 0)

	err := filepath.WalkDir(".", func(p string, d fs.DirEntry, _ error) error {
		if d.IsDir() && strings.EqualFold(d.Name(), folder) {
			foldersToRemove = append(foldersToRemove, p)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Failed to walk directory: %v", err)
	}

	for _, f := range foldersToRemove {
		log.Printf("Deleting '%s'", f)
		if !*simulate {
			err = os.RemoveAll(f)
			if err != nil {
				log.Printf("Failed to delete '%s': %v", f, err)
			}
		}
	}
}
