package utils

import (
	"os"

	"golang.org/x/sys/unix"
)

type FileProcessor struct {
	Files []string
}

func (p *FileProcessor) FileExists(path string) bool {
	_, err := os.Stat(path)
	os.IsNotExist(err)

	return !os.IsNotExist(err)
}

func (p *FileProcessor) FileIsWritable(path string) bool {
	return unix.Access(path, unix.W_OK) == nil
}

func (p *FileProcessor) Register(path string) *FileProcessor {
	if !fileProcessor.FileIsWritable(path) {
		LoggerService.Error(path + " is unwritable!")
		return p
	}

	p.Files = append(p.Files, path)

	return p
}

func (p *FileProcessor) Proceed(cb func(path string)) {
	if len(p.Files) == 0 {
		LoggerService.Info("No log file were found. Exiting")
		os.Exit(0)
	}

	for _, file := range p.Files {
		cb(file)
	}
}
