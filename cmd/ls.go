package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// ls allows you to list all the projects saved in the path.json file
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "This command will list all the projects saved in the path.json file",
	Long: `This command will list all the project saved in the path.json file
			example: gproject ls`,
	Run: func(cmd *cobra.Command, args []string) {
		//the environment variable is stored in a variable in order to create and find the path.json file in the directory where the app is located
		filEnv := os.Getenv("goproject")
		file, err := os.ReadFile(filEnv + "/path.json")
		if err != nil {
			panic(err)
		}
		var data map[string]interface{}
		err = json.Unmarshal(file, &data)
		if err != nil {
			panic(err)
		}
		for k, _ := range data {
			fmt.Println(k)
		}

	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

}
