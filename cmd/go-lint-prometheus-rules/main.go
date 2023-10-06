package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type Config struct {
	ErrorsAndDowntime ErrorsAndDowntime `yaml:"errors-and-downtime"`
}

type ErrorsAndDowntime struct {
	Groups []Group `yaml:"groups"`
}

type Group struct {
	Name  string `yaml:"name"`
	Rules []Rule `yaml:"rules"`
}

type Rule struct {
	Alert      string            `yaml:"alert"`
	Expr       string            `yaml:"expr"`
	For        string            `yaml:"for"`
	Labels     map[string]string `yaml:"labels"`
	Annotations map[string]string `yaml:"annotations"`
}

func main() {
	var filePath string
	flag.StringVar(&filePath, "file", "", "Path to the YAML file")
	flag.Parse()

	if filePath == "" {
		log.Fatal("Please provide the path to the YAML file using the -file flag.")
	}

	yamlData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	var cfg Config
	err = yaml.Unmarshal(yamlData, &cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	for _, group := range cfg.ErrorsAndDowntime.Groups {
		for _, rule := range group.Rules {
			// Check labels
			if _, ok := rule.Labels["owner"]; !ok {
				fmt.Printf("Alert %s is missing 'owner' label\n", rule.Alert)
			}
			if _, ok := rule.Labels["severity"]; !ok {
				fmt.Printf("Alert %s is missing 'severity' label\n", rule.Alert)
			}

			// Check annotations
			if _, ok := rule.Annotations["runbook_url"]; !ok {
				fmt.Printf("Alert %s is missing 'runbook_url' annotation\n", rule.Alert)
			}
			if _, ok := rule.Annotations["description"]; !ok {
				fmt.Printf("Alert %s is missing 'description' annotation\n", rule.Alert)
			}
			if _, ok := rule.Annotations["summary"]; !ok {
				fmt.Printf("Alert %s is missing 'summary' annotation\n", rule.Alert)
			}
		}
	}
}

