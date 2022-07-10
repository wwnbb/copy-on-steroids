package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/wwnbb/go-cp/cp"
)

var (
	offset int
	limit  int
)

var rootCmd = &cobra.Command{
	Use:   "go-cp",
	Short: "Unix cp clone",
	Long:  `Copy file with offset and limit and progress bar`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		source := args[0]
		dest := args[1]
		cp.CopyFile(source, dest, offset, limit)
	},
}

func init() {
	rootCmd.Flags().IntVar(&offset, "offset", 0, "Offset")
	rootCmd.Flags().IntVar(&limit, "limit", 0, "Limit")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
