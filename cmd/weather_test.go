package cmd

import (
	"testing"

	"github.com/freecracy/news/cmd"
)

func TestWeather_GetData(t *testing.T) {
	tests := []struct {
		name    string
		w       *cmd.Weather
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"test",
			&cmd.Weather{},
			"",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &cmd.Weather{}
			got, err := w.GetData()
			if (err != nil) != tt.wantErr {
				t.Errorf("Weather.GetData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Weather.GetData() = %v, want %v", got, tt.want)
			}
		})
	}
}
