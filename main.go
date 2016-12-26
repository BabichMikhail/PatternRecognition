package main

import (
	"math/rand"
	"time"

	"github.com/BabichMikhail/PatternRecognition/gopyscripts"
	_ "github.com/BabichMikhail/PatternRecognition/routers"
	"github.com/astaxie/beego"
)

func sum(a float64, b float64) float64 {
	return a + b
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	gopyscripts.PrintLab3()
	gopyscripts.PrintLab8()
	beego.AddFuncMap("sum", sum)
	beego.Run()
}
