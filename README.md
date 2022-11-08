
Do you sometimes forget where you stored your project ?


You search in several folders where you think it might be and in the end you find it but you have wasted a lot of time and energy with
`gproject` you can store the path and access your projects in one command line.

## How to use 

This CLI has 3 sub commands (add, go and ls)


The add command allows you to add a project by specifying the name of the project and its path, this will be saved in a json file:

```bash
$ gproject add <project_name> <project_path>
```

The ls command allows you to list the projects to store:

```bash 
$ gproject ls 
```

The go command will allow you to move directly into your project by specifying the project you want to access:

```bash
$ gproject go <project_name>
```


## How install

You can install the CLI with [go](https://go.dev/dl/) by typing the following command:

```bash
$ go install github.com/Dar-rius/gproject@latest
```


Then create an environment variable naming it `gproject`.


The stable version is not available, if you encounter a problem do not hesitate to create an [issue](https://github.com/Dar-rius/gproject/issues).
