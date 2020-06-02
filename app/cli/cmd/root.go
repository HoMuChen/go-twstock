package cmd

import (
    "fmt"
    "log"

    "github.com/spf13/cobra"

    "github.com/HoMuChen/go-park/domain"
    _companyService "github.com/HoMuChen/go-park/uc/company"
    _companyRepo    "github.com/HoMuChen/go-park/infra/companyRepository/redis"
    _companySource  "github.com/HoMuChen/go-park/infra/companySource/file"
)

var companyService domain.CompanyService

var rootCmd = &cobra.Command{
  Use:   "go-park",
  Short: "taiwan stock realtime price crawler",
  Long:  `taiwan stock realtime price crawler`,
  Run: func(cmd *cobra.Command, args []string) {
     fmt.Println("Hello from go-park")
  },
}

func init() {
  comapnyRepo := _companyRepo.New()
  companySource := _companySource.New("./data/companies.csv")
  companyService = _companyService.New(comapnyRepo, companySource)

  rootCmd.AddCommand(listCmd)
  rootCmd.AddCommand(getCmd)
  rootCmd.AddCommand(priceCmd)
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    log.Fatal(err)
  }
}
