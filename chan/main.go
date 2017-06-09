package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// -------------------------------------------
	//unbuffered chan : conduit d'un seul "slot"
	//L'emetteur' place une donnée dans le channel, le recepteur le recoit
	fmt.Println(" ---- Chan Simple ---- ")
	ch := make(chan int)
	wg := sync.WaitGroup{}

	//on lance une goroutine consomateur
	go func() {
		fmt.Println("debut consommateur")
		for v := range ch { //revois les message jusqu'a fermeture par l'emetteur
			//traitement du message
			fmt.Println("recu : ", v)
			<-time.After(100 * time.Millisecond)

			wg.Add(1)
			defer wg.Done()
		}
	}()

	//producteur
	fmt.Println("debut producteur")
	for i := 0; i < 10; i++ {
		//le chan n'ayant qu'un seul emplacement, il est bloqué
		//jusqu' a que le message soit consommé
		ch <- i
		fmt.Println("emis : ", i)
	}
	//fin d'emission
	close(ch)

	//attend que le consomateur aient fini
	fmt.Println("wait")
	wg.Wait()

	// -------------------------------------------
	//timer
	fmt.Println(" ---- Timer ---- ")
	tick := time.Tick(100 * time.Millisecond) //inutile de fermer ces channel (sinon il faut passer par un NewTicker)
	end := time.After(500 * time.Millisecond)

	for bContinue := true; bContinue; {
		select {
		case <-tick:
			fmt.Println("Tick.")
		case <-end:
			fmt.Println("Fin!")
			bContinue = false
		default:
			fmt.Println("default")
			time.Sleep(50 * time.Millisecond)
		}
	}

	// -------------------------------------------
	//unbuffered chan : conduit d'un seul "slot"
	//L'emetteur' place une donnée dans le channel, le recepteur le recoit
	fmt.Println(" ---- Chan buffered ---- ")
	iBatch := 10
	ch2 := make(chan int, 2*iBatch)
	wg2 := sync.WaitGroup{}

	//on lance une goroutine consomateur
	go func() {
		fmt.Println("debut consommateur")
		var batch []int = make([]int, iBatch)
		iBPos := 0
		for {
			select {
			case x := <-ch2:
				if iBPos > 9 {
					iBPos = 0
					//record batch
					fmt.Println("batch a! ", batch)
				}
				batch[iBPos] = x
				iBPos++
				wg2.Done()

			case <-time.After(200 * time.Millisecond):
				//recodd
				fmt.Println("batch t! ")

			}
		}
	}()

	//producteur
	fmt.Println("debut producteur : 200 messages")
	for i := 0; i < 205; i++ {
		//le chan ayant plusieurs emplacement, il sera bloquant
		//s'il est plein
		wg2.Add(1)
		ch2 <- i
		if i%30 == 0 {
			time.Sleep(300 * time.Millisecond)
		} else {
			time.Sleep(50 * time.Millisecond)
		}
	}
	//attend la fin
	wg2.Wait()
	close(ch2)

}
