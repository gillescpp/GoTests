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
	tick := time.Tick(100 * time.Millisecond)
	end := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("Tick.")
		case <-end:
			fmt.Println("Fin!")
			return
		default:
			fmt.Println("default")
			time.Sleep(50 * time.Millisecond)
		}
	}
	close(tick)
	close(end)


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
		var batch []int
		iBPos := 0
		for {
			select {
			x := <- ch2 :
				if iBPos > 9 {
					iBPos = 0
					//record batch
					fmt.Println("batch a! ")
				}
				batch[iBPos] = x
				iBPos++

				<-time.After(500*time.Millisecond) :
				//recodd
				fmt.Println("batch t! ")

			}
		}
	}()

	//producteur
	fmt.Println("debut producteur : 200 messages")
	for i := 0; i < 200; i++ {
		//le chan n'ayant qu'un seul emplacement, il est bloqué
		//jusqu' a que le message soit consommé
		ch <- i
	}
	//fin d'emission
	close(ch2)

	//attend que le consomateur aient fini
	fmt.Println("wait")
	wg2.Wait()
}
