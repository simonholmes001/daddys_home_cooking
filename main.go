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
	http.HandleFunc("/ratatouille", ratatouille)
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

// THIS WAS THE LAST ONE WORKING
type Intro struct {
	PageTitle string   `yaml:"pagetitle"`
	Navbar    []string `yaml:"navbar"`
	Brand     string   `yaml:"brand"`
	Intro     struct {
		Title1 string   `yaml:"title1"`
		Title2 string   `yaml:"title2"`
		Title3 []string `yaml:"title3"`
	} `yaml:"intro"`
	RecipeList []string `yaml:"recipe_list"`
	NewRecipes []struct {
		Recipe            string   `yaml:"recipe"`
		PreparationTime   string   `yaml:"preparation_time"`
		CookingTime       string   `yaml:"cooking_time"`
		Difficulty        string   `yaml:"difficulty"`
		Ingredients       []string `yaml:"ingredients"`
		Method            string   `yaml:"method"`
		Tags              []string `yaml:"tags"`
		PhotoURL          string   `yaml:"photo_url"`
		VideoURL          string   `yaml:"video_url"`
		RatatouilleEnable bool     `yaml:"ratatouille_enable"`
		Homo2Enable       bool     `yaml:"homo2_enable"`
		Homo3Enable       bool     `yaml:"homo3_enable"`
	} `yaml:"new_recipes"`
}

// recipe_templates

func ratatouille(w http.ResponseWriter, r *http.Request) {
	var test Intro
	yamlFile, err := ioutil.ReadFile("./assets/structures/homepage.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &test)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	tpl.ExecuteTemplate(w, "ratatouille.html", test)
}
