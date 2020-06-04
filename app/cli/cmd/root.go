package cmd

import (
    "fmt"
    "log"

    "github.com/spf13/cobra"

    "github.com/HoMuChen/go-twstock/domain"
    _companyService "github.com/HoMuChen/go-twstock/uc/company"
    _companyRepo    "github.com/HoMuChen/go-twstock/infra/companyRepository/file"
    _companySource  "github.com/HoMuChen/go-twstock/infra/companySource/file"

    _priceService   "github.com/HoMuChen/go-twstock/uc/price"
    _priceRepo      "github.com/HoMuChen/go-twstock/infra/priceRepository/memory"
    _priceSouce     "github.com/HoMuChen/go-twstock/infra/priceHttpSource"

    _followService "github.com/HoMuChen/go-twstock/uc/follow"
)

var companyService  domain.CompanyService
var followService   domain.FollowService
var priceService    domain.PriceService

var rootCmd = &cobra.Command{
  Use:   "go-twstock",
  Short: "taiwan stock realtime price crawler",
  Long:  `taiwan stock realtime price crawler`,
  Run: func(cmd *cobra.Command, args []string) {
     fmt.Println("Hello from go-twstock")
  },
}

func init() {
  companyRepo := _companyRepo.New("/tmp/follows.csv")
  companySource := _companySource.New()
  priceRepo := _priceRepo.New()
  priceHttpSource := _priceSouce.New()

  companyService = _companyService.New(companySource)
  followService = _followService.New(companyRepo)
  priceService = _priceService.New(priceHttpSource, priceRepo, companyRepo)

}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    log.Fatal(err)
  }
}
