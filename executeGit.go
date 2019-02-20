package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"log"
	"os"
)

func main() {
	var firstArgs string
	var path string
	var option string

	if len(os.Args) == 1 {
	  log.Fatal("No arguments given, Please specify arguments.")
	  return
	}
	if firstArgs = os.Args[1]; firstArgs == "" {
		log.Fatal("No option given, Please specify option.")
		return
	}
	if firstArgs == "--help" {
		fmt.Printf("-s\tgit status\n-P\tgit pull\n-c\tgit-clean\n-b\tgit branch\n")
		return
	}
	if path = os.Args[2]; path == "" {
		log.Fatal("No path given, Please specify path.")
		return
	}
	if firstArgs == "-s" {
		option = "git status"
	} else if firstArgs == "-P" {
		option = "git pull"
	} else if firstArgs == "-c" {
		option = "git remote prune origin && git branch -vv | grep \"origin/.*: gone]\" | awk \"{print }\" | xargs git branch -D 2>/dev/null"
	} else if firstArgs == "-b" {
		option = "git branch"
	} else {
		log.Fatal("Wrong option given. Please call --help")
		return
	}

	ret := executeGit(path, option)
	if ret == false {
		fmt.Printf("Error with the script\n")
	}
}

func executeGit(path string, option string) (bool) {
	green := "\x1b[32;1m"
	black := "\x1b[0m"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	sep := "--------------------------------------\n--------------------------------------"
	for _, file := range files {
		name := file.Name()
		if file.IsDir() {
			var filePath string   = filepath.Join(path, name)
			var filePathSh string = path+"/"+name

			if isGit(filePath) {
				var cmd string = "cd "+filePathSh+" && "+option
				out, err := exec.Command("bash","-c", cmd).Output()
				if err != nil {
					fmt.Printf("%s\n", err)
				} else {
					fmt.Printf("%s\n%s%s%s\n%s\n%s\n", sep, green, name, black, sep, out)
				}
				} else {
					executeGit(filePathSh, option)
			}
		}
	}
	return true
}

func isGit(path string) (bool) {
	Files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, File := range Files {
		if File.IsDir() && File.Name() == ".git" {
			return true
		}
	}
	return false
}

func checkDir(path string) ([]string) {
	Files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	var listDir []string
	for _, File := range Files {
		if File.IsDir() {
			listDir = append(listDir, File.Name())
		}
	}
	return listDir
}