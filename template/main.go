package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type TestContext struct {
	TextDivers string
	RawHtml    template.HTML
	Dog        Dog
	Map        map[string]string
}

type Dog struct {
	Name string
	Age  int
}

func main() {
	// html/template et text/template sont le moteur de template dans la librairie
	// standard. La version html ajoute l'encodage des entités html et la protection
	// contre les injections de code dans la page html. html/template gére de maniere
	// contextuel (contexte determiné automatiquement) le html, le css et le js.
	// A l'usage html/template et text/template sont similaire.

	// Note : certains editeur gere le language de template go si l'extension est ".gohtml"
	fmt.Println("test template")

	// I -
	//hello world :
	// On parse le template (ParseGlob permet de parser directement une chaine)
	// ex : template.ParseGlob("<h1>Hello {{.Name}}!</h1>  ")
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	//Puis on execute le template en lui passant une structure en paramétre
	data := struct {
		Name string
	}{"John Smith"}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	// II -
	//Encodage contextuel
	t, err = template.ParseFiles("context.gohtml")
	if err != nil {
		panic(err)
	}

	fmt.Println("\r\n", "test template avec contexte", "\r\n")
	//template.HTML permet d'indiquer que cette chaine NE SERA PAS encoder
	//(permet d'injecter du html brut)
	//Une chaine de characteres sera protégé selon le contexte (sauf type template.HTML) :
	// - Dans le html, les characteres spéciaux seront encodés en entité html
	// - Dans un lien href, les espaces seront remplacé par %20, etc. pour faire des url valide
	// - Dans du script, les textes seront encodé en utf8, les doubles sont
	//   automatiquement ajoutés sur les chaines
	// - Les objets seront au format json insérable dans un code javascript
	// Note : Par defaut, html/template supprime les commentaires html (<!-- bla -->)
	// présent dans le template
	datac := TestContext{
		TextDivers: "Texte divers éô<>",
		RawHtml:    template.HTML("<h1>Html Brute</h1>"),
		Dog:        Dog{"Fido", 6},
		Map: map[string]string{
			"clé":  "valeur",
			"clé2": "valeur 2",
		},
	}

	err = t.Execute(os.Stdout, datac)
	if err != nil {
		panic(err)
	}

	// III -
	//structure conditionnelles et fonctions

	// serveur pour le rendu des templates
	http.HandleFunc("/1", handler1)
	http.HandleFunc("/2", handler2)

	// le 1er template est le template principale suivi des templates dont il dépand
	testTemplateCond1, err = template.ParseFiles("cond1.gohtml", "foot.gohtml")
	if err != nil {
		panic(err)
	}
	testTemplateCond2, err = template.ParseFiles("cond2.gohtml", "foot.gohtml")
	if err != nil {
		panic(err)
	}

	fmt.Println("\r\n", "serveur http sur port 3000...", "\r\n")
	http.ListenAndServe(":3000", nil)
}

// III -
//structure conditionnelles
var testTemplateCond1 *template.Template
var testTemplateCond2 *template.Template
var testCond1RequestCnt int
var testCond2RequestCnt int

//structure de donnée transmi au templa cond1
type ViewData1 struct {
	Name        string
	StringArray []string
	DogArray    []Dog
}

//structure de donnée transmi au templa cond2
type ViewData2 struct {
	Login     string
	Name      string
	Admin     bool
	Level     int
	ArrayLink []string
}

//page template avec conditions
func handler1(w http.ResponseWriter, r *http.Request) {
	testCond1RequestCnt++ //note : devrait être protegé par un mutex ou incrementé par atomic.AddInt64

	vd := ViewData1{
		Name:        "John Smith",
		StringArray: []string{"item 1", "item 2", "item 3", "item 4"},
		DogArray: []Dog{
			{Name: "Rex", Age: 8},
			{Name: "Lisa", Age: 3},
		},
	}
	// on met des données qui change une fois sur deux
	if testCond1RequestCnt%2 == 0 {
		vd.Name = ""
	}

	w.Header().Set("Content-Type", "text/html")
	err := testTemplateCond1.Execute(w, vd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//page template avec fonctions
func handler2(w http.ResponseWriter, r *http.Request) {
	testCond2RequestCnt++ //note : devrait être protegé par un mutex ou incrementé par atomic.AddInt64

	vd := ViewData2{}
	// on met des données qui change a chaque refresh
	if testCond2RequestCnt == 1 {
		vd = ViewData2{
			Login:     "TheAdmin",
			Name:      "IT Team",
			Admin:     true,
			Level:     3,
			ArrayLink: []string{"Link1", "", "Link2", "Link3"},
		}
	} else if testCond2RequestCnt == 2 {
		vd = ViewData2{
			Login:     "puser3",
			Name:      "Modo!",
			Admin:     false,
			Level:     3,
			ArrayLink: []string{"My test 1", "My test 21", "My test 61", "My test 8"},
		}
	} else if testCond2RequestCnt == 3 {
		vd = ViewData2{
			Login:     "advuser2",
			Name:      "John Adv",
			Admin:     false,
			Level:     2,
			ArrayLink: []string{"Cassandra", "redis", "couchdb", "influxdb", "mysql", "mssql"},
		}

	} else {
		testCond2RequestCnt = 0
		vd = ViewData2{
			Login:     "newb1",
			Name:      "coco 1",
			Admin:     false,
			Level:     1,
			ArrayLink: []string{"MyLink", ""},
		}
	}

	w.Header().Set("Content-Type", "text/html")
	err := testTemplateCond2.Execute(w, vd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
