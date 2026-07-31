package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/kaique/portscanner-go/output"
	"github.com/kaique/portscanner-go/scanner"
	"github.com/kaique/portscanner-go/utils"
)

func main() {
	hostFlag := flag.String("host", "", "Host único ou rede CIDR (ex: 192.168.1.1 ou 192.168.1.0/24)")
	portsFlag := flag.String("ports", "top100", "Portas a escanear (ex: '80,443', '1-1024', 'top100')")
	threadsFlag := flag.Int("threads", 100, "Número de goroutines concorrentes (worker pool)")
	timeoutFlag := flag.Int("timeout", 1000, "Timeout em milissegundos por porta")
	outputFlag := flag.String("output", "", "Arquivo de saída (opcional)")
	formatFlag := flag.String("format", "json", "Formato de saída ('json' ou 'csv')")
	
	flag.Usage = func() {
		utils.PrintBanner()
		fmt.Println("Uso:")
		flag.PrintDefaults()
	}

	flag.Parse()

	utils.PrintBanner()

	if *hostFlag == "" {
		fmt.Println(utils.Red + "Erro: A flag -host é obrigatória!" + utils.Reset)
		flag.Usage()
		os.Exit(1)
	}

	// Parsing Hosts
	hosts, err := utils.ParseHosts(*hostFlag)
	if err != nil {
		fmt.Println(utils.Red + "Erro ao processar host: " + err.Error() + utils.Reset)
		os.Exit(1)
	}

	// Parsing Ports
	ports, err := utils.ParsePorts(*portsFlag)
	if err != nil {
		fmt.Println(utils.Red + "Erro ao processar portas: " + err.Error() + utils.Reset)
		os.Exit(1)
	}

	fmt.Printf(utils.Cyan+"[*] Iniciando scan em %d host(s) e %d porta(s)...\n"+utils.Reset, len(hosts), len(ports))
	fmt.Printf(utils.Cyan+"[*] Usando %d goroutines com timeout de %d ms\n"+utils.Reset, *threadsFlag, *timeoutFlag)

	start := time.Now()
	timeout := time.Duration(*timeoutFlag) * time.Millisecond

	scan := scanner.NewScanner(hosts, ports, *threadsFlag, timeout)
	
	// Executa o scan
	results := scan.Run()

	elapsed := time.Since(start)
	fmt.Printf(utils.Cyan+"\n[*] Scan concluído em %v\n"+utils.Reset, elapsed)

	// Lida com o output
	output.HandleOutput(results, *formatFlag, *outputFlag)
}
