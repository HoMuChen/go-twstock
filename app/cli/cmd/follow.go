package cmd

import (
    "log"
    "fmt"

    "github.com/spf13/cobra"
)

var followCmd = &cobra.Command{
    Use: "follow",
    Short: "follow {id}",
    Long: `Follow compainy with specific ID`,
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        id := args[0]

        company, err := companyService.GetById(id)
        if err != nil {
            log.Fatal(err)
        }

        err = followService.Follow(company)

        if err != nil {
            log.Fatal(err)
        }

        fmt.Println("Follow success")
    },
}

func init() {
    rootCmd.AddCommand(followCmd)
}
