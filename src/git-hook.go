package main

import (
	//"github.com/fsouza/go-dockerclient"

	"fmt"
	"log"
	"os/exec"
	"strings"
)

func exeCmd(cmd string) string {
	parts := strings.Fields(cmd)
	systemCmd := parts[0]
	args := parts[1:len(parts)]

	out, err := exec.Command(systemCmd, args...).Output()
	if err != nil {
		log.Fatal(err)
	}

	cleanOutput := string(out)

	return cleanOutput[0 : len(cleanOutput)-1]
	//return cleanOutput
}

func isIn(strList []string, item string) bool {
	for _, value := range strList {
		if value == item {
			return true
		}
	}
	return false
}

func addStringToArray(list []string) {
	for i := 0; i < len(list)+1; i++ {

	}
}

func main() {
	/*endpoint := "unix:///var/run/docker.sock"
	  dockerClient, _ := docker.NewClient(endpoint)
	  imgs, _ := dockerClient.ListImages(docker.ListImagesOptions{All: false})
	  for _, img := range imgs {
	    fmt.Println("ID: ", img.ID)
	    fmt.Println("RepoTags: ", img.RepoTags)
	    fmt.Println("Created: ", img.Created)
	    fmt.Println("Size: ", img.Size)
	    fmt.Println("VirtualSize: ", img.VirtualSize)
	    fmt.Println("ParentId: ", img.ParentID)
	  }*/
	gitRoot := exeCmd("git rev-parse --show-toplevel")
	gitActualBranch := exeCmd("git rev-parse --abbrev-ref HEAD")
	gitCommitIDOriginBranch := exeCmd(fmt.Sprintf("git merge-base %s master", gitActualBranch))
	gitDiff := exeCmd(fmt.Sprintf("git diff %s HEAD --name-only", gitCommitIDOriginBranch))
	gitDiffList := strings.Split(gitDiff, "\n")
	modifiedArtifact := make([]string, 0, len(gitDiffList))
	var artifact string

	for _, modif := range gitDiffList {
		artifact = strings.Split(modif, "/")[0]
		if !isIn(modifiedArtifact, artifact) {
			fmt.Println(len(modifiedArtifact))
			modifiedArtifact[len(modifiedArtifact)] = artifact
		}
	}
	fmt.Println(gitRoot)
	fmt.Println(gitActualBranch)
	fmt.Println(modifiedArtifact)
	//fmt.Println(gitDiff)
	//fmt.Println(gitDiff)
	//git diff `git merge-base {0} master` HEAD --name-only
}
