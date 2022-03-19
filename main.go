package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"time"
)

func main() {
	cmd := exec.Command("gh", "repo", "list", "--limit=200", "--json=owner,name")
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	fmt.Printf("running command: %s\n", cmd)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	type Repos struct {
		Name  string `json:"name"`
		Owner struct {
			ID    string `json:"id"`
			Login string `json:"login"`
		} `json:"owner"`
	}
	var repo []Repos

	if err := json.NewDecoder(&outb).Decode(&repo); err != nil {
		log.Fatal(err)
	}
	for _, r := range repo {
		repoName := fmt.Sprintf("%s/%s", r.Owner.Login, r.Name)
		prCommand := exec.Command("gh", "pr", "list", "--repo", repoName, "--json=number,author")
		var outbs, errbs bytes.Buffer
		prCommand.Stdout = &outbs
		prCommand.Stderr = &errbs
		fmt.Printf("running command: %s\n", prCommand)
		fmt.Println("PRs: ", outbs.String())
		if err := prCommand.Run(); err != nil {
			log.Fatal(`err`, err, errbs.String())
		}

		type PR struct {
			Author struct {
				Login string `json:"login"`
			} `json:"author"`
			Number int `json:"number"`
		}
		var pr []PR
		if err := json.NewDecoder(&outbs).Decode(&pr); err != nil {
			log.Fatal(err)
		}
		if len(pr) != 0 {
			for _, p := range pr {
				if p.Author.Login == "dependabot" {
					prCommand := exec.Command("gh", "pr", "merge", fmt.Sprintf("%d", p.Number), "-m", "--repo", repoName)
					var outbss, errbss bytes.Buffer
					prCommand.Stdout = &outbss
					prCommand.Stderr = &errbss
					fmt.Printf("running command: %s\n", prCommand)
					if err := prCommand.Run(); err != nil {
						log.Fatal(`err`, err, errbss.String())
					}
					fmt.Println("Merged: ", outbss.String())
				}
			}
		}
		fmt.Println("pr", pr)
		time.Sleep(1 * time.Second)
	}
}
