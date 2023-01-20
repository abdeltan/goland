package fs

import (
	"fmt"
	"os"
)

const PERMS_DIRECTORY = 0755

func Copy(origin string, target string) error {
	fi, err := os.Stat(origin)
	if os.IsNotExist(err) {
		return err
	}

	if !fi.IsDir() {
		if err := os.RemoveAll(target); err != nil {
			return err
		}

		return file(origin, target)
	}

	if fi, _ := os.Stat(target); fi != nil {
		if err := os.RemoveAll(target); err != nil {
			return err
		}
	}

	if err := os.MkdirAll(target, PERMS_DIRECTORY); err != nil {
		return err
	}

	return folder(origin, target)
}

func file(origin string, target string) error {
	fi, err := os.Open(origin)
	if err != nil {
		return err
	}
	defer fi.Close()

	fo, err := os.Create(target)
	if err != nil {
		return err
	}
	defer fo.Close()

	buffer := make([]byte, 1024)
	for {
		nb, err := fi.Read(buffer)
		if nb == 0 {
			break
		} else if err != nil {
			return err
		}

		if _, err := fo.Write(buffer[:nb]); err != nil {
			return err
		}
	}

	if DEBUG {
		fmt.Println("COPY:", origin, "->", target, "FILE")
	}

	if PROGRESS {
		PROGRESS_BAR.Add(1)
	}

	return nil
}

func folder(origin string, target string) error {
	entries, err := os.ReadDir(origin)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			err := CopyFolder(origin+"/"+entry.Name(), target+"/"+entry.Name())
			if err != nil {
				return err
			}

			folder(origin+"/"+entry.Name(), target+"/"+entry.Name())
		} else {
			err := CopyFile(origin+"/"+entry.Name(), target+"/"+entry.Name())
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func CopyFolder(origin string, target string) error {
	if DEBUG {
		fmt.Println("COPY FOLDER:", origin, "->", target)
	}

	if PROGRESS {
		PROGRESS_BAR.Add(1)
	}

	return os.Mkdir(target, PERMS_DIRECTORY)
}

func CopyFile(origin string, target string) error {
	if DEBUG {
		fmt.Println("COPY FILE:", origin, "->", target)
	}

	if PROGRESS {
		PROGRESS_BAR.Add(1)
	}

	return file(origin, target)
}
