package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// goCmd represents the go command
var goCmd = &cobra.Command{
	Use:   "go",
	Short: "",
	Long:  ``,
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

	//write file
	writeBash(&dir)

	if err := exec.Command("cmd", "/C", "start", "C:/Users/MOHAM/Desktop/project/goproject/script.sh").Run(); err != nil {
		log.Fatal(err)
	}
}

func writeBash(dir *string) {
	commande := "cd " + *dir + "\n bash \n"
	data := []byte(commande)
	err := ioutil.WriteFile("C:/Users/MOHAM/Desktop/project/goproject/script.sh", data, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(goCmd)
}
