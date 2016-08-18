package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// I - Lecture d'un fichier, méthode de base
	//Ouverture d'un fichier, en lecture seule par defaut.
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Erreur", err)
	}
	defer file.Close()

	//lecture par block de 100 octet :
	data := make([]byte, 100)
	for err == nil {
		count := 0
		count, err = file.Read(data)
		if err == io.EOF || err == nil {
			fmt.Println("Read", count, "bytes", string(data[:count]))
			//fin du fichier atteint
			if err == io.EOF || count == 0 {
				break
			}

		} else if err != nil {
			fmt.Println("Erreur", err)
		} else {
			// on déplace le curseur pour la prochaine lecture
			// Le premier parametre est l'offset à appliquer, 1 en seconde
			// parametre indique un déplacement relatif à la positon en cours
			// Seek(0,0) permet de revenir au début du reader
			_, err = file.Seek(int64(count), 1)
			if err != nil {
				fmt.Println("Erreur", err)
			}
		}
	}

	// II - ioutils permet de charger simplement
	// tout le fichier en mémoire
	data2, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Erreur", err)
	}
	fmt.Println("ReadFull", string(data2))
}
