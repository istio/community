package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v2"
)

type Data struct {
	Members []string `yaml:"members"`
}

func main() {
	// Get the current directory
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current directory: %v", err)
	}

	// Define the path to the members.yaml file
	membersFilePath := filepath.Join(currentDir, "tempname.yaml")

	// Read the YAML file
	yamlFile, err := ioutil.ReadFile(membersFilePath)
	if err != nil {
		log.Fatalf("Failed to read YAML file: %v", err)
	}

	// Unmarshal YAML data
	var data Data
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
	}

	// Sort the names alphabetically (case-insensitive)
	sort.Slice(data.Members, func(i, j int) bool {
		return strings.ToLower(data.Members[i]) < strings.ToLower(data.Members[j])
	})

	// Marshal YAML data
	updatedYAML, err := yaml.Marshal(&data)
	if err != nil {
		log.Fatalf("Failed to marshal YAML: %v", err)
	}

	// Write the updated data back to the YAML file
	err = ioutil.WriteFile(membersFilePath, updatedYAML, 0644)
	if err != nil {
		log.Fatalf("Failed to write YAML file: %v", err)
	}

	// Create a copy of members.yaml in the parent directory
	parentDir := filepath.Dir(currentDir)
	copyFilePath := filepath.Join(parentDir, "members.yaml")
	err = ioutil.WriteFile(copyFilePath, updatedYAML, 0644)
	if err != nil {
		log.Fatalf("Failed to create a copy of members.yaml: %v", err)
	}

	fmt.Println("Names sorted successfully!")
}

