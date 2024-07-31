package cmd

import (
	"fmt"
	"log"

	"github.com/Sanpeta/stress-test-pos-go-expert/internal/loadtest"
	"github.com/spf13/cobra"
)

var (
	urlLoad     string
	requests    int
	concurrency int
)

// loadTestCmd represents the loadtest command
var loadTestCmd = &cobra.Command{
	Use:   "loadtest",
	Short: "Executa testes de carga em um serviço web",
	Run: func(cmd *cobra.Command, args []string) {
		if urlLoad == "" {
			log.Fatal("URL do serviço é obrigatória")
		}
		fmt.Printf("Iniciando teste de carga:\nURL: %s\nTotal de Requests: %d\nConcorrência: %d\n", urlLoad, requests, concurrency)
		loadtest.StartLoadTest(urlLoad, requests, concurrency)
	},
}

func init() {
	rootCmd.AddCommand(loadTestCmd)
	loadTestCmd.Flags().StringVar(&urlLoad, "url", "", "URL do serviço a ser testado (obrigatório)")
	loadTestCmd.Flags().IntVar(&requests, "requests", 100, "Número total de requests")
	loadTestCmd.Flags().IntVar(&concurrency, "concurrency", 10, "Número de chamadas simultâneas")
	loadTestCmd.MarkFlagRequired("url")
}
