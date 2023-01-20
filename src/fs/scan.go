package fs

import (
	"os"

	"github.com/schollz/progressbar/v3"
)

var PROGRESS_BAR *progressbar.ProgressBar

func Scan(origin string, target string) ([]string, []string, error) {
	files := make([]string, 0)
	folders := make([]string, 0)
	entries, err := os.ReadDir(origin)
	if err != nil {
		return folders, files, err
	}

	folders = append(folders, origin+":"+target)

	if len(entries) == 0 {
		return folders, files, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			fo, fi, _ := Scan(origin+"/"+entry.Name(), target+"/"+entry.Name())
			folders = append(folders, fo...)
			files = append(files, fi...)
		} else {
			files = append(files, origin+"/"+entry.Name()+":"+target+"/"+entry.Name())
		}
	}

	return folders, files, nil
}
