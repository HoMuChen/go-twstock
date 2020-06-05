package cmd

import (
    "log"
    "fmt"

    "github.com/spf13/cobra"
)

var unfollowCmd = &cobra.Command{
    Use: "unfollow",
    Short: "unfollow {id}",
    Long: `Unfollow compainy with specific ID`,
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        id := args[0]

        company, err := companyService.GetById(id)
        if err != nil {
            log.Fatal(err)
        }

        err = followService.Unfollow(company)

        if err != nil {
            log.Fatal(err)
        }

        fmt.Println("ok")
    },
}

func init() {
    rootCmd.AddCommand(unfollowCmd)
}
