package estudos

import (
	"fmt"
	"sync"
)

var semaforo = make(chan struct{}, 2)

func Exemplo03() {
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("Exemplo 3 - semaforo")
	fmt.Println("------------------------------------------------------------------")

	var wg sync.WaitGroup

	// Adiciona 3 ao WaitGroup
	wg.Add(3)

	// Goroutines concorrentes usando um semáforo
	for i := 1; i <= 3; i++ {
		go funcao_semaforo(i, &wg)
	}

	// Aguarda a conclusão das goroutines
	wg.Wait()
}

func funcao_semaforo(id int, wg *sync.WaitGroup){
	defer wg.Done()

	fmt.Printf("Goroutine %d aguardando semáforo\n", id)
	semaforo <- struct{}{} // Adquire o semáforo
	defer func() {
		fmt.Printf("Goroutine %d liberando semáforo\n", id)
		<-semaforo // Libera o semáforo
	}()

	fmt.Printf("Goroutine %d está executando\n", id)
}