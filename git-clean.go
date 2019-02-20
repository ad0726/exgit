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
	 var root string

	if len(os.Args) == 1 {
	  log.Fatal("No path given, Please specify path.")
	  return
	 }
	 if root = os.Args[1]; root == "" {
	  log.Fatal("No path given, Please specify path.")
	  return
	 }
    files, err := ioutil.ReadDir(root)
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
		if file.IsDir() {
			var filePath string = filepath.Join(root, file.Name())

			if checkGit(filePath) {
				var filePathSh string = root+"/"+file.Name()
				var cmd string = "cd "+filePathSh+" && git remote prune origin && git branch -vv | grep \"origin/.*: gone]\" | awk \"{print }\" | xargs git branch -D 2>/dev/null"
				// fmt.Printf("%s\n", cmd)
					out, err := exec.Command("bash","-c", cmd).Output()
					if err != nil {
						fmt.Printf("%s\n", err)
					} else {
							fmt.Printf("%s\n", out)
					}
			} else {
				// todo
				// listDir = checkDir(filePath)
			}


		}
	}
}

func checkGit(path string) (bool) {
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