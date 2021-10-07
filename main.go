package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Project struct {
	AddedToGithub    string
	ZilliqaToAdd     string
	Name             string
	ShortDescription string
	Logo             string
	Messaged         string
	CategoryTag      string
	Status           string
	TwitterURL       string
	WebsiteURL       string
	DiscordURL       string
	TelegramURL      string
}

func main() {
	file, err := os.Open("projects.tsv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "\t")
		name := fields[2]
		description := fields[3]
		logo := fields[4]
		categories := fields[6]
		status := fields[7]
		twitter := fields[8]
		website := fields[9]
		discrod := fields[10]
		telegram := fields[11]
		err := os.Mkdir(name, 0755)
		if err != nil {
			log.Fatal(err.Error())
		}
		fileName := fmt.Sprintf("%s/project.md", name)
		text := fmt.Sprintf("%s\nname: %s\ndescription: %s\nlogo: %s\ncategories: %s\nstatus: %s\nttwitter: %s\nwebsite: %s\ndiscord: %s\ntelegram: %s\n%s", "---", name, description, logo, categories, status, twitter, website, discrod, telegram, "---")
		appendFile(fileName, text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func appendFile(filename string, text string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		log.Fatal(err.Error())
	}
}
