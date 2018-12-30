package cmd

import (
	"fmt"
	"testing"
)

func TestGetV1Data(t *testing.T) {
	var j Jue
	d, _ := j.GetV1Data()
	fmt.Println(d)
}
