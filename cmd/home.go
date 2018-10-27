package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	HOME_FORMAT = `
<h1>
    <a href='%s'>%s</a>
</h1>
<h2>
    %s
</h2>
<br />
`
)

type HomeInterface interface {
	GetData() string
}

type Home struct {
}

type info struct {
	Success int `json:"Success`
	Result  []a
}

type a struct {
	//Newsid        int         `json:"newsid"`
	Title string `json:"title"`
	// V             string      `json:"v"`
	// Orderdate     string      `json:"orderdate"`
	// Postdate      string      `json:"postdate"`
	Description string `json:"description"`
	// Image         string      `json:"image"`
	// Slink         string      `json:"slink"`
	// Hitcount      int         `json:"hitcount"`
	// Commentcount  int         `json:"commentcount"`
	// Cid           int         `json:"cid"`
	// Url           string      `json:"url"`
	// Live          int         `json:"live"`
	// Lapinid       int         `json:"lapinid"`
	// Forbidcomment string      `json:"forbidcomment"`
	// Imagelist     interface{} `json:"imagelist"`
	// C             string      `json:"c"`
	// Client        string      `json:"client"`
	Isad bool `json:"isad"`
	// Sid           int         `json:"sid"`
	// PostDateStr   string      `json:"PostDateStr"`
	// HitCountStr   string      `json:"HitCountStr"`
	WapNewsUrl string `json:"WapNewsUrl"`
	// NewsTips      interface{} `json:"NewsTips"`
}

func (h *Home) GetData() (z string, err error) {
	d, _ := base64.StdEncoding.DecodeString(homeUrl)
	m := string(d)
	str := m + homePath
	t := strconv.FormatInt(time.Now().Unix(), 10) + "000"
	param := url.Values{}
	u, _ := url.Parse(str)
	param.Set("Tag", "")
	param.Set("ot", t)
	param.Set("page", "0")
	u.RawQuery = param.Encode()
	uPath := u.String()
	resp, err := http.Get(uPath)
	data, _ := ioutil.ReadAll(resp.Body)
	info := info{}
	err = json.Unmarshal(data, &info)
	if err != nil {
		panic(err)
	}
	z = ""
	for _, v := range info.Result {
		z += fmt.Sprintf(HOME_FORMAT, v.WapNewsUrl, v.Title, v.Description)
	}
	return z, err
	return m, nil
}

func NewHome() *Home {
	return &Home{}
}
