package utils

import (
	"fmt"
	"os"
	"os/exec"
)

// CheckArgs is...
func CheckArgs() {
	//Repo
	repo := GetRepoURL()
	cArgs := os.Args[1:]
	//if flag present, but no url
	if repo == "" {
		BuilderLog.Fatal("No Repo Url Provided")
	}

	//check to see if repo exists
	//git ls-remote lists refs/heads & tags of a repo, if none exists, exit status thrown
	//returns the exit status in err
	_, err := exec.Command("git", "ls-remote", repo, "-q").Output()
	if err != nil {
		BuilderLog.Fatal("Provided repository does not exist")
	}

	//check if artifact path is passed in
	var artifactPath string
	for i, v := range cArgs {
		if v == "--output" || v == "-o" {
			if len(cArgs) <= i+1 {
				BuilderLog.Fatal("No Output Path Provided")

			} else {
				artifactPath = cArgs[i+1]
				val, present := os.LookupEnv("BUILDER_OUTPUT_PATH")
				if !present {
					os.Setenv("BUILDER_OUTPUT_PATH", artifactPath)
				} else {
					fmt.Println("BUILDER_OUTPUT_PATH", val)
					fmt.Println("Output Path already present")
					BuilderLog.Error("Output path already present")
				}
			}
		}
		if v == "--compress" || v == "-z" || v == "-C" {
			os.Setenv("ARTIFACT_ZIP_ENABLED", "true")
		}
		if v == "--hidden" || v == "-H" {
			os.Setenv("HIDDEN_DIR_ENABLED", "true")
		}
	}
}
