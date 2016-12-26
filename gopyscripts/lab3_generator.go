package gopyscripts

import "os"

func PrintLab3() {
	f, _ := os.Create("lab3_main.py")
	f.WriteString(`
import numpy as np
import matplotlib.pyplot as plt
from math import *
import random
from tkinter import *
from copy import *

import numpy as np
from matplotlib.mlab import bivariate_normal
from matplotlib import pylab as plt

lines = []
fin = open('lab3_input.txt', 'r')
for line in fin:
    lines.append(line)
inputData = lines[0].split()
M = int(inputData[0])
N = int(inputData[1])
First = int(inputData[2])
Second = int(inputData[3])
inputData = lines[1].split()
leftX = float(inputData[0])
rightX = float(inputData[1])
inputData = lines[2].split()
bottomY = float(inputData[0])
topY = float(inputData[1])
inputData = lines[3].split()
meds = []
for i in range(M):
    meds.append(float(inputData[i]))
matrix = []
inputData = lines[4].split()
for i in range(M):
    newRow = []
    for j in range(M):
        newRow.append(float(inputData[i*M + j]))
    matrix.append(newRow)
offest = M + M*M + 3

x = np.arange(leftX, rightX, 0.1)
y = np.arange(bottomY, topY, 0.1)
X, Y = np.meshgrid(x, y)
sigmaX = sqrt(matrix[First][First])
sigmaY = sqrt(matrix[Second][Second])
Z = bivariate_normal(X, Y, sigmaX, sigmaY, meds[First], meds[Second], matrix[First][Second])
plt.figure(1)
plt.xlim([leftX - 1, rightX + 1])
plt.ylim([bottomY - 1, topY + 1])
plt.contour(X, Y, Z)
plt.savefig("static/img/img1.png")
plt.figure(2)
plt.xlim([leftX - 1, rightX + 1])
plt.ylim([bottomY - 1, topY + 1])
plt.contour(X, Y, Z)
x, y = np.random.multivariate_normal([ meds[First], meds[Second] ], [ [matrix[First][First], matrix[First][Second]], [ matrix[Second][First], matrix[Second][Second] ] ], N).T
plt.plot(x, y, 'b.')
plt.savefig("static/img/img2.png")
plt.figure(3)
plt.xlim([leftX - 1, rightX + 1])
plt.ylim([bottomY - 1, topY + 1])
plt.plot(x, y, 'b.')
plt.savefig("static/img/img3.png")
#plt.show()
    `)
}
