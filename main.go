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

// type Intro struct {
// 	PageTitle string `yaml:"pagetitle"`
// 	Intro     struct {
// 		Title1 string `yaml:"title1"`
// 		Title2 string `yaml:"title2"`
// 	} `yaml:"intro"`
// 	Navbar      []string `yaml:"navbar"`
// 	Brand       string   `yaml:"brand"`
// 	Occupations []string `yaml:"occupations"`
// 	Sliderimage []string `yaml:"sliderimage"`
// 	Button      struct {
// 		Enable string `yaml:"enable"`
// 		Lable  string `yaml:"label"`
// 		Link   string `yaml:"link"`
// 	} `yaml:"button"`
// }

type Intro struct {
<<<<<<< HEAD
	Pagetitle string   `yaml:"pagetitle"`
=======
	PageTitle string   `yaml:"pagetitle"`
>>>>>>> d413102... minor updates
	Navbar    []string `yaml:"navbar"`
	Brand     string   `yaml:"brand"`
	Intro     struct {
		Title1 string   `yaml:"title1"`
		Title2 string   `yaml:"title2"`
		Title3 []string `yaml:"title3"`
	} `yaml:"intro"`
<<<<<<< HEAD
	Recipes    []string `yaml:"recipes"`
	RecipeList []struct {
		Recipe          string   `yaml:"recipe"`
		PreparationTime string   `yaml:"preparation_time"`
		CookingTime     string   `yaml:"cooking_time"`
		Difficulty      string   `yaml:"difficulty,omitempty"`
=======
	RecipeList []string `yaml:"recipe_list"`
	Recipes    []struct {
		Recipe          string   `yaml:"recipe"`
		PreparationTime string   `yaml:"preparation_time"`
		CookingTime     string   `yaml:"cooking_time"`
		Difficulty      string   `yaml:"difficulty"`
>>>>>>> d413102... minor updates
		Ingredients     []string `yaml:"ingredients"`
		Method          string   `yaml:"method"`
		Tags            []string `yaml:"tags"`
		PhotoURL        string   `yaml:"photo_url"`
		VideoURL        string   `yaml:"video_url"`
<<<<<<< HEAD
		Ifficulty       string   `yaml:"ifficulty,omitempty"`
	} `yaml:"recipe_list"`
=======
	} `yaml:"recipes"`
>>>>>>> d413102... minor updates
}

// type Intro struct {
// 	Pagetitle string   `yaml:"pagetitle"`
// 	Navbar    []string `yaml:"navbar"`
// 	Brand     string   `yaml:"brand"`
// 	Intro     struct {
// 		Title1 string   `yaml:"title1"`
// 		Title2 string   `yaml:"title2"`
// 		Title3 []string `yaml:"title3"`
// 	} `yaml:"intro"`
// 	RecipeList    []string `yaml:"recipe_list"`
// 	Recipes []struct {
// 		Recipe          string   `yaml:"recipes"`
// 		PreparationTime string   `yaml:"preparation_time"`
// 		CookingTime     string   `yaml:"cooking_time"`
// 		Difficulty      string   `yaml:"difficulty,omitempty"`
// 		Ingredients     []string `yaml:"ingredients"`
// 		Method          string   `yaml:"method"`
// 		Tags            []string `yaml:"tags"`
// 		PhotoURL        string   `yaml:"photo_url"`
// 		VideoURL        string   `yaml:"video_url"`
// 	} `yaml:"recipe_list"`
// }
