package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var priceCmd = &cobra.Command{
  Use:   "price",
  Short: "show realtime price of certen compainy",
  Long: `price ({id} | follows)`,
  Run: func(cmd *cobra.Command, args []string) {
    if args[0] == "all" {
      prices, _ := priceService.FetchRealtimeAll(0, 10)

      for _, price := range prices {
        fmt.Println(price.Company.ID, price.Company.Name, price.Value)
      }
    } else {
      company, _ := companyService.GetById(args[0])
      price, _ := priceService.FetchRealtime(company)

      fmt.Println(price.Company.ID, price.Company.Name, price.Value)
    }
  },
}

func init() {
  rootCmd.AddCommand(priceCmd)
}
