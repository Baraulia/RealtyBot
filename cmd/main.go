package cmd

import (
	"github.com/Baraulia/RealtyBot/realtyParser"
	"time"
)

func main() {
	parser := realtyParser.NewParser()
	ticker := time.NewTicker(5 * time.Minute)
}
