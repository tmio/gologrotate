package main

import (
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jasonlvhit/gocron"
)

var format = flag.String("format", "2006-02-01", "Time format")

func findFiles(searchDir string, suffix string) []string {
	fileList := []string{}
	filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if strings.HasSuffix(path, suffix) {
			fileList = append(fileList, path)
		}
		return nil
	})

	return fileList
}

func removeUnderlyingFile(name string, file *os.File) {
	fi, err := file.Stat()
	if err != nil {
		fmt.Fprintln(os.Stderr, "couldn't stat", name, err)
		return
	}
	nfi, err := os.Stat(name)
	if err != nil {
		fmt.Fprintln(os.Stderr, "couldn't delete", name, err)
		return
	}
	if os.SameFile(fi, nfi) {
		os.Remove(name)
	}
}

func copyTruncate(nameIn, nameOut string) (err error, needPanic bool) {
	// copy with gzip and truncate a file.
	// filename may be relative or absolute.
	// this will cd to the directory with the file.
	out, err := os.Create(nameOut)
	if err != nil {
		return err, false
	}
	in, err := os.OpenFile(nameIn, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err, false
	}
	defer out.Close()
	defer in.Close()
	defer func() {
		// if the function fails, let's delete the file we were creating to save space.  OKAY!
		if r := recover(); r != nil {
			removeUnderlyingFile(nameOut, out)
		}
	}()

	gzout := gzip.NewWriter(out)
	defer gzout.Close()

	// Don't truncate until we've got 10 empty reads in a row
	for i := 0; i < 10; i++ {
		written, err := io.Copy(gzout, in)
		if written > 0 {
			i = 0
			fmt.Print(".")
		}
		if err != nil {
			return err, true
		}
	}
	err = in.Truncate(0)
	if err != nil {
		return fmt.Errorf("Couldn't truncate file: %v", err), true
	}
	in.Sync()
	fmt.Println(fmt.Sprintf("Done - %s", nameIn))
	return nil, false
}

func fileExists(name string) bool {
	_, err := os.Stat(name)
	return (err == nil)
}

func run(searchDir string) {
	fileList := findFiles(searchDir, ".log")
	now := time.Now().Format(*format)
	for _, file := range fileList {
		outname := fmt.Sprintf("%s.%s.gz", file, now)
		n := 0
		for fileExists(outname) {
			n += 1
			outname = fmt.Sprintf("%s.%s.%d.gz", file, now, n)
		}
		err, needPanic := copyTruncate(file, outname)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error copyTruncate from %s to %s: %v", file, outname, err)
			if needPanic {
				panic(err)
			}
		}
	}
}

func main() {
	flag.Parse()
	for _, arg := range flag.Args() {
		gocron.Every(1).Day().At("01:00").Do(run, arg)
	}
	gocron.Start()
}
