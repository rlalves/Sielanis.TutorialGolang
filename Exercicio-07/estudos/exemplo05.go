package estudos

import (
	"fmt"
	"sync"
)

func Exemplo05() {
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("Exemplo 5 - uso de goroutines e WaitGroup - evolução do exemplo 4")
	fmt.Println("------------------------------------------------------------------")

	urls_variadas := []string{"https://example.com", "https://golang.org", "https://www.google.com"}
	urls_ti := []string{"https://microsoft.com"}
	urls_bancos := []string{"https://itau.com.br", "https://santander.com.br"}
	fmt.Println("montou as listas")
	
	resultados_v := make(chan int, len(urls_variadas))
	resultados_t := make(chan int, len(urls_ti))
	resultados_b := make(chan int, len(urls_bancos))
	fmt.Println("criou os canais")

	var wg_v sync.WaitGroup
	wg_v.Add(len(urls_variadas))

	var wg_t sync.WaitGroup
	wg_t.Add(len(urls_ti))

	var wg_b sync.WaitGroup
	wg_b.Add(len(urls_bancos))
	fmt.Println("criou os WaitGroups")

	for _, urls_variadas := range urls_variadas {
		go obterTamanhoDasUrls(urls_variadas, resultados_v, &wg_v)
	}

	for _, urls_ti := range urls_ti {
		go obterTamanhoDasUrls(urls_ti, resultados_t, &wg_t)
	}

	for _, urls_bancos := range urls_bancos {
		go obterTamanhoDasUrls(urls_bancos, resultados_b, &wg_b)
	}
	fmt.Println("iniciou as goroutines")

	fechar(resultados_v, &wg_v)
	fechar(resultados_t, &wg_t)
	fechar(resultados_b, &wg_b)
	fmt.Println("fechou os canais")

	for resultado := range resultados_v {
		fmt.Println(resultado)
	}
	for resultado := range resultados_t {
		fmt.Println(resultado)
	}
	for resultado := range resultados_b {
		fmt.Println(resultado)
	}
}

