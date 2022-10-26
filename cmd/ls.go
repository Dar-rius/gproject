package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// lsCmd permet de lsiter tous les projets contenus dans l'appli
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		viper.SetConfigName("project")
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
