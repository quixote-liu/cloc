package main

// type Command struct {
// 	path  string
// 	sorts map[string]string
// 	order map[string]string
// }

// func parseArguments(args []string) Command {
// 	if len(args) == 1 {
// 		log.Println("[ERROR]: missing arguments, see 'cloc help'")
// 		os.Exit(ExitCodeFailed)
// 	}
// 	if args[1] == "help" {
// 		fmt.Printf("%v\n", helptext)
// 		os.Exit(ExitCodeFailed)
// 	}
// 	if !strings.HasPrefix(args[1], "/") && !strings.HasPrefix(args[1], "./") {
// 		log.Println("[ERROR]: the first argument must be file path")
// 		os.Exit(ExitCodeFailed)
// 	}
// 	c := Command{path: args[1]}

// 	if len(args) == 2 {
// 		return c
// 	}

// 	m := make(map[string]string)
// 	isSubCommand := true
// 	for _, v := range args[2:] {
// 		if isSubCommand {
// 			if !strings.HasPrefix(v, "-") {
// 				log.Println("the subcommand argument must begin with '-'")
// 				os.Exit(ExitCodeFailed)
// 			}

// 		}
// 	}

// 	return c
// }
