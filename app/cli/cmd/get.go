package cmd

import (
    "fmt"
    "log"

    "github.com/spf13/cobra"
)

var id string

var getCmd = &cobra.Command{
  Use:   "get",
  Short: "get company by ID",
  Long: `Get the specific company infomation`,
  Run: func(cmd *cobra.Command, args []string) {
    company, err := companyService.GetById(id)
    if err != nil {
      log.Fatal(err)
    }

    fmt.Println(company)
  },
}

func init() {
  getCmd.Flags().StringVarP(&id, "id", "i", "2330", "company ID")
}
