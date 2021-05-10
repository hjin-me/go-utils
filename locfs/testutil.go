package locfs

import (
	"embed"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func FSPathMust(efs embed.FS, p string) string {
	tmpDir, err := ioutil.TempDir("", "fs_path_")
	if err != nil {
		panic(err)
	}
	s, err := fsPathDir(efs, p, tmpDir)
	if err != nil {
		panic(err)
	}
	return s
}

func FSPath(efs embed.FS, p string) (string, error) {
	tmpDir, err := ioutil.TempDir("", "fs_path_")
	if err != nil {
		return "", err
	}
	return fsPathDir(efs, p, tmpDir)
}

func fsPathDir(efs embed.FS, p string, baseDir string) (string, error) {
	//log.Printf("fsPathDir: %s, %s\r\n", p, baseDir)
	f, err := efs.Open(p)
	if err != nil {
		return "", err
	}
	s, err := f.Stat()
	if err != nil {
		return "", err
	}
	if s.IsDir() {
		err := os.Mkdir(filepath.Join(baseDir, s.Name()), os.ModePerm)
		if err != nil {
			return "", err
		}
		edir, err := efs.ReadDir(p)
		if err != nil {
			return "", err
		}
		for _, entry := range edir {
			_, err := fsPathDir(efs, filepath.Join(p, entry.Name()), filepath.Join(baseDir, s.Name()))
			if err != nil {
				return "", err
			}
		}
		return filepath.Join(baseDir, s.Name()), nil
	}
	tf, err := os.Create(filepath.Join(baseDir, s.Name()))
	if err != nil {
		return "", err
	}

	_, err = io.Copy(tf, f)
	if err != nil {
		return "", err
	}
	return tf.Name(), nil
}
