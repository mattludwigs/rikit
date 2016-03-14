package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"regexp"

	"github.com/codegangsta/cli"
)

type Config struct {
	Sites SitesConfig `json:"sites, omitempty"`
}

type SitesConfig map[string]ConfigObject

type ConfigObject struct {
	URL  string `json:"url,omitempty"`
	Auth string `json:"auth,omitempty"`
}

func main() {
	app := cli.NewApp()

	app.Name = "Rikit"
	app.Version = "1.0.0"
	app.Usage = "API testing CLI"

	var inConfig bool
	var url string
	var path string
	var authToken bool
	var conf Config

	app.Commands = []cli.Command{
		{
			Name:  "get",
			Usage: "Send get request",
			Action: func(c *cli.Context) {

				u, _ := user.Current()
				home := u.HomeDir

				match, _ := regexp.MatchString(`http:\/\/([a-zA-Z0-9]\w+.\w+)`, c.Args().First())

				if !match {
					file, e := ioutil.ReadFile(home + "/.rikit.json")

					if e != nil {
						log.Fatal(e)
					}

					json.Unmarshal(file, &conf)

					_, inConfig = conf.Sites[c.Args().First()]

					if inConfig {
						url = conf.Sites[c.Args().First()].URL
						url += path
					} else {
						url = c.Args().First()
					}
				}

				fmt.Printf("Using: %v\n", url)

				client := &http.Client{}
				req, err := http.NewRequest("GET", url, nil)
				req.Header.Set("User-Agent", "Rikit")

				if authToken {
					req.Header.Set("Authorization", conf.Sites[c.Args().First()].Auth)
				}

				res, err := client.Do(req)

				defer res.Body.Close()

				if err != nil {
					log.Fatal(err)
				}

				contents, err := ioutil.ReadAll(res.Body)

				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("%v\n", string(contents))

			},
			Flags: []cli.Flag{
				cli.StringFlag{Name: "path, p", Destination: &path},
				cli.BoolFlag{Name: "auth, a", Destination: &authToken},
			},
		},
	}

	app.Run(os.Args)
}
