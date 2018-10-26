package cmd

import (
	"net/http"
	"os"
)

type HomeInterface interface {
	GetData() string
}

type Home struct {
}

func (h *Home) GetData() (data string) {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		os.Exit(1)
	}
	return resp.Status
}

func NewHome() *Home {
	return &Home{}
}
