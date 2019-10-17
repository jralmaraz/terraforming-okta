package main

import (
  "log"
  "os"
  "sort"

  "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()
  app.Name = "terraforming-okta"
  app.Flags = []cli.Flag {
    cli.StringFlag{
      Name: "generate, g",
      Value: "generate",
      Usage: "Default option - connects to an Okta org and maps them to terraform resources",
    },
    cli.StringFlag{
      Name: "import, i",
      Usage: "Import terraform resource from `FILE`",
    },
  }

  app.Commands = []cli.Command{
    {
      Name:    "generate",
      Aliases: []string{"g"},
      Usage:   "generatte",
      Action:  func(c *cli.Context) error {
        return nil
      },
    },
    {
      Name:    "import",
      Aliases: []string{"i"},
      Usage:   "import",
      Action:  func(c *cli.Context) error {
        return nil
      },
    },
  }

  sort.Sort(cli.FlagsByName(app.Flags))
  sort.Sort(cli.CommandsByName(app.Commands))

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}

