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
	recent := []string{}
	stale := []string{}
	deprecated := []string{}

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
	recentDate := time.Now().AddDate(0, 0, -30)
	staleDate := time.Now().AddDate(0,0,-60)
	deprecatedDate := time.Now().AddDate(0,0,-90)

	
	// Iterate over the directory entries
	for _, file := range files {
		// Check if the entry is a directory (that is not hidden)
		if file.IsDir() && file.Name()[0] != '.' {
			// Get the modification time of the directory
			modTime := file.ModTime()

            // Check if the modification time is after the stale date
            if modTime.After(staleDate) {
				active = append(active, file.Name())
			} else if modTime.After(recentDate) {
				recent = append(recent, file.Name())
			} else if modTime.After(deprecatedDate) {
				stale = append(stale, file.Name())
			} else {
				deprecated = append(deprecated, file.Name())
			}
		}
	}

	// Sort the slices
	sort.Strings(active)
	sort.Strings(recent)
	sort.Strings(stale)
	sort.Strings(deprecated)

	// Format the Strings
	act := strings.Join(active, ", ")
	actives := fmt.Sprintf("🟢 %v Active projects: %s", len(active),act)
	rct := strings.Join(recent, ", ")
	recents := fmt.Sprintf("🟡 %v Recent projects: %s", len(recent), rct)
	stl := strings.Join(stale, ", ")
	stales := fmt.Sprintf("🔴 %v Stale projects: %s",len(stale), stl)
	dep := strings.Join(stale, ", ")
	depres := fmt.Sprintf("⚫ %v Deprecated projects: %s", len(deprecated), dep)
	
	//Add colors (courtesy of Fatih's color pkg)
	color.Green(actives)
	color.Yellow(recents)
	color.Red(stales)
	color.HiBlack(depres)
}
