package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/fatih/color"
)

func main() {
	// Create slices for active and stale projects
	active := []string{}
	stale := []string{}

	// Get the root directory
	root, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Open the root directory
	dir, err := os.Open(root)
	if err != nil {
		log.Fatal(err)
	}
	defer dir.Close()

	// Read the contents of the root directory
	files, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}

	// Calculate the threshold date (30 days ago)
	thresholdDate := time.Now().AddDate(0, 0, -30)

	// Iterate over the directory entries
	for _, file := range files {
		// Check if the entry is a directory (that is not hidden)
		if file.IsDir() && file.Name()[0] != '.' {
			// Get the modification time of the directory
			modTime := file.ModTime()

			// Check if the modification time is before the threshold date
			if modTime.Before(thresholdDate) {
				stale = append(stale, file.Name())
			} else {
				active = append(active, file.Name())
			}
		}
	}

	// Sort the slices
	sort.Strings(active)
	sort.Strings(stale)

	// Format the Strings
	act := strings.Join(active, ", ")
	actives := fmt.Sprintf("🟢 Active projects: %s", act)
	stl := strings.Join(stale, ", ")
	stales := fmt.Sprintf("🟡 Stale projects: %s", stl)

	//Add colors (courtesy of Fatih's color pkg)
	color.Green(actives)
	fmt.Println("")
	color.Yellow(stales)
}
