package estudos

import (
	"fmt"
	"time"
)

func Exemplo01(){
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("Exemplo 1 - Uma goroutine em um canal")
	fmt.Println("------------------------------------------------------------------")

	canal1 := make(chan int)

	go funcao(canal1)

	fmt.Println("iniciou a goroutine")

	for valor := range canal1 {
		fmt.Println("valor recebido na gorotine principal = ", valor)
	}

	fmt.Println("Fim")
}

func funcao(canal chan int) {
	for i:=0; i < 5; i++ {
		time.Sleep(time.Millisecond * 500)
		fmt.Println("valor antes de gravar no canal = ", i)
		canal <- i
		fmt.Println("valor gravado no canal = ", i)
	}

	close(canal)
}