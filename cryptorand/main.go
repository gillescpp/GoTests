package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// Génération d'un tableau de byte de taille s en base64 avec la librairie crypto/rand.
// (qui s'appui sur /dev/urandom/ ou l'api Windows de cryptographie, de maniere
// a generer un nombre aléatoire sécurisé)
// usage : jeton CSRF ou variable de session
func GenerateRandomString(s int) (string, error) {
	b := make([]byte, s)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), err
}

func main() {
	token, err := GenerateRandomString(32)
	if err != nil {.
		fmt.Println("Erreur", err)
	}
	fmt.Println("clé aléatoire : ", token)
}

//links
//	https://elithrar.github.io/article/generating-secure-random-numbers-crypto-rand/
