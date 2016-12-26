package controllers

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type Lab8Controller struct {
	beego.Controller
}

type Point struct {
	X float64
	Y float64
}

type ColorPoint struct {
	Point
	Color string
}

type Cluster struct {
	N    int
	Disp []float64
	Med  []float64
	Cor  [][]float64
	P    float64
	//Points []Point
	//Center Point
}

func FloatToString(floatArr []float64) string {
	ans := []string{}
	for i := 0; i < len(floatArr); i++ {
		ans = append(ans, strconv.FormatFloat(floatArr[i], 'E', -1, 64))
	}
	return strings.Join(ans, " ")
}

func FloatMatToString(floatMat [][]float64) string {
	ans := []string{}
	for i := 0; i < len(floatMat); i++ {
		ans = append(ans, FloatToString(floatMat[i]))
	}
	return strings.Join(ans, " ")
}

type DataFormat struct {
	N    int
	Med  []float64
	Disp []float64
}

func ReadMyFormatFile(fname string, dim int) [][]DataFormat {
	f, err := os.Open(fname)
	fmt.Print(err)
	reader := io.Reader(f)
	data := [][]DataFormat{}
	for true {
		newData := []DataFormat{}
		num := -1
		fmt.Fscan(reader, &num)
		//fmt.Println(num)
		if num == -1 {
			break
		}

		for j := 0; j < num; j++ {
			subData := DataFormat{}
			fmt.Fscan(reader, &subData.N)
			for i := 0; i < dim; i++ {
				var f float64
				fmt.Fscan(reader, &f)
				subData.Med = append(subData.Med, f)
			}
			for i := 0; i < dim; i++ {
				var f float64
				fmt.Fscan(reader, &f)
				subData.Disp = append(subData.Disp, f)
			}
			fmt.Fscanln(reader)
			newData = append(newData, subData)
		}
		data = append(data, newData)
	}
	//fmt.Print(data)
	return data
}

func (this *Lab8Controller) Main() {
	rand.Seed(time.Now().UTC().UnixNano())
	D, _ := this.GetInt("D", 2)
	fmt.Print(D)
	N, _ := this.GetInt("N", 4)
	T, _ := this.GetInt64("T", 10)
	K, _ := this.GetInt("K", 4)
	NN, _ := this.GetInt("NN", 1000)
	fmt.Print(NN)
	Alpha, _ := this.GetFloat("Alpha", 10)
	clusters := make([]Cluster, N)
	f, _ := os.Create("lab8_input.txt")
	Koef := 100.0
	leftX, _ := this.GetFloat("left_x", -Koef)
	rightX, _ := this.GetFloat("right_x", Koef)
	bottomY, _ := this.GetFloat("bottom_y", -Koef)
	topY, _ := this.GetFloat("top_y", Koef)
	f.WriteString(fmt.Sprintf("%d %d %d %d %e\n", D, N, T, K, Alpha))
	f.WriteString(fmt.Sprintf("%e %e\n", leftX, rightX))
	f.WriteString(fmt.Sprintf("%e %e\n", bottomY, topY))
	Comp1, _ := this.GetInt("Comp1", 1)
	Comp2, _ := this.GetInt("Comp2", 2)
	f.WriteString(fmt.Sprintf("%d %d\n", Comp1-1, Comp2-1))
	this.Data["LeftX"] = leftX
	this.Data["RightX"] = rightX
	this.Data["BottomY"] = bottomY
	this.Data["TopY"] = topY

	for i := 0; i < N; i++ {
		for j := 0; j < D; j++ {
			cor := []float64{}
			for k := 0; k < D; k++ {
				var defValue float64
				if j == k {
					defValue = 1.0
				} else {
					defValue = 0.0
				}
				newCor, _ := this.GetFloat(fmt.Sprintf("cluster-сor-%d-%d-%d", i, j, k), defValue)
				cor = append(cor, newCor)
			}
			clusters[i].Cor = append(clusters[i].Cor, cor)
		}
		for j := 0; j < D; j++ {
			disp, _ := this.GetFloat(fmt.Sprintf("cluster-disp-%d-%d", i, j), 1.0)
			clusters[i].Disp = append(clusters[i].Disp, disp)
		}
		for j := 0; j < D; j++ {
			med, _ := this.GetFloat(fmt.Sprintf("cluster-med-%d-%d", i, j), rand.Float64()*Koef*2-Koef)
			clusters[i].Med = append(clusters[i].Med, med)
		}
		clusters[i].N, _ = this.GetInt(fmt.Sprintf("cluster-N-%d", i), 100)
		clusters[i].P, _ = this.GetFloat(fmt.Sprintf("cluster-P-%d", i), float64(1)/float64(N))
		f.WriteString(fmt.Sprintf("%d %e\n", clusters[i].N, clusters[i].P))
		f.WriteString(FloatToString(clusters[i].Disp) + "\n")
		f.WriteString(FloatToString(clusters[i].Med) + "\n")
		f.WriteString(FloatMatToString(clusters[i].Cor) + "\n")
	}

	fmt.Println(N, T)
	this.Data["D"] = D
	this.Data["N"] = N
	this.Data["T"] = T
	this.Data["K"] = K
	this.Data["Comp1"] = Comp1
	this.Data["Comp2"] = Comp2
	this.Data["Alpha"] = Alpha
	this.Data["Clusters"] = clusters
	cmd := exec.Command("python", "lab8_main.py")
	err := cmd.Run()
	fmt.Println(err)
	cmd.Wait()
	this.Data["Data0"] = ReadMyFormatFile("lab8_output_0.txt", D)
	this.Data["Data1"] = ReadMyFormatFile("lab8_output_1.txt", D)
	this.Data["Data2"] = ReadMyFormatFile("lab8_output_2.txt", D)
	this.Data["Data3"] = ReadMyFormatFile("lab8_output_3.txt", D)
	this.TplName = "lab8.tpl"
}
