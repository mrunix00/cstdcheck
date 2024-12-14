package main

import (
	"flag"
	"os"
	"path/filepath"
	"runtime"
)

func getCCPath() string {
	var showVersion bool
	var showHelp bool

	flag.BoolVar(&showVersion, "v", false, "Show version")
	flag.BoolVar(&showHelp, "h", false, "Show help")

	flag.Parse()
	if showVersion {
		println("Cstdcheck v0.0.0")
		os.Exit(0)
	}

	if showHelp {
		println("Usage: cstdcheck [options] [files]")
		println("Options:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	args := flag.Args()

	if len(args) != 1 {
		println("Usage: cstdcheck [options] /path/to/cc")
		println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	return args[0]
}

func getCachePath() (error, string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err, ""
	}

	switch runtime.GOOS {
	case "windows":
		return nil, filepath.Join(homeDir, "AppData", "Local", "Cache", "cstdcheck")
	case "darwin":
		return nil, filepath.Join(homeDir, "Library", "Caches", "cstdcheck")
	default:
		return nil, filepath.Join(homeDir, ".cache", "cstdcheck")
	}
}

func isExecutable(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	} else if runtime.GOOS == "windows" {
		return filepath.Ext(path) == ".exe"
	}
	return fileInfo.Mode()&0111 != 0
}

func findCFiles(path string) []string {
	var files []string
	entries, err := os.ReadDir(path)
	if err != nil {
		return files
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		ext := filepath.Ext(name)
		if ext != ".c" && ext != ".h" && ext != ".x" {
			continue
		}
		files = append(files, filepath.Join(path, name))
	}
	return files
}
