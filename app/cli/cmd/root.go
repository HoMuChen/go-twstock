package cmd

import (
    "fmt"
    "log"

    "github.com/spf13/cobra"

    "github.com/HoMuChen/go-park/domain"
    _companyService "github.com/HoMuChen/go-park/uc/company"
    _companyRepo    "github.com/HoMuChen/go-park/infra/companyRepository/redis"
    _companySource  "github.com/HoMuChen/go-park/infra/companySource/file"

    _priceService   "github.com/HoMuChen/go-park/uc/price"
    _priceRepo      "github.com/HoMuChen/go-park/infra/priceRepository/memory"
    _priceSouce     "github.com/HoMuChen/go-park/infra/priceHttpSource"

    _followService "github.com/HoMuChen/go-park/uc/follow"
)

var companyService  domain.CompanyService
var followService   domain.FollowService
var priceService    domain.PriceService

var rootCmd = &cobra.Command{
  Use:   "go-park",
  Short: "taiwan stock realtime price crawler",
  Long:  `taiwan stock realtime price crawler`,
  Run: func(cmd *cobra.Command, args []string) {
     fmt.Println("Hello from go-park")
  },
}

func init() {
  companySource := _companySource.New("./data/companies.csv")
  companyService = _companyService.New(companySource)

  companyRepo := _companyRepo.New()
  followService = _followService.New(companyRepo)

  priceRepo := _priceRepo.New()
  priceHttpSource := _priceSouce.New()
  priceService = _priceService.New(priceHttpSource, priceRepo, companyRepo)

}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    log.Fatal(err)
  }
}
