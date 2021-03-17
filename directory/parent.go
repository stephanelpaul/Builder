package directory

import (
	"fmt"
	"log"
	"os"
	"time"
	"strconv"

	"github.com/ilarocca/Builder/utils"
	"github.com/manifoldco/promptui"
)

func parentDir(path string) (bool, error) {
	//check if file path exists, returns err = nil if file exists
	_, err := os.Stat(path)

	if err == nil {
		fmt.Println("Path already exists")
	}

	// should return true if file doesn't exist
	if os.IsNotExist(err) {
		//bypass Prompt msg
		if bypassPrompt() {
			errDir := os.MkdirAll(path, 0755)
			//should return nil once directory is made, if not, throw err
			if errDir != nil {
				log.Fatal(err)
			}
		} else {
			//prompt user if they'd like dir to be created
			mk := yesNo()

			if mk == true {
				errDir := os.MkdirAll(path, 0755)
				//should return nil once directory is made, if not, throw err
				if errDir != nil {
					log.Fatal(err)
				}

			} else {
				fmt.Println("Please create a directory for the Builder")
				os.Exit(1)
				return true, err
			}

		}

	}

	//check workspace env exists, if not, create it
	val, present := os.LookupEnv("BUILDER_PARENT_DIR")
	if !present {
		os.Setenv("BUILDER_PARENT_DIR", path)
	} else {
		fmt.Println("BUILDER_PARENT_DIR", val)
	}
	return true, err
}

func yesNo() bool {
	prompt := promptui.Select{
		Label: "Create A Directory? [Yes/No]",
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return result == "Yes"
}

//MakeParentDir does...
func MakeParentDir() {
	args := os.Args[1:]

	//handles -n flag
	name := utils.GetName(args)

	//add Unix timestamp to dir name
	currentTime := time.Now().Unix()

	// local path for now
	path := "./" + name +"_"+strconv.FormatInt(currentTime, 10)

	parentDir(path)

	MakeHiddenDir(path)
	MakeLogDir(path)
	MakeWorkspaceDir(path)
}

func bypassPrompt() bool {
	args := os.Args[1:]

	yesFlag := false

	for _, val := range args {
		if val == "--yes" || val == "-y" {
			yesFlag = true
		}
	}

	return yesFlag
}
