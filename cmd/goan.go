package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-annotate/pkg/goan"
)

var goanCmd = &cobra.Command{
	Use:     "goan",
	Aliases: []string{"goan"},
	Short:   "GenerateSwaggerAnnotations",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := goan.GenerateSwaggerAnnotations(args[0])
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(goanCmd)
}
