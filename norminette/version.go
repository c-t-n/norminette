package norminette

import (
	"fmt"
	
	"github.com/spf13/cobra"
)

const Version = "0.0.1"

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "displays the current version of Norminette",
	Long: "https://semver.org | Because it's better",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version ", Version)
	},
}