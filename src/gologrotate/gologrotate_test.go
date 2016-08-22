package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestFindLogFilesRecursively(t *testing.T) {
	t.Log("Testing finding log files recursively")
	fileList := findFiles("tests/find_all", ".log")
	root_log := false
	sub_log := false
	for _, file := range fileList {
		if file == "tests/find_all/root.log" {
			root_log = true
		} else if file == "tests/find_all/not_picking.loga" || file == "tests/find_all/notpicking.txt" {
			t.Error("Picked a file it shouldn't pick")
		} else if file == "tests/find_all/log/sub_log/sub.log" {
			sub_log = true
		}
	}
	if !root_log {
		t.Error("Root log not picked")
	}
	if !sub_log {
		t.Error("Sub log not picked")
	}
}

func ReadFile(path string) string {
	bytes, _ := ioutil.ReadFile(path)
	return string(bytes)
}

func TestRun(t *testing.T) {
	t.Log("Testing running the process")

	root_log_name := "tests/find_all/root.log"

	format := "2016-02-16"

	fi, err := os.OpenFile(root_log_name, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	fi.WriteString("FOO\n")
	fi.WriteString("BAR\n")

	t.Log(fmt.Sprintf("File contents: %q", ReadFile(root_log_name)))

	run("tests/find_all", format)

	fi.WriteString("NEW LINE\n")

	defer fi.Close()

	data := ReadFile(root_log_name)
	t.Log(fmt.Sprintf("File contents: %q", data))
	if data != "NEW LINE\n" {
		t.Error("File was not truncated correctly")
	}

	os.Remove(root_log_name)
	os.OpenFile(root_log_name, os.O_RDONLY|os.O_CREATE, 0644)

	allGz := findFiles("tests/find_all", ".gz")
	now := time.Now().Format(format)
	datesuffix := fmt.Sprintf("%s", now)
	root_log := false
	sub_log := false
	for _, file := range allGz {
		if file == "tests/find_all/root.log."+datesuffix+".gz" {
			root_log = true
		} else if file == "tests/find_all/not_picking.loga."+datesuffix+".gz" || file == "tests/find_all/notpicking.txt."+datesuffix+".gz" {
			t.Error("Picked a file it shouldn't pick")
		} else if file == "tests/find_all/log/sub_log/sub.log."+datesuffix+".gz" {
			sub_log = true
		}
		os.Remove(file)
	}
	if !root_log {
		t.Error("Root log not picked")
	}
	if !sub_log {
		t.Error("Sub log not picked")
	}
}
