package cmd

import (
    "fmt"
    "log"

    "github.com/spf13/cobra"
)

var from int
var to int

var listCmd = &cobra.Command{
  Use:   "list",
  Short: "List (compainies | follows)",
  Long: `List all compainies or the companies that you followed`,
  Args: cobra.MinimumNArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    if args[0] == "companies" {
       companies, err := companyService.List(from, to)
       if err != nil {
        log.Fatal(err)
      }

      for _, company := range companies {
        fmt.Println(company)
      }
    }

    if args[0] == "follows" {
       companies, err := companyService.ListFollow(from, to)
       if err != nil {
        log.Fatal(err)
      }

      for _, company := range companies {
        fmt.Println(company)
      }
    }
  },
}

func init() {
  listCmd.Flags().IntVarP(&from, "from", "f", 0, "List from")
  listCmd.Flags().IntVarP(&to, "to", "t", 10, "List to")
}
