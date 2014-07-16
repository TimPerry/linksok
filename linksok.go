package main

import (
	"fmt"
  "os"
  "net/http"
	"io/ioutil"
  "github.com/codegangsta/cli"
	"strings"
)

var brokenLinks = make([]string, 0)

func checkLinks(body []byte) {
	
	fmt.Println("test")
	
}

func downloadPage(url string) {
	
	if strings.HasPrefix(url, "http://") != true && strings.HasPrefix(url, "https://") != true {
		url = "http://" + url
	}
	
	resp, err := http.Get(url)
	
	if err != nil {
		brokenLinks = append(brokenLinks, url)
	  return
	}
	
	if(strings.Contains(resp.Header["Content-Type"][0], "text/html")) {
	
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		
		if err == nil {
			checkLinks(body)
		}
		
	}
	
}

func main() {
	
  app := cli.NewApp()
	
  app.Name = "Linksok"
  app.Usage = "Check a website for broken links"
	
	app.Commands = [] cli.Command {
	  {
	    Name:      "check",
	    ShortName: "c",
	    Usage:     "Check the given website for broken links",
	    Action: func(c *cli.Context) {
				
				website := c.Args().First()
				
				if website == "" {
					println("Please provide a website to check for broken links")
					return;
				}
				
				println("Checking " + website + " for broken links...")
				downloadPage(website)
								
				if len(brokenLinks) > 0 {
				
					println("The following links are broken:")
				
					for _, link := range brokenLinks {
					
						println(link)
					
					}
				
				}
				
	    },
	  },
	}

  app.Run(os.Args)
	
}