package cmd

import (
	"encoding/base64"
	//"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type HomeInterface interface {
	GetData() string
}

type Home struct {
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
	z = string(data)
	return z, err
	return m, nil
}

func NewHome() *Home {
	return &Home{}
}
