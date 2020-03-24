package cmd

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func (j *Jue) GetV1Data() (string, error) {
	haveMore := true
	s := ""
	t, _ := time.LoadLocation("Asia/Shanghai")
	u, _ := base64.StdEncoding.DecodeString(jueV1Url)
	p := string(u) + jueV1Path
	after := ""
	for haveMore {
		time.Sleep(2 * time.Microsecond)
		haveMore = false
		r := &JueReq{}
		r.Variables.Size = 2
		r.Variables.After = after
		r.Extensions.Query.Id = "249431a8e4d85e459f6c29eb808e76d0"
		x, _ := json.Marshal(r)
		req, _ := http.NewRequest(http.MethodPost, p, bytes.NewBuffer(x))
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("X-Agent", "Juejin/Web")
		req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.79 Safari/537.36")
		client := &http.Client{}
		resp, _ := client.Do(req)
		data, _ := ioutil.ReadAll(resp.Body)
		log.Println(string(data))
		var JueV1Results JueV1Result
		if err := json.Unmarshal(data, &JueV1Results); err != nil {
			return s, err
		}
		for _, v := range JueV1Results.Data.RecommendedActivityFeed.Items.Edges {
			for _, vv := range v.Node.Targets {
				now, err := time.ParseInLocation("2006-01-02T15:04:05Z", vv.CreatedAt, t)
				if err != nil {
					continue
				}
				if now.Format("2006-01-02") != time.Now().In(t).Format("2006-01-02") {
					continue
				}
				for _, vvv := range vv.Pictures {
					s += fmt.Sprintf("<a href='%s'><img src='%s' width='600' height='auto'/></a>", vvv, vvv)
				}
				s += fmt.Sprintf("<h2>%s %s</h2><br />", vv.Content, vv.Url)
				after = JueV1Results.Data.RecommendedActivityFeed.Items.PageInfo.EndCursor
				haveMore = true
				time.Sleep(time.Millisecond)
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
