package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ls allows you to list all the projects saved in the path.json file
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "This command will list all the projects saved in the path.json file",
	Long: `This command will list all the project saved in the path.json file
			example: gproject ls`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.SetConfigName("path")
		viper.AutomaticEnv()
		viper.SetConfigType("json")
		viper.AddConfigPath(".")
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal(err)
		}
		for _, v := range viper.AllKeys() {
			fmt.Println(v)
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

}
