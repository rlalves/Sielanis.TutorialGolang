package estudos

import (
	"time"
	"fmt"
	"sync"
)

func Exemplo02(){
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("Exemplo 2 - n goroutines em um canal")
	fmt.Println("------------------------------------------------------------------")
	meuCanal := make(chan int)
	var wg sync.WaitGroup

	// Adiciona 2 ao WaitGroup (uma para o produtor e outra para o consumidor)
	wg.Add(2)

	// Goroutine para o produtor
	go produtor(meuCanal, &wg)
	fmt.Println("finalizou a goroutine produtora")

	// Goroutine para o consumidor
	go consumidor(meuCanal, &wg)
	fmt.Println("finalizou a goroutine consumidora")

	// Aguarda a conclusão das goroutines
	wg.Wait()
}

func produtor(canal chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 500)
		canal <- i
	}
	close(canal)
}

func consumidor(canal chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for valor := range canal {
		fmt.Println("Valor recebido:", valor)
	}
}