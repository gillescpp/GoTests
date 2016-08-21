package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	// *string et byte propose un grand nombre de fonction identiques*
	// (les fonctions de strings ont quasiment systématiquement leurs pendants
	// pour les tableay de byte dans bytes, et vice versa)
	// string : utf8, adapté pour le travail sur texte, immutable (chaqeu modfication,
	// concatenation, ... revient a créer un nouvelle instance de string)
	// [] byte : octet brut, adapté au travail sur flux brut, meilleur erfomance
	// que string, modifiable, reimensionable
	//Note : "Rune" = charactere unicode
	myByteArray := []byte{1, 2, 3, 4, 5, 'x'}
	myByteArray2 := []byte{7, 8, 9, 'y'}
	myString := "Some words here"
	myMultilineSting := `Some word here.
And a new line.`

	// I - Comparaison / Inspection ---------
	if bytes.Equal(myByteArray, myByteArray2) {
		fmt.Println("myByteArray == myByteArray2, comparaison byte à byte")
	}
	if myString == myMultilineSting {
		fmt.Println("myString == myMultilineSting, case sensitive")
	}
	//version non sensible à la casse :
	// > bytes.EqualFold(s, t []byte) bool
	// > strings.EqualFold(s, t string) bool

	//Pour faire du tri à bulle :
	// > bytes.Compare(a, b []byte) int
	// > strings.Compare(a, b string) int
	//retourne 1 si a > b, -1 si b < a et 0 s'ils sont égaux

	//Contains
	// > bytes.Contains(b, subslice []byte) bool
	// > strings.Contains(s, substr string) bool
	if strings.Contains(myString, "Some") {
		fmt.Println("myString contient 'Some'")
	}

	//Tests prefix / sufix
	// > bytes.HasPrefix(s, prefix []byte) bool
	// > strings.HasPrefix(s, prefix string) bool
	// > bytes.HasSuffix(s, suffix []byte) bool
	// > strings.HasSuffix(s, suffix string) bool
	if strings.HasPrefix(myString, "Some") {
		fmt.Println("myString commence par 'Some'")
	}

	//compter un nombre de mot
	// > bytes.Count(s, sep []byte) int
	// > strings.Count(s, sep string) int
	fmt.Println("myString contient", strings.Count(myString, ""), "mot")

	//Position d'un charactere (base à, -1 si non trouvé)
	// > bytes.Index(s, sep []byte) int				//1er trouvé
	// > strings.Index(s, sep string) int
	// > bytes.LastIndex(s, sep []byte) int			//dernier trouvé
	// > strings.LastIndex(s, sep string) int
	fmt.Println("myString contient 'words' à la position", strings.Index(myString, "words"))

	// II - Manipulation basiques ---------
	// > bytes.ToUpper(s []byte) []byte
	// > strings.ToUpper(s string) string
	fmt.Println(strings.ToUpper("En majuscule"))

	// > bytes.ToLower(s []byte) []byte
	// > strings.ToLower(s string) string
	fmt.Println(strings.ToLower("En minuscule"))

	// > bytes.Title(s []byte) []byte
	// > strings.Title(s string) string
	fmt.Println(strings.Title("Une phrase en snake case"))

	// > bytes.TrimSpace(s []byte) []byte
	// > strings.TrimSpace(s string) string
	fmt.Println(strings.TrimSpace("  Trimed string"))

	// > bytes.Trim(s []byte, cutset string) []byte
	// > strings.Trim(s string, cutset string) string
	fmt.Println(strings.Trim("###  Trimed string###", "#"))

	// > bytes.TrimPrefix(s, prefix []byte) []byte
	// > strings.TrimPrefix(s, prefix string) string
	// > bytes.TrimSuffix(s, suffix []byte) []byte
	// > strings.TrimSuffix(s, suffix string) string
	fmt.Println(strings.TrimPrefix("###Trimed prefix###", "###"))
	fmt.Println(strings.TrimSuffix("###Trimed suffix###", "###"))

	// > bytes.Replace(s, old, new []byte, n int) []byte
	// > strings.Replace(s, old, new string, n int) string
	fmt.Println(strings.Replace("Replace all this, this, this", "this", "that", -1))
	fmt.Println(strings.Replace("Replace 2 firsts this, this, this", "this", "that", 2))

	//sous chaine  opérateur [idx char de debut, base 0:idx char de fin] (deb doit être > fin)
	fmt.Println("myString du charactere 0 à 5 ([:6]) :", myString[:6])
	fmt.Println("myString du charactere 3 à 5 ([3:6]) :", myString[3:6])
	fmt.Println("myString du charactere 6 à la fin ([6:]) :", myString[6:])

	// Cas ou on souhaite remplacer plusieurs couple texte source=nouveau texte
	myReplacer := strings.NewReplacer("this", "that", "@FIRSTNAME", "John", "@LASTNAME", "Doe")
	fmt.Println(myReplacer.Replace("Hi @FIRSTNAME @LASTNAME... this is your real name ?"))

	// Cloner
	// > bytes.Join(s [][]byte, sep []byte) []byte
	// > strings.Join(a []string, sep string) string
	fmt.Println("Can you", strings.Repeat("repeat ", 5), "please ?")

	// Concatener
	// > bytes.Join(s [][]byte, sep []byte) []byte
	// > strings.Join(a []string, sep string) string
	myStringArray := []string{"Hi, how", "are you", "my dear ?"}
	fmt.Println(strings.Join(myStringArray, " "))

	// Eclater
	// > bytes.Split(s []byte, sep []byte) []byte
	// > strings.Split(s string, sep string) string
	// > bytes.Fields(s []byte) []byte
	// > strings.Fields(a string) string
	mySplitted := strings.Split("Split me    please", " ")
	fmt.Println(len(mySplitted), "sustrings in 'Split me    please' :", mySplitted)
	myFields := strings.Fields("Split me    please")
	fmt.Println(len(myFields), "words in 'Split me    please' :", myFields)

	// Transformer en Reader ou Writer (sous copie !, on accéde directement
	// aux tableau en mémoire sous-jacent)
	// > bytes.NewReader(s string) Reader
	// > strings.NewReader(s string) Reader
	// permet de transformer une chaine en stream sans copie préalable dans buffer
	r := strings.NewReader("foobar")
	http.Post("http://::1", "text/plain", r)

}
