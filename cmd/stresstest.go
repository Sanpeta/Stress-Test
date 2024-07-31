/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/Sanpeta/stress-test-pos-go-expert/internal/stresstest"
	"github.com/spf13/cobra"
)

var (
	urlStress          string
	initialConcurrency int
	maxConcurrency     int
	increment          int
)

// stresstestCmd represents the stresstest command
var stressTestCmd = &cobra.Command{
	Use:   "stresstest",
	Short: "Executa testes de estresse em um serviço web",
	Run: func(cmd *cobra.Command, args []string) {
		if urlStress == "" {
			log.Fatal("URL do serviço é obrigatória")
		}
		if initialConcurrency <= 0 || maxConcurrency <= 0 || increment <= 0 {
			log.Fatal("Os valores de concorrência inicial, máxima e incremento devem ser maiores que zero")
		}
		fmt.Printf("Iniciando teste de estresse:\nURL: %s\nConcorrência Inicial: %d\nConcorrência Máxima: %d\nIncremento: %d\n", urlStress, initialConcurrency, maxConcurrency, increment)
		stresstest.StartStressTest(urlStress, initialConcurrency, maxConcurrency, increment)
	},
}

func init() {
	rootCmd.AddCommand(stressTestCmd)
	stressTestCmd.Flags().StringVar(&urlStress, "url", "", "URL do serviço a ser testado (obrigatório)")
	stressTestCmd.Flags().IntVar(&initialConcurrency, "initial-concurrency", 10, "Número inicial de chamadas simultâneas")
	stressTestCmd.Flags().IntVar(&maxConcurrency, "max-concurrency", 100, "Número máximo de chamadas simultâneas")
	stressTestCmd.Flags().IntVar(&increment, "increment", 10, "Incremento na concorrência a cada passo")
	stressTestCmd.MarkFlagRequired("url")
}
