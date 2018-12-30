package cmd

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func (j *Jue) GetV1Data() (string, error) {
	have_more := true
	s := ""
	t, _ := time.LoadLocation("Asia/Shanghai")
	u, _ := base64.StdEncoding.DecodeString(jueV1Url)
	p := string(u) + jueV1Path
	after := ""
	for have_more {
		have_more = false
		r := &JueReq{}
		r.Variables.Size = 2
		r.Variables.After = after
		r.Extensions.Query.Id = "964dab26a3f9997283d173b865509890"
		x, _ := json.Marshal(r)
		req, _ := http.NewRequest(http.MethodPost, p, bytes.NewBuffer(x))
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("X-Agent", "Juejin/Web")
		client := &http.Client{}
		resp, _ := client.Do(req)
		data, _ := ioutil.ReadAll(resp.Body)
		var JueV1Results JueV1Result
		if err := json.Unmarshal(data, &JueV1Results); err != nil {
			return s, err
		}
		for _, v := range JueV1Results.Data.RecommendedActivityFeed.Items.Edges {
			for _, vv := range v.Node.Targets {
				now, err := time.ParseInLocation("2006-01-02T15:04:05Z", vv.CreatedAt, t)
				if err != nil {
					panic(err)
				}
				if now.Format("2006-01-02") != time.Now().In(t).Format("2006-01-02") {
					continue
				}
				s += fmt.Sprintf("<h2>%s %s</h2><br />", vv.Content, vv.Url)
				for _, vvv := range vv.Pictures {
					s += fmt.Sprintf("<img src='%s' width='600' height='auto'/>", vvv)
				}
				after = JueV1Results.Data.RecommendedActivityFeed.Items.PageInfo.EndCursor
				have_more = true
			}
		}
	}
	return s, nil
}

type JueReq struct {
	OperationName string `json:"operationName"`
	Query         string `json:"query"`
	Variables     struct {
		Size  int    `json:"size"`
		After string `json:"after"`
	} `json:"variables"`
	Extensions struct {
		Query struct {
			Id string `json:"id"`
		} `json:"query"`
	} `json:"extensions"`
}

type JueV1Result struct {
	Data struct {
		RecommendedActivityFeed struct {
			Items struct {
				Edges []struct {
					Node struct {
						Id      string `json:"id"`
						Action  string `json:"action"`
						Targets []struct {
							Id        string   `json:"id"`
							Content   string   `json:"content"`
							CreatedAt string   `json:createdAt`
							Url       string   `json:"url"`
							Pictures  []string `json:"pictures"`
						} `json:"targets"`
					} `json:"node"`
				} `json:"edges"`
				PageInfo struct {
					EndCursor string `json:"endCursor"`
				} `json:"pageInfo"`
				PositionInfo struct {
				} `json:"positionInfo"`
			} `json:"items"`
			NewItemCount int `json:"newItemCount"`
		} `json:"recommendedActivityFeed"`
	} `json:"data"`
}
