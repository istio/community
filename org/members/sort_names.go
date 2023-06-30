package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"gopkg.in/yaml.v2"
)

type Data struct {
	Members []string `yaml:"members"`
}

func main() {
	// Read the YAML file
	yamlFile, err := os.ReadFile("members.yaml")
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
	err = os.WriteFile("members.yaml", updatedYAML, 0644)
	if err != nil {
		log.Fatalf("Failed to write YAML file: %v", err)
	}

	fmt.Println("Names sorted successfully!")
}
