package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	f1         *string
	f2         *bool
	f3         *int
	f4         *float64
	a1         string
	a2         string
	helpWanted *bool
	// >fSub1 *flag.FlagSet
	// >fSub2 *flag.FlagSet
)

func init() {
	//Le package flag permet de lire les parametres passés a un programme
	//en ligne de commande.
	//Les "flags" peuvent être défini dans l'init du package main, mais la lecture
	//des paramètres passés au programme lors du lancement doit forcement être
	//faite dans le main (appel de flag.Parse())
	//Vocabulaire : "./monappli.exe --flag=valeur --autreflag=10 argument1 argument2"

	//Exemple simple : ./monapplication -truc=machin
	//Note : "./monapplication --truc=machin" est aussi valide
	f1 = flag.String("truc", "Valeur par defaut!", "Le truc à traiter")
	//Cas d'un flag type présent/non présent : ./monapplication -verbose
	f2 = flag.Bool("verbose", false, "Mode verbeux activé")
	//un nombre entier ou décimal :
	f3 = flag.Int("qty", 0, "Quantité de truc")
	f4 = flag.Float64("price", 0, "Prix du truc")
	//flag --help si l'utilisateur veut voir comment se servir du programme
	helpWanted = flag.Bool("help", false, "aide")

	// Tous ces pointeurs seront ensuite renseignés lors par l'appel à "Parse"

	// ----------------------
	// Note : Cas des sous commandes
	// Si un même programme a plusieurs cas d'usage qu'on spécifie comme suit :
	// ./monappli.exe Analyse fichier1   #Sous commande pour lancer une analyse
	// ./monappli.exe Traite --verbose fichier     #Sous commande pour traiter
	// Chaque sous commandes peut avoir des flags/arguments propres.
	// Pour gerer ces cas, on crée des flags spécifique pour chaque sous commande :
	// >fSub1 = flag.NewFlagSet("Analyse", flag.ExitOnError)
	// >fSub2 = flag.NewFlagSet("Traite", flag.ExitOnError)
	// Puis on renseigne pour chacun les flags de maniere classique
	// >fS2 = fSub2.Bool("verbose", false, "Mode verbeux activé")
	//etc.
	// Puis dans le *main*, il faudra alors s'appuyer sur os.Args[1] pour savoir
	// quel sous commandes a été utilisé :
	// >switch os.Args[1] {
	// >case "Analyse":
	// >    fSub1.Parse(os.Args[2:])
	// >case "Traite":
	// >    fSub2.Parse(os.Args[2:])
	// >default:
	// >    flag.PrintDefaults()
	// >    os.Exit(1)
	// >}
	// >
	// >if fSub1.Parsed() {
	// >	//...
	// >} else if fSub2.Parsed() {
	// >	//...
	// >}
	// ----------------------
}

func main() {
	//Lecture des flags
	flag.Parse()

	//Aprés les flags (elements avec tiret comme ci-dessus, peuvent suivrent des
	//arguments classiques récupérables dans l'ordre ou ils sont fournis.
	//Note : Ceux ne sont forcement récupérable qu'aprés l'appel à Parse!
	//Si l'index spécifié n'existe pas, la fonction Arg() retourne vide
	a1 = flag.Arg(0)
	a2 = flag.Arg(1)

	// si un ou plusieurs flag sont obligatoire, c'est à nous de vérifier si celui-ci
	// n'est pas resté à la valeur par defaut.
	// On peut a alors rappeler à l'utilisateur comme il doit appeler notre appication.
	// Par exemple, si "Quantité" n'est pas spécifié, ou contient un valeur inadéquate,
	// on affiche la notice d'utilisation :
	if *f3 <= 0 || *helpWanted {
		flag.PrintDefaults()
		os.Exit(1)
	}
	fmt.Printf("Flag 'truc' = %v, 'verbose' = %v, 'qty' = %v, 'price' = %v \r\n", *f1, *f2, *f3, *f4)
	fmt.Println("Arguments : ", a1, a2)

}

// links :
// - https://golang.org/pkg/flag/
// - https://blog.komand.com/build-a-simple-cli-tool-with-golang
