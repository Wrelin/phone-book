package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Mat struct {
	DockerImage string    `yaml:"docker"`
	Version     []float32 `yaml:",flow"`
}

type YAML struct {
	Image  string
	Matrix Mat
}

func main() {
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	data := YAML{}

	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("After Unmarshal (Structure):\n%v\n\n", data)

	d, err := yaml.Marshal(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("After Marshal (YAML code):\n%s\n", string(d))
}
