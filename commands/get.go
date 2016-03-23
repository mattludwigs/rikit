package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/user"
	"regexp"

	"github.com/mattludwigs/rikit/config"

	"github.com/codegangsta/cli"
)

var inConfig bool
var url string
var path string
var query string
var conf config.Config

func isHTTPCall(apiURL string) bool {
	match, err := regexp.MatchString(`http:\/\/([a-zA-Z0-9]\w+.\w+)`, apiURL)

	if err != nil {
		panic(err)
	}

	return match
}

func GET() cli.Command {

	return cli.Command{
		Name:  "get",
		Usage: "Send get request",
		Action: func(c *cli.Context) {

			apiURL := c.Args().First()
			home := getHomeDir()

			if !isHTTPCall(apiURL) {
				file, e := ioutil.ReadFile(home + "/.rikit.json")

				if e != nil {
					log.Fatal(e)
				}

				json.Unmarshal(file, &conf)

				_, inConfig = conf.Sites[apiURL]

				if inConfig {
					url = fmt.Sprintf("%s%s", conf.Sites[apiURL].URL, path)
				} else {
					url = fmt.Sprintf("%s", apiURL)
				}
			}

			if query != "" {
				url = fmt.Sprintf("%s?%s", url, query)
			}

			fmt.Printf("Using: %v\n", url)

			client := &http.Client{}
			req, err := http.NewRequest("GET", url, nil)
			req.Header.Set("User-Agent", "Rikit")

			if conf.Sites[apiURL].Auth != "" {
				req.Header.Set("Authorization", conf.Sites[apiURL].Auth)
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
			cli.StringFlag{Name: "query, q", Destination: &query},
		},
	}
}

func getHomeDir() string {
	u, _ := user.Current()
	return u.HomeDir
}
