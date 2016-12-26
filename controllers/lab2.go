package controllers

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func F_rasp(x float64) float64 {
	if x <= 0 {
		return 0
	}
	if x >= 1 {
		return 1
	}
	return 2.0 / math.Pi * math.Asin(x)
}

func F_plot(x float64) float64 {
	if x <= 0 || x >= 1 {
		return 0
	}
	return 2.0 / math.Pi / math.Sqrt(1-x*x)
}

func UpdateN(this *MainController) int {
	newN := this.GetString("N")
	oldN := this.GetSession("N")
	if oldN == nil || oldN.(int) == 0 {
		if i, err := strconv.Atoi(newN); len(newN) > 0 && err == nil && i > 0 {
			this.SetSession("N", i)
		} else {
			this.SetSession("N", 10000)
		}
		newN = strconv.Itoa(this.GetSession("N").(int))
	}
	i, _ := strconv.Atoi(newN)
	this.Data["N"] = newN
	return i
}

func UpdateIntervals(this *MainController) int {
	newN := this.GetString("Gist")
	oldN := this.GetSession("Gist")
	if oldN == nil || oldN.(int) == 0 {
		if i, err := strconv.Atoi(newN); len(newN) > 0 && err == nil && i > 0 {
			this.SetSession("Gist", i)
		} else {
			this.SetSession("Gist", 10)
		}
		newN = strconv.Itoa(this.GetSession("Gist").(int))
	}
	i, _ := strconv.Atoi(newN)
	this.Data["GistIntervals"] = newN
	return i
}

func myrand(a int, b int) float64 {
	return float64(a) + (float64(b)-float64(a))*math.Sin(math.Pi/2*rand.Float64())
}

var L float64

func appendInterval(vector *[]int, v float64, a int, b int, Int int) {
	L = (float64(b) - float64(a)) / float64(Int)
	num := int((v - float64(a)) / L)
	(*vector)[num]++
}

func med(v []float64) float64 {
	sum := float64(0.0)
	for i := 0; i < len(v); i++ {
		sum += v[i]
	}
	return sum / float64(len(v))
}

func cmoment(v []float64, med float64, I float64) float64 {
	sum := float64(0.0)
	for i := 0; i < len(v); i++ {
		sum += math.Pow(v[i]-med, I)
	}
	return sum / float64(len(v))
}

func disp(v []float64, med float64) float64 {
	return cmoment(v, med, 2)
}

func cmoment3(v []float64, med float64) float64 {
	return cmoment(v, med, 3)
}

func cmoment4(v []float64, med float64) float64 {
	return cmoment(v, med, 4)
}

func assim(v []float64, med float64) float64 {
	return cmoment3(v, med) / math.Pow(disp(v, med), 3.0/2.0)
}

func ekscess(v []float64, med float64) float64 {
	return cmoment4(v, med)/math.Pow(disp(v, med), 4.0/2.0) - 3
}

func (this *MainController) Main() {
	N := UpdateN(this)
	Int := UpdateIntervals(this)
	v := []float64{}
	a, b := 0, 1
	vecInt := make([]int, Int)
	for i := 0; i < N; i++ {
		val := (myrand(a, b))
		v = append(v, val)
		appendInterval(&vecInt, val, a, b, Int)
	}
	srednee := med(v)
	this.Data["A"] = a
	this.Data["B"] = b
	this.Data["Gist"] = vecInt
	this.Data["N"] = N
	this.Data["L"] = L
	this.Data["Intervals"] = Int
	this.Data["Srednee"] = fmt.Sprintf("%.8f", srednee)
	this.Data["SredneeAn"] = 0.63662977
	this.Data["SredneeDif"] = fmt.Sprintf("%.8f", srednee-0.63662977)
	d := disp(v, srednee)
	this.Data["Disp"] = fmt.Sprintf("%.8f", d)
	this.Data["DispAn"] = 0.094715
	this.Data["DispDif"] = fmt.Sprintf("%.8f", d-0.094715)
	e := ekscess(v, srednee)
	this.Data["Ekscess"] = fmt.Sprintf("%.8f", e)
	this.Data["EkscessAn"] = -1.06842795
	this.Data["EkscessDif"] = fmt.Sprintf("%.8f", e-(-1.06842795))
	as := assim(v, srednee)
	this.Data["Assim"] = fmt.Sprintf("%.8f", as)
	this.Data["AssimAn"] = -0.49716835
	this.Data["AssimDif"] = fmt.Sprintf("%.8f", as-(-0.49716835))
	this.TplName = "lab2.tpl"
}
