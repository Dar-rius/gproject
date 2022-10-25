package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// goCmd represents the go command
var goCmd = &cobra.Command{
	Use:   "go",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples`,
	Run: func(cmd *cobra.Command, args []string) {
		if args == nil && args[0] == "" {
			panic("Erreur sur la commande")
		}
		goPath(&args[0])
	},
}

func goPath(project *string) {
	viper.SetConfigName("project")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	path := viper.GetString(*project)

	os.Chdir(path)
	dir, err := os.Getwd()
	if err != nil {
		panic("ca existe pas")
	}
	fmt.Println(dir)
}

func init() {
	rootCmd.AddCommand(goCmd)

}
