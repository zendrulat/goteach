package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Model struct {
	Names  []string
	Whys   []string
	Titles []string
}

/*templates stuff
************************************************************* */
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

/*Main!
************************************************************* */
func main() {
	/*routes
	************************************************************* */
	http.HandleFunc("/", index)
	http.HandleFunc("/golang", golang)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/gocli", gocli)
	//	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/gohelpfuls", gohelpfuls)
	http.HandleFunc("/tempdata", tempdata)
	http.HandleFunc("/linux", linux)
	http.HandleFunc("/resources", resources)

	http.ListenAndServe(":8080", nil)
	var a = "whatever"
	var b = 33
	var c = []int{33, 55, 66}
	var d = []string{"june", "may", "october"}
	icheck(a)
	icheck(b)
	icheck(c)
	icheck(d)

}

/*responses
************************************************************* */
func index(res http.ResponseWriter, req *http.Request) {

	//array

	//serving the css file
	//http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))

	//actual file
	err := tpl.ExecuteTemplate(res, "index.html", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}
func golang(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "golang.html", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}
func signup(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "signup.html", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}
func tempdata(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "tempdata.html", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}
func gocli(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "gocli.html", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}
func gohelpfuls(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "gohelpfuls.html", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}
func linux(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "linux.html", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}
func resources(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "resources.html", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}
func loginPage(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "login.html", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}

func icheck(c interface{}) {
	switch v := c.(type) {

	case string:
		fmt.Println("string", "[1:]", v[1:], "[:1]", v[:1], v)
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case []string:
		fmt.Println("string array", "[1:]", v[1:], "[:1]", v[:1], v)
	case []int:
		fmt.Println("int array", "[1:]", v[1:], "[:1]", v[:1], v)
	default:
		//	fmt.Println("default", "[1:]", c[1:], "[:1]", c[:1], c)
	}
}
