package output

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/ispectr3/portscanner/scanner"
	"github.com/ispectr3/portscanner/utils"
)

// PrintTerminal imprime os resultados coloridos no terminal
func PrintTerminal(results []scanner.Result) {
	if len(results) == 0 {
		fmt.Println(utils.Yellow + "\nNenhuma porta aberta encontrada." + utils.Reset)
		return
	}

	fmt.Println(utils.Green + "\nResultados do Scan:" + utils.Reset)
	fmt.Printf("%-15s %-10s %-10s %s\n", "IP", "PORTA", "ESTADO", "SERVIÇO/BANNER")
	fmt.Println("------------------------------------------------------------")
	for _, r := range results {
		if r.State == "open" {
			fmt.Printf("%-15s %s%-10d%s %s%-10s%s %s\n", r.IP, utils.Green, r.Port, utils.Reset, utils.Green, r.State, utils.Reset, r.Service)
		}
	}
}

// ExportJSON exporta os resultados para um arquivo JSON
func ExportJSON(results []scanner.Result, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(results)
}

// ExportCSV exporta os resultados para um arquivo CSV
func ExportCSV(results []scanner.Result, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"IP", "Porta", "Estado", "Serviço"})
	for _, r := range results {
		writer.Write([]string{r.IP, strconv.Itoa(r.Port), r.State, r.Service})
	}

	return nil
}

// HandleOutput direciona o output dependendo do formato
func HandleOutput(results []scanner.Result, format, outputFile string) {
	PrintTerminal(results)

	if outputFile != "" {
		if format == "json" {
			err := ExportJSON(results, outputFile)
			if err != nil {
				fmt.Printf(utils.Red+"Erro ao exportar JSON: %v\n"+utils.Reset, err)
			} else {
				fmt.Printf(utils.Cyan+"\nResultados exportados para %s (JSON)\n"+utils.Reset, outputFile)
			}
		} else if format == "csv" {
			err := ExportCSV(results, outputFile)
			if err != nil {
				fmt.Printf(utils.Red+"Erro ao exportar CSV: %v\n"+utils.Reset, err)
			} else {
				fmt.Printf(utils.Cyan+"\nResultados exportados para %s (CSV)\n"+utils.Reset, outputFile)
			}
		}
	}
}
