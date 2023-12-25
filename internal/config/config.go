package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Host string `yaml:"host"`
}

func New() *Config {
	cfg := Config{}
	pwd, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = cleanenv.ReadConfig(filepath.Dir(pwd)+"/fart.yaml", &cfg)
	if err != nil {
		log.Fatal("cannot parse config file: fart.yaml")
	}
	return &cfg
}

func Write(host string) {
	cfg := Config{Host: host}
	yamlFile, err := yaml.Marshal(cfg)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(yamlFile))
	pwd, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(filepath.Dir(pwd))
	f, err := os.Create(filepath.Dir(pwd) + "/fart.yaml")
	if err != nil {
		fmt.Printf("cannot create config file at path: %s, because error is: %s", filepath.Dir(pwd), err)
	}
	defer f.Close()

	_, err = io.WriteString(f, string(yamlFile))
	if err != nil {
		fmt.Printf("cannot write in the fart.yaml because error is: %s", err)
	}
}
