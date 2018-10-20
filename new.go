package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/go-ozzo/ozzo-config"
	"github.com/urfave/cli"
)

func main() {

	color.Red("this is a test")
	app := cli.NewApp()
	app.Name = "new"
	app.Usage = "fight the loneliness!"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello friend!")
		return nil
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "type,t",
			Value: "main",
			Usage: "work space of todo list",
		},
		cli.StringFlag{
			Name:  "a,b",
			Value: "main",
			Usage: "work space of todo list",
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(0)

	var hello = `woain
woaini
`

	fmt.Println(hello)
	fmt.Println(hello)

	// 创建一个新的 Config 对象
	c := config.New()

	c.Load("config/app.json")

	// 尝试在配置中读取 "Version" 的值，若找不到，则返回默认值 "1.0"
	version := c.GetString("Version", "1.0")
	a := c.GetString("a", "b")

	var author struct {
		Name, Email string
	}
	// 用 "Author" 部分的配置填充 author 对象的属性。
	c.Configure(&author, "Author")

	fmt.Println(version)
	fmt.Println(author.Name)
	fmt.Println(author.Email)
	fmt.Println(a)
}
