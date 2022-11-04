package realtyParser

import "time"

var sources = []Source{
	{
		url:      "testURL",
		tags:     []string{"a"},
		lastTime: time.Time{},
	},
}

type Source struct {
	url      string
	tags     []string
	lastTime time.Time
}

type Result struct {
	link     string
	date     time.Time
	title    string
	summary  string
	price    float32
	currency string
}
