package estudos

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func Exemplo04() {
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("Exemplo 4 - uso de goroutines e WaitGroup para obter o tamanho de várias urls")
	fmt.Println("------------------------------------------------------------------")

	urls := []string{"https://example.com", "https://golang.org", "https://www.google.com"}
	resultados := make(chan int, len(urls))

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		go obterTamanhoDasUrls(url, resultados, &wg)
	}

	fechar(resultados, &wg)

	for resultado := range resultados {
		fmt.Println(resultado)
	}
}

func obterTamanhoDasUrls(url string, canal chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	resposta, erro := http.Get(url)
	if erro != nil {
		fmt.Println("Erro ao obter o tamanho da url", url)
		return
	}
	fmt.Println("Obtendo o tamanho da url ", url)

	defer resposta.Body.Close()

	conteudo, erro := ioutil.ReadAll(resposta.Body)
	if erro != nil {
		fmt.Println("Erro ao ler o conteúdo da url", url)
		return
	}

	canal <- len(conteudo)
	fmt.Println("Url ", url, " de tamanho")
}

func fechar(canal chan int, wg *sync.WaitGroup) {
	wg.Wait()
	close(canal)
}