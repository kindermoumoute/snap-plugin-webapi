package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/codegangsta/cli"
)

type Plugin struct {
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Type        string `json:"type"`
	Owner       string `json:"owner"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Forks       int    `json:"fork_count"`
	Stars       int    `json:"star_count"`
	Watchers    int    `json:"watch_count"`
	Issues      int    `json:"issues_count"`
}

func printPlugin(p Plugin) {
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("FullName: %s\n", p.FullName)
	fmt.Printf("Type: %s\n", p.Type)
	fmt.Printf("Owner: %s\n", p.Owner)
	fmt.Printf("Description: %s\n", p.Description)
	fmt.Printf("URL: %s\n", p.URL)
	fmt.Printf("Forks: %d\n", p.Forks)
	fmt.Printf("Stars: %d\n", p.Stars)
	fmt.Printf("Watchers: %d\n", p.Watchers)
	fmt.Printf("Issues: %d\n\n", p.Issues)
}

func printType(plugins []Plugin, title string, pluginType string) {
	fmt.Printf("\n========%s========\n", title)
	for _, v := range plugins {
		if v.Type == pluginType {
			fmt.Println(v.Name)
		}
	}

}

func infoByName(ctx *cli.Context) {
	if len(ctx.Args()) > 1 {
		fmt.Println("Incorrect usage of list command--please enter in format: plugin list [plugin_name]")
	} else {
		resp, err := http.Get("http://localhost:8080/plugins/")
		if err != nil {
			fmt.Printf("error is: %s\n", err)
		} else {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
			}
			defer resp.Body.Close()
			plugins := make([]Plugin, 0)
			err = json.Unmarshal(body, &plugins)
			if err != nil {
				fmt.Println(err)
			}
			// List out all plugins
			if len(ctx.Args()) == 0 {
				printType(plugins, "Collectors", "collector")
				printType(plugins, "Processors", "processor")
				printType(plugins, "Publishers", "publisher")
			} else { // List out a specific plugin
				name := ctx.Args()[0]
				found := false
				for _, v := range plugins {
					if v.Name == name {
						found = true
						printPlugin(v)
					}
				}
				if found == false {
					fmt.Printf("There is no plugin with the name %s\n", name)
				}
			}
		}
	}
}
