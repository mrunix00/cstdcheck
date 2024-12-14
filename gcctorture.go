package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

type Results struct {
	all    int
	passed int
	failed int
}

func FetchGccTorture() error {
	const gccTortureURL = "svn://gcc.gnu.org/svn/gcc/trunk/gcc/testsuite/gcc.c-torture/execute"
	err, cachePath := getCachePath()
	if err != nil {
		return err
	}

	pathDownload := filepath.Join(cachePath, "gcc-torture")
	cmd := exec.Command("svn", "co", gccTortureURL, pathDownload)
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func TestGccTorture(cc string) (error, Results) {
	err, cachePath := getCachePath()
	if err != nil {
		return err, Results{}
	}
	pathToFiles := filepath.Join(cachePath, "gcc-torture")
	files := findCFiles(pathToFiles)

	all, passed, failed := len(files), 0, 0
	for _, file := range files {
		cmd := exec.Command(cc, file, "-o", "/dev/null")
		err := cmd.Run()
		if err != nil {
			failed++
		} else {
			passed++
		}
	}

	return nil, Results{all, passed, failed}
}

func TestIeeeTorture(cc string) (error, Results) {
	err, cachePath := getCachePath()
	if err != nil {
		return err, Results{}
	}
	pathToFiles := filepath.Join(cachePath, "gcc-torture", "ieee")
	files := findCFiles(pathToFiles)

	all, passed, failed := len(files), 0, 0
	for _, file := range files {
		cmd := exec.Command(cc, file, "-o", "/dev/null")
		err := cmd.Run()
		if err != nil {
			failed++
		} else {
			passed++
		}
	}

	return nil, Results{all, passed, failed}
}
