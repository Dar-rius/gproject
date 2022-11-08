package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// go is a sub-command to move to the project of our choice
var goCmd = &cobra.Command{
	Use:   "go",
	Short: "This command will allow you to move to the repository of a project that you have chosen",
	Long:  `This command will allow you to move to the repository of a project that you have chosen`,
	Run: func(cmd *cobra.Command, args []string) {
		if args == nil && args[0] == "" || len(args) > 1 || len(args) < 1 {
			log.Fatal("Command error")
		} else {
			goPath(&args[0])
		}
	},
}

//A function to search for the path of a project in the path.json
// file and change directory by executing a shell command
func goPath(project *string) {
	//the environment variable is stored in a variable in order to create and find the path.json file in the directory where the app is located
	filEnv := os.Getenv("gproject")
	sys := runtime.GOOS
	viper.SetConfigName("path")
	viper.SetConfigType("json")
	viper.AddConfigPath(filEnv)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	pathProject := viper.GetString(*project)
	if pathProject == "" {
		log.Fatal("Error, this project is not saved")
	}

	os.Chdir(pathProject)
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	writeBash(&dir, &sys)

	if sys == "windows" {
		if err := exec.Command("cmd", "/C", "start", filEnv+`/script.sh`).Run(); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := exec.Command("bash", "~", filEnv+`script.sh`).Run(); err != nil {
			log.Fatal(err)
		}
	}

}

//The writeBash function allows to write scripts in the script.sh file to be executed
func writeBash(dir *string, sys *string) {
	if *sys == "windows" {
		command := "cd " + *dir + "\n cmd \n"
		pathEnv := os.Getenv("gproject")
		data := []byte(command)
		err := ioutil.WriteFile(pathEnv+`/script.sh`, data, 0666)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		command := "cd " + *dir + "\n bash \n"
		pathEnv := os.Getenv("gproject")
		data := []byte(command)
		err := ioutil.WriteFile(pathEnv+`script.sh`, data, 0666)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func init() {
	rootCmd.AddCommand(goCmd)
}
