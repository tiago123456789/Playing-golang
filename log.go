package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	http "net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const quantidadeVezExecutarMonitoramento = 3
const delayVerificacao = 1 * time.Minute

func main() {
	displayMessageIntroduction()

	for {
		displayMenu()
		executeTask(getCommandTyping())
	}
}

func getCommandTyping() int {
	var command int
	fmt.Scan(&command)
	return command
}

func displayMessageIntroduction() {
	version := 1.0
	fmt.Println("Version software:", version)
}

func displayMenu() {
	fmt.Println("################## MENU ############################")
	fmt.Println("1 - Iniciar monitoramento.")
	fmt.Println("2 - Exibir logs.")
	fmt.Println("3 - Sair programa.")
}

func executeTask(numberTyping int) {
	switch numberTyping {
	case 1:
		monitorar()
	case 2:
		displayLogs()
	case 3:
		fmt.Println("Sair programa.")
		os.Exit(0)
	default:
		os.Exit(-1)
	}
}

func monitorar() {
	sites := getSites()

	fmt.Println("Iniciar monitoramento....")

	for indice := 0; indice < quantidadeVezExecutarMonitoramento; indice++ {
		for _, site := range sites {
			checkSite(site)
		}
		time.Sleep(delayVerificacao)
		fmt.Println("")
	}
}

func checkSite(site string) {
	response, _ := http.Get(site)

	if response.StatusCode != 200 {
		fmt.Println("O", site, "está fora do ar!")
		registerLog(site, false)
	} else {
		fmt.Println("O", site, "está funcionando!")
		registerLog(site, true)
	}
}

func getSites() []string {
	arquivo, erro := os.Open("sites.txt")
	fmt.Println(erro)
	sites := []string{}
	leitorArquivo := bufio.NewReader(arquivo)
	for {
		linha, err := leitorArquivo.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}

func registerLog(site string, onlineOuOffline bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Occurred a error:", err)
	}

	message := time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online:" + strconv.FormatBool(onlineOuOffline) + "\n"
	arquivo.WriteString(message)
	arquivo.Close()
}

func displayLogs() {
	fmt.Println("Exibir logs ....")
	file, err := ioutil.ReadFile("log.txt")

	if err == nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}
