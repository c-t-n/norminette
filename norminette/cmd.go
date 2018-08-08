package norminette

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "norminette",
	Short: "Norminette - 42 Norm Linter",
	Long: `Norminette - 42 Norm Linter
- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
A C source code linter builded with love by and for the 42 Community.
Checks every source file, header file and C Makefile of a specific directory
and displays every norm errors.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello world!")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Panic(err)
	}
}
