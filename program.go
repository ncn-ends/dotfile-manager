package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println(color("yellow", "Starting..."))

	homeDir := findHomeDir()
	fmt.Printf(
		color("purple", "Searching from home directory for csgo cfg directory: \n\t%s\n"),
		color("cyan", homeDir))

	cfgDir := findCfgDir(homeDir)
	fmt.Printf(
		color("purple", "Found output directory: \n\t%s\n"),
		color("cyan", cfgDir))

	fmt.Println(color("purple", "Looking for cfg files in local cfg directory..."))
	cfgFiles := findCfgContents()

	fmt.Printf(color("purple", "Copying over to output directory:\n"))
	for _, path := range cfgFiles {
		fileName := strings.TrimLeft(path, "cfg/")
		fmt.Printf(
			color("cyan", "\t%s\n"),
			fileName)
		moveFile(path, cfgDir+"/"+fileName)
	}

	fmt.Println(color("green", "Complete!"))
}

func color(colorString string, message string) string {
	resetColor := "\033[0m"

	matchedColor := matchColor(colorString)
	return matchedColor + message + resetColor
}

func colorPrint(colorString string, message string) {
	resetColor := "\033[0m"

	matchedColor := matchColor(colorString)
	fmt.Printf(matchedColor + message + resetColor + "\n")
}

func matchColor(colorString string) string {
	if colorString == "green" {
		return "\033[32m"
	} else if colorString == "yellow" {
		return "\033[33m"
	} else if colorString == "purple" {
		return "\033[35m"
	} else if colorString == "cyan" {
		return "\033[36m"
	}
	log.Fatal("Color passed as argument does not match an accepted color.")
	return ""
}

func findHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return homeDir
}

func moveFile(pathToIn string, pathToOut string) {
	input, err := ioutil.ReadFile(pathToIn)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(pathToOut, input, 0644)
	if err != nil {
		fmt.Println("Error creating", pathToOut)
		fmt.Println(err)
		return
	}
}

func findCfgDir(homeDir string) string {

	COMMON_CFG_PATH := "Steam/steamapps/common/Counter-Strike Global Offensive/csgo/cfg"
	// COMMON_CFG_PATH := "out/testcfg"
	var files []string

	err := filepath.Walk(homeDir, func(path string, info os.FileInfo, err error) error {
		if err != nil && strings.Contains(path, "permission denied") {
			fmt.Println(err)
			return nil
		}

		if strings.HasSuffix(path, COMMON_CFG_PATH) && info.IsDir() {
			files = append(files, path)

		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return files[0]
}

func findCfgContents() []string {
	var files []string

	err := filepath.Walk("./cfg", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}

		if !info.IsDir() && filepath.Ext(path) == ".cfg" {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return files
}
