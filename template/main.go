package main

import (
	"fmt"
	"html/template"
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
	//structure conditionnelles

	// todo p2

}
