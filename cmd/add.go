package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//add allows to add projects in the name and path of a project in the path.json file
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "This command allows you to add a new project and its path",
	Long: `This command allows you to add a new project and its path
			example: gproject add name_project path_project
			or to add a project with the current path: gproject add .`,
	Run: func(cmd *cobra.Command, args []string) {
		//init struct projecto
		var project Project

		if args == nil && args[0] == "" && args[1] == "" || len(args) > 2 {
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

// a structure with the necessary fields to find a project
type Project struct {
	name, path string
}

//a function retrieves the current directory of the project
func addProjectActually(project *Project) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	project.path = dir
	addProject(project)
}

//the add Project function allows you to search the json file and save the data concerning the project (name and path)
func addProject(project *Project) {
	vp := viper.New()
	vp.SetConfigName("path")
	vp.SetConfigType("json")
	vp.AddConfigPath(`C:\Users\MOHAM\GoProject\`)
	err := vp.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	vp.Set(project.name, project.path)
	vp.WriteConfig()

	vp.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("le projet %s ete ajouter", in.Name)
	})

	vp.WatchConfig()
}

func init() {
	rootCmd.AddCommand(addCmd)
}
