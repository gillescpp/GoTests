package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"sync/atomic"

	"errors"

	"golang.org/x/sync/errgroup"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Le paquage errgroup permet principalement a gestion des erreurs dans les
	// goroutines lancés en paralelle pour traiter un même tache.

	//Exemple de 3 goroutines gérés avec un waitgroup classique
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1) //incrémentation waitgroup
		// lancement de la goroutine
		go func(worknum int) {
			defer wg.Done() //décrémentation waitgroup en fin de traitement
			// Simulation d'un travail d'un durée variable
			workDuration := time.Duration(1000+rand.Intn(2500)) * time.Millisecond
			time.Sleep(workDuration)
			fmt.Println("WaitGroup : work", worknum, "done in ", workDuration)
		}(i)
	}
	// on attend la fin des goroutines
	wg.Wait()
	fmt.Println("WaitGroup terminé")

	// errgroup permet de faire la même en y intégrant la gestion d'erreur dans
	// la goroutine
	var eg errgroup.Group
	var aworknum int64
	for i := 0; i < 3; i++ {
		//pas de système d'incrémentation ici, on fouri la goroutine au errgroup
		eg.Go(func() error {
			// Simulation d'un travail d'un durée variable
			worknum := atomic.AddInt64(&aworknum, 1)
			workDuration := time.Duration(1000+rand.Intn(2500)) * time.Millisecond
			time.Sleep(workDuration)
			fmt.Println("errgroup : work", worknum, "done in ", workDuration)
			//simulation d'une erreur au travail : la 1ere erreur anule
			//le groupe (mais les go routine en cours ne sont pas arrétés).
			//la derniere erreur retourné sera a son tour retourné par wait
			if worknum > 1 {
				return errors.New(fmt.Sprintf("errgroup : work %v fail", worknum))
			} else {
				return nil
			}
		})
	}
	// on attend la fin des goroutines, et on regarde si une a échoué
	if err := eg.Wait(); err != nil {
		fmt.Println("errgroup terminé avec erreur", err)
	} else {
		fmt.Println("errgroup terminé")
	}

	//TODO : voir func errgroup.WithContext(ctx context.Context) (*Group, context.Context)

}

// Links :
//	https://godoc.org/golang.org/x/sync/errgroup
