package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.html"))
}

func main() {
	http.HandleFunc("/", index)
	// The following two lines of code allows go to serve the static asset files
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	// Listen and Serve
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	var test Intro
	yamlFile, err := ioutil.ReadFile("./assets/structures/homepage.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &test)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	tpl.ExecuteTemplate(w, "index.html", test)
}

type Intro struct {
	PageTitle string `yaml:"pagetitle"`
	Intro     struct {
		Title1 string `yaml:"title1"`
		Title2 string `yaml:"title2"`
	} `yaml:"intro"`
	Navbar      []string `yaml:"navbar"`
	Brand       string   `yaml:"brand"`
	Occupations []string `yaml:"occupations"`
	Sliderimage []string `yaml:"sliderimage"`
	Button      struct {
		Enable string `yaml:"enable"`
		Lable  string `yaml:"label"`
		Link   string `yaml:"link"`
	} `yaml:"button"`
}
