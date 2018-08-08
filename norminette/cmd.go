package norminette

import (
	"log"

	"github.com/spf13/cobra"
)

type Config struct {
	Verbose bool
	Debug   bool
	Path    string
	Options Options
}

type Options struct {
	BypassPrototypeInCFile bool
	BypassLineWidth        bool
	BypassFunctionHeight   bool
}

var Cfg Config

var RootCmd = &cobra.Command{
	Use:   "norminette",
	Short: "Norminette - 42 Norm Linter",
	Long: `Norminette - 42 Norm Linter
- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
A C source code linter builded with love by and for the 42 Community.
Checks every source file, header file and C Makefile of a specific directory
and displays every norm errors.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.PersistentFlags().BoolVarP(&Cfg.Debug, "debug", "d", false, "display debug informations")
	RootCmd.PersistentFlags().BoolVarP(&Cfg.Verbose, "verbose", "v", false, "set the verbosity level a little higher")
	RootCmd.Flags().StringVarP(&Cfg.Path, "path", "p", "$PWD", "runs norminette on this path (default: $PWD)")
	RootCmd.Flags().BoolVarP(&Cfg.Options.BypassPrototypeInCFile, "bypass-prototypes", "", false, "skip errors related to prototypes functions in C files")
	RootCmd.Flags().BoolVarP(&Cfg.Options.BypassLineWidth, "bypass-line-width", "", false, "skip errors related to line width")
	RootCmd.Flags().BoolVarP(&Cfg.Options.BypassFunctionHeight, "bypass-function-height", "", false, "skip errors related to the number of lines inside a function block")
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Panic(err)
	}
}
