package main

import (
  "os"
  "github.com/codegangsta/cli"
  "time"
  "os/exec"
)

func main() {
  app := cli.NewApp()
  app.Name = "jarvis"
  app.Usage = "I am a nice butler, who can speak and listen"
  app.Action = func(c *cli.Context) {
    println("Hello friend, how can I help you!")
  }
  app.Commands = []cli.Command{
    {
      Name:      "ask",
      ShortName: "a",
      Usage:     "Ask me anything",
      Action: ask,
    },
    {
      Name:      "clock",
      ShortName: "c",
      Usage:     "tells you the current time",
      Action: clock,
    },
    {
      Name:      "date",
      ShortName: "d",
      Usage:     "tells you the current date",
      Action: date,
    },     
    {
      Name:      "introduce",
      ShortName: "i",
      Usage:     "Introduces yourself",
      Action: introduce,
    },     
    {
      Name:      "timer",
      ShortName: "t",
      Usage:     "set the timer to X seconds, minutes or hours",
      Action: timer,
    },
    {
      Name:      "timer_finish",
      ShortName: "tf",
      Usage:     "informs everyone that the timer has finished",
      Action: timer_finish,
    },	
    {
      Name:      "weather",
      ShortName: "w",
      Usage:     "Checks todays weather, e.g. Hamburg, Germany",
      Action: weather,
    },	
  }
  app.Run(os.Args)
}


func ask(c *cli.Context) {
	
}

func clock(c *cli.Context) {
  	const layout = "January 2, 2006"
  	t := time.Now()
	say("Today is " + t.UTC().Format(layout))
}

func date(c *cli.Context) {
  	const layout = "03:04PM"
  	t := time.Now()
	say("It is " + t.Format(layout))
}

func introduce(c *cli.Context) {
	
}

func timer(c *cli.Context) {
	
}

func timer_finish(c *cli.Context) {
	
}

func weather(c *cli.Context) {
	
}



func say(text string) {
	exec.Command("say", text).Run()
}