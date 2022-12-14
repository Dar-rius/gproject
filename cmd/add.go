package cmd

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//add allows to add projects in the name and path of a project in the path.json file
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "This command allows you to add a new project and its path",
	Long: `This command allows you to add a new project and its path. `,
	Run: func(cmd *cobra.Command, args []string) {
		//init struct project
		var project Project

		if args == nil && args[0] == "" && args[1] == "" || len(args) > 2 || len(args) < 2 {
			log.Fatal("Command error")
		} else if args[0] != "" && args[1] == "." {
			project.name = args[0]
			addProjectActually(&project)
		} else {
			project.name = args[0]
			project.path = args[1]
			addProject(&project)
		}
	},
}

// A structure with the necessary fields to find a project
type Project struct {
	name, path string
}

//This function retrieves the current directory of the project
func addProjectActually(project *Project) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	project.path = dir
	addProject(project)
}

//The add Project function allows to search the json file and save the data concerning the project (name and path)
func addProject(project *Project) {
	//the environment variable is stored in a variable in order to create and find the path.json file in the directory where the app is located
	filEnv := os.Getenv("gproject")
	sys := runtime.GOOS
	if sys == "linux" || sys == "darwin"{
		_, errs := os.OpenFile(filEnv+"path.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if errs != nil{
			log.Fatal(errs)
		}
	} else {
		_, errs := os.OpenFile(filEnv+"\\path.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if errs != nil{
			log.Fatal(errs)
		}

	}
	
	viper.SetConfigName("path")
	viper.SetConfigType("json")
	viper.AddConfigPath(filEnv)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	viper.Set(project.name, project.path)
	viper.WriteConfig()

	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("project add")
	})
	viper.WatchConfig()
}

func init() {
	rootCmd.AddCommand(addCmd)
}
