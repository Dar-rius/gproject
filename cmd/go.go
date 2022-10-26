package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// go is a command to move to the project of our choice
var goCmd = &cobra.Command{
	Use:   "go",
	Short: "This command will allow you to move to the repository of a project that you have chosen",
	Long: `This command will allow you to move to the repository of a project that you have chosen 
			example: gproject go projectA`,
	Run: func(cmd *cobra.Command, args []string) {
		if args == nil && args[0] == "" || len(args) > 1 {
			log.Fatal("Command error")
		} else {
			goPath(&args[0])
		}
	},
}

//A function to search for the path of a project in the path.json
// file and change directory by executing a shell command
func goPath(project *string) {
	viper.SetConfigName("path")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	path := viper.GetString(*project)
	if path == "" {
		log.Fatal("Error, this project is not saved")
	}

	os.Chdir(path)
	dir, err := os.Getwd()
	if err != nil {
		panic("ca existe pas")
	}

	//We pass the value of the path in the writeBash function
	writeBash(&dir)

	//We execute the script contained in the file: script.sh
	if err := exec.Command("cmd", "/C", "start", "C:/Users/MOHAM/Desktop/project/goproject/script.sh").Run(); err != nil {
		log.Fatal(err)
	}
}

//The writeBash function allows you to write scripts in the script.sh file to be executed
func writeBash(dir *string) {
	//the command
	command := "cd " + *dir + "\n bash \n"
	data := []byte(command)
	err := ioutil.WriteFile("C:/Users/MOHAM/Desktop/project/goproject/script.sh", data, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(goCmd)
}
