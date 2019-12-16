package common

import (
	"flag"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/freecracy/news/cmd"
	"github.com/freecracy/news/etc"

	"github.com/fatih/color"
	config "github.com/go-ozzo/ozzo-config"
	"github.com/urfave/cli"
)

type mailServer struct {
	host     string
	port     int
	username string
	password string
	from     string
	fromName string
	to       string
}

var open = flag.Bool("o", false, "this is a test")

func parse() *mailServer {
	host := flag.String("h", "", "your email host")
	port := flag.Int("p", 0, "mail server port")
	username := flag.String("u", "", "your email username")
	password := flag.String("P", "", "your email password")
	from := flag.String("f", "", "send from")
	fromName := flag.String("n", "", "from name")
	to := flag.String("t", "", "send to")
	help := flag.Bool("help", false, "help")
	version := flag.Bool("v", false, "show version info")
	flag.Parse()
	if *version {
		fmt.Println("version:", etc.APP_VERSION)
		os.Exit(0)
	}
	if *help ||
		*host == "" ||
		*port == 0 ||
		*username == "" ||
		*password == "" ||
		*from == "" ||
		*to == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}
	return &mailServer{
		host:     *host,
		port:     *port,
		username: *username,
		password: *password,
		from:     *from,
		fromName: *fromName,
		to:       *to,
	}
}

func Exec() {

	m := parse()

	// w := &cmd.Weather{}
	// wd, _ := w.GetData()
	var wg sync.WaitGroup

	h := &cmd.Home{}
	j := &cmd.Jue{}

	r1 := make(chan string)
	r2 := make(chan string)
	defer func() {
		close(r1)
		close(r2)
	}()
	go func() {
		wg.Add(1)
		defer wg.Done()
		data, _ := h.GetOneData(*open)
		r1 <- data
	}()
	go func() {
		wg.Add(1)
		defer wg.Done()
		data, _ := j.GetV1Data()
		r2 <- data
	}()

	wg.Wait()
	var Hdata, Jdata string
	Hdata = <-r1
	Jdata = <-r2
	//content := wd + "<hr />" + Hdata + "<hr />" + Jdata
	content := Hdata + "<hr />" + Jdata
	//发送邮件
	auth := smtp.PlainAuth("cn", m.username, m.password, m.host)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{m.to}
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	local, _ := time.LoadLocation("Asia/Shanghai")
	subject := fmt.Sprintf("Subject:%s\n", time.Now().In(local).Format(time.RFC850))
	msg := []byte(subject + mime + content)
	if err := smtp.SendMail(m.host+":"+strconv.Itoa(m.port), auth, m.from, to, msg); err != nil {
		log.Println(err.Error())
	}
	os.Exit(0)
	s := NewCnMail()
	s.Setup(m)
	s.SendMail(content)
	os.Exit(0)

	// os.Exit(0)
	color.Red("this is a test")
	app := cli.NewApp()
	app.Name = etc.APP_NAME

	app.Usage = "fight the loneliness!"
	app.Action = func(c *cli.Context) error {
		host := c.GlobalString("H")
		fmt.Println(host)
		return nil
	}
	app.Version = etc.APP_VERSION
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "type,t",
			Value: "main",
			Usage: "work space of todo list",
		},
		cli.StringFlag{
			Name:  "H,host",
			Value: "email host",
			Usage: "your mail host",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:      "list",
			Aliases:   []string{"l"},
			Usage:     "a",
			UsageText: "b",
			Action: func(c *cli.Context) error {
				fmt.Println("hello", c.Args().First())
				return nil
			},
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
