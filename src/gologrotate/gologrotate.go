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
			return err, false
		}
	}
	err = in.Truncate(0)
	if err != nil {
		return fmt.Errorf("Couldn't truncate file: %v", err), false
	}
	in.Sync()
	fmt.Println(fmt.Sprintf("Done - %s", nameIn))
	return nil, false
}

func fileExists(name string) bool {
	_, err := os.Stat(name)
	return (err == nil)
}

func run(searchDir string, format string) {
	fmt.Println(fmt.Sprintf("Starting execution on %s", searchDir))
	fileList := findFiles(searchDir, ".log")
	now := time.Now().Format(format)
	for _, file := range fileList {
		outname := fmt.Sprintf("%s.%s.gz", file, now)
		n := 0
		for fileExists(outname) {
			n += 1
			outname = fmt.Sprintf("%s.%s.%d.gz", file, now, n)
		}
		fmt.Println(fmt.Sprintf("Truncating %s into %s", file, outname))
		err, needPanic := copyTruncate(file, outname)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error copyTruncate from %s to %s: %v", file, outname, err)
			if needPanic {
				panic(err)
			}
		}
	}
	fmt.Println(fmt.Sprintf("Done with execution on %s", searchDir))
}

func main() {
	fmt.Println("Starting gologrotate")
	now := flag.Bool("now", false, "Run now")
	time := flag.String("time", "23:55", "Local time at which the cron job runs")
	format := flag.String("format", "2006-02-16", "Time format")
	flag.Parse()
	if *now {
		fmt.Println("Running a one-time execution of gologrotate")
		for _, arg := range flag.Args() {
			fmt.Println(fmt.Sprintf("Running on %s", arg))
			run(arg, *format)
		}
	} else {
		fmt.Println("Running a cron job of gologrotate")
		for _, arg := range flag.Args() {
			fmt.Println(fmt.Sprintf("Adding %s to watchlist", arg))
			gocron.Every(1).Day().At(*time).Do(run, arg, *format)
		}
		_, time := gocron.NextRun()
		fmt.Println(fmt.Sprintf("Cron job will run next at %s", time))
		<-gocron.Start()
	}
}
