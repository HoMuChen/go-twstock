package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var Company string

var priceCmd = &cobra.Command{
  Use:   "price",
  Short: "show realtime price of certen compainy",
  Long: `price ({id} | all)`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("fetch realtime price with id:", Company)
  },
}

func init() {
  priceCmd.Flags().StringVarP(&Company, "company", "c", "2330", "Comapny to fetch realtime price")
}
