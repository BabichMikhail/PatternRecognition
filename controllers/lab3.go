package controllers

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type Lab3Controller struct {
	beego.Controller
}

func GetNormMatrix(mat [][]float64) (ans [][]float64, disp []float64) {
	disp = make([]float64, len(mat))
	ans = make([][]float64, len(mat))
	for i := 0; i < len(mat); i++ {
		ans[i] = make([]float64, len(mat[i]))
		disp[i] = mat[i][i]

	}
	for i := 0; i < len(mat); i++ {
		ans[i][i] = 1
		for j := i + 1; j < len(mat[i]); j++ {
			ans[i][j] = mat[i][j] / (math.Sqrt(disp[i]) * math.Sqrt(disp[j]))
			ans[j][i] = ans[i][j]
		}
	}
	return ans, disp
}

func (this *Lab3Controller) Main() {
	rand.Seed(time.Now().UTC().UnixNano())
	M, _ := this.GetInt("M", 3)
	arr := make([][]float64, M)
	N, _ := this.GetInt("N-points", 100)
	this.Data["NPoints"] = N
	leftX, _ := this.GetFloat("left_x", -10.0)
	rightX, _ := this.GetFloat("right_x", 10.0)
	bottomY, _ := this.GetFloat("bottom_y", -10.0)
	topY, _ := this.GetFloat("top_y", 10.0)
	this.Data["LeftX"] = leftX
	this.Data["RightX"] = rightX
	this.Data["BottomY"] = bottomY
	this.Data["TopY"] = topY
	for i := 0; i < M; i++ {
		arr[i] = make([]float64, M)
		for j := i; j < M; j++ {
			if i == j {
				arr[i][j], _ = this.GetFloat("cor"+strconv.Itoa(i)+strconv.Itoa(j), 1.0)
			} else {
				arr[i][j], _ = this.GetFloat("cor"+strconv.Itoa(i)+strconv.Itoa(j), 0.0)
			}
		}
	}
	for i := 0; i < M; i++ {
		for j := i + 1; j < M; j++ {
			arr[j][i] = arr[i][j]
		}
	}
	normMatrix, disp := GetNormMatrix(arr)
	med := make([]float64, M)
	for i := 0; i < M; i++ {
		med[i], _ = this.GetFloat("med"+strconv.Itoa(i), 0.0)
	}
	arg1 := strings.Join(func(meds []float64) []string {
		ans := []string{}
		for i := 0; i < len(meds); i++ {
			ans = append(ans, strconv.FormatFloat(meds[i], 'E', -1, 64))
		}
		return ans
	}(med), " ")
	arg2 := strings.Join(func(cor [][]float64) []string {
		ans := []string{}
		for i := 0; i < len(cor); i++ {
			for j := 0; j < len(cor[i]); j++ {
				ans = append(ans, strconv.FormatFloat(cor[i][j], 'E', -1, 64))
			}
		}
		return ans
	}(arr), " ")
	//fmt.Println(arg1)
	//fmt.Println(arg2)

	cmd := exec.Command("python", "lab3_main.py")
	genRand, _ := this.GetBool("rand", false)
	comp1, err1 := this.GetInt("comp1", 0)
	if err1 != nil || comp1 <= 0 || genRand || comp1 > M {
		comp1 = rand.Intn(M)
	} else {
		comp1--
	}
	comp2, err2 := this.GetInt("comp2", 0)
	//fmt.Print(genRand)
	if err2 != nil || comp2 <= 0 || genRand || comp2 > M {
		for ; comp2 == comp1; comp2 = rand.Intn(M) {
		}
	} else {
		comp2--
	}
	f, _ := os.Create("lab3_input.txt")
	n, _ := f.WriteString(strconv.Itoa(M) + " " + strconv.Itoa(N) + " " + strconv.Itoa(comp1) + " " + strconv.Itoa(comp2) + "\n")
	f.WriteString(strconv.FormatFloat(leftX, 'E', -1, 64) + " " + strconv.FormatFloat(rightX, 'E', -1, 64) + "\n")
	f.WriteString(strconv.FormatFloat(bottomY, 'E', -1, 64) + " " + strconv.FormatFloat(topY, 'E', -1, 64) + "\n")
	f.WriteString(arg1 + "\n")
	f.WriteString(arg2 + "\n")
	fmt.Println("size: ", n)
	err := cmd.Run()
	fmt.Println(err)
	cmd.Wait()
	this.Data["MArr"] = med
	this.Data["Med"] = med
	this.Data["M"] = M
	this.Data["Cor"] = arr
	this.Data["Comp1"] = comp1 + 1
	this.Data["Comp2"] = comp2 + 1
	this.Data["Rand"] = genRand
	this.Data["Disp"] = disp
	this.Data["NormMatrix"] = normMatrix
	this.TplName = "lab3.tpl"
}
