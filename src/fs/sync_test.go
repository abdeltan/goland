package fs

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"
)

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Dir(path.Dir(b) + "/../../")
}

func TestCopyFile(t *testing.T) {
	cwd := RootDir()

	if err := os.RemoveAll(cwd + "/test/targets"); err != nil {
		t.Error("Impossible de supprimer le dossier de test", err)
	}

	if err := os.MkdirAll(cwd+"/test/targets", PERMS_DIRECTORY); err != nil {
		t.Error("Impossible de creer le dossier de test", err)
	}

	if err := file(cwd+"/test/sources/a", cwd+"/test/targets/a"); err != nil {
		t.Error("Impossible de copier le fichier", err)
	}
}

func TestCopyFolder(t *testing.T) {
	cwd := RootDir()

	if err := os.RemoveAll(cwd + "/test/targets"); err != nil {
		t.Error("Impossible de supprimer le dossier de test", err)
	}

	if err := os.MkdirAll(cwd+"/test/targets", PERMS_DIRECTORY); err != nil {
		t.Error("Impossible de creer le dossier de test", err)
	}

	if err := folder(cwd+"/test/sources", cwd+"/test/targets"); err != nil {
		t.Error("Impossible de copier le dossier", err)
	}
}
