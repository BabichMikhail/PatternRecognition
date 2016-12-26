package gopyscripts

import "os"

func PrintLab8() {
	f, _ := os.Create("lab8_main.py")
	f.WriteString(`
import numpy as np
import matplotlib.pyplot as plt
from math import *
import random
from tkinter import *
from copy import *

from matplotlib.mlab import bivariate_normal
from matplotlib import pylab as plt

lines = []
fin = open('lab8_input.txt', 'r')
for line in fin:
    lines.append(line)
inputData = lines[0].split()
D = int(inputData[0])
N = int(inputData[1])
T = int(inputData[2])
K = int(inputData[3])
Alpha = float(inputData[4])
marker = 5

inputData = lines[1].split()
leftX = float(inputData[0])
rightX = float(inputData[1])
inputData = lines[2].split()
bottomY = float(inputData[0])
topY = float(inputData[1])

inputData = lines[3].split()
componentA = int(inputData[0])
componentB = int(inputData[1])

clusters = []
for i in range(int((len(lines) - 4) / N)):
    j = 4*i + 4
    NP = lines[j].split()
    disp = lines[j + 1].split()
    med = lines[j + 2].split()
    cor = lines[j + 3].split()
    clusters.append([ int(NP[0]), float(NP[1]), disp, med, cor ])

x = np.arange(leftX, rightX, (rightX - leftX) / 5000)
y = np.arange(bottomY, topY, (topY - bottomY) / 5000)
X, Y = np.meshgrid(x, y)

def distance(point1, point2):
    dist = 0
    for i in range(len(point1)):
        if i == componentA or i == componentB:
            dist += ((point1[i] - point2[i])**2)
    return dist

def cmp_len(point, clusters, T):
    length = T
    idx = -1
    for i in range(len(clusters)):
        c = clusters[i]
        newLen = distance(c.center, point)
        if newLen < T and newLen < length:
            length = newLen
            idx = i
    return idx

def cmp(point1, point2):
    flag = True
    for i in range(len(point1)):
        flag &= point1[i] == point2[i]
    return flag

def getStrPoint(point):
    pointStr = ""
    for x in point:
        pointStr += " " + str(x)
    return pointStr

class Cluster:
    def __init__(self, point):
        self.init = False
        self.points = [ [ point[i] ] for i in range(len(point)) ]
        #print(len(self.points))
        self.center = copy(point)
        #print(self.center)
        self.medium = None
        self.dim = len(point)
        self.count = 1
        self.color = [ random.randint(1, 256) / 256, random.randint(1, 256) / 256, random.randint(1, 256) / 256 ]

    def appendPoint(self, point):
        if not self.init and cmp(point, self.center):
            self.init = True
            return
        for i in range(self.dim):
            self.points[i].append(point[i])
        self.count += 1

    def redefineCenter(self):
        medium = self.getMedium()
        self.center = medium

    def x(self):
        return self.points[componentA]

    def y(self):
        return self.points[componentB]

    def getDisp(self):
        med = self.getMedium()
        disp = []
        for j in range(self.dim):
            compDisp = 0
            for i in range(self.count):
                compDisp += (self.points[j][i] - med[j])**2
            disp.append(compDisp / self.count)
        return disp

    def getMedium(self):
        #if self.medium != None:
        #    return self.medium
        med = [ 0 for x in range(self.dim) ]
        for i in range(self.count):
            for j in range(self.dim):
                med[j] += self.points[j][i]
        self.medium = [ med[i] / self.count for i in range(self.dim) ]
        idx = -1
        minDist = 0
        for i in range(self.count):
            newDist = distance(self.medium, [ self.points[j][i] for j in range(self.dim) ])
            if idx == -1 or newDist < minDist:
                minDist = newDist
                idx = i
        self.medium = [ self.points[j][idx] for j in range(self.dim) ]
        return self.medium

    def flush(self):
        self.count = 1
        self.init = False
        self.points = [ [ self.center[i] ] for i in range(self.dim) ]
        self.medium = None

plt.figure(1)
points = []
sourceClusters = []
for i in range(len(clusters)):
    #clusters.append([ int(NP[0]), float(NP[1], disp, med, cor ])
    corMat = [ [ 0 for x in range(D) ] for y in range(D) ]
    med = [ 0 for x in range(D) ]
    for j in range(D):
        med[j] = float(clusters[i][3][j])
        for k in range(D):
            if k < j:
                corMat[k][j] = corMat[j][k]
                continue
            corMat[k][j] = float(clusters[i][4][j*D + k])*sqrt(float(clusters[i][2][j]))*sqrt(float(clusters[i][2][k]))
    pointsLocal = np.random.multivariate_normal(med, corMat, clusters[i][0])
    sourceClusters.append(Cluster(pointsLocal[0]))
    for j in range(len(pointsLocal)):
        points.append(pointsLocal[j])
        sourceClusters[i].appendPoint(pointsLocal[j])
    #plt.plot(x, y, 'b.', markersize = marker)

f = open("lab8_output_0.txt", "w")
f.write(str(len(sourceClusters)) + "\n")
for c in sourceClusters:
    f.write(str(c.count) + getStrPoint(c.getMedium()) + getStrPoint(c.getDisp()) + "\n")
    plt.plot(c.x(), c.y(), 'b.', markersize = marker, color = c.color)
plt.savefig("static/img/img_lab8_0.png")

plt.figure(2)
clusters1 = []
for i in range(len(points)):
    idx = cmp_len(points[i], clusters1, T)
    if idx == -1:
        clusters1.append(Cluster(points[i]))
    else:
        clusters1[idx].appendPoint(points[i])

print(len(clusters1))
f = open("lab8_output_1.txt", "w")
f.write(str(len(clusters1)) + "\n")
for i in range(len(clusters1)):
    c = clusters1[i]
    f.write(str(c.count) + getStrPoint(c.getMedium()) + getStrPoint(c.getDisp()) + "\n")
    plt.plot(c.x(), c.y(), 'b.', markersize = marker, color = c.color)
plt.savefig("static/img/img_lab8_1")

plt.figure(3)
usedPoints = []
for i in range(len(points)):
    usedPoints.append(False)
clusters2 = []
clusters2.append(Cluster(points[0]))
usedPoints[0] = True
minDistance = -1
while True:
    dist = 0
    idx = -1
    for i in range(len(points)):
        if usedPoints[i]:
            continue
        minDist = -1
        for j in range(len(clusters2)):
            newDistance = distance(clusters2[j].center, points[i])
            if minDist == -1 or newDistance < minDist:
                minDist = newDistance
        if minDist > dist:
            idx = i
            dist = minDist

    if idx != -1 and len(clusters2) >= 2 and dist > minDistance * Alpha:
        newDistance = 0
        for i in range(len(clusters2)):
            newDistance += distance(clusters2[i].center, points[idx])
        lenClusters = len(clusters2)
        minDistance = (minDistance*(lenClusters*(lenClusters - 1) / 2) + newDistance) / ((lenClusters + 1)*lenClusters / 2)
        clusters2.append(Cluster(points[idx]))
        usedPoints[idx] = True
    elif len(clusters2) == 1:
        clusters2.append(Cluster(points[idx]))
        minDistance = distance(clusters2[0].center, clusters2[1].center)
    else:
        break;

for i in range(len(points)):
    if not usedPoints[i]:
        dist = minDistance
        idx = -1
        for j in range(len(clusters2)):
            newDist = distance(clusters2[j].center, points[i])
            if idx == -1 or newDist < dist:
                idx = j
                dist = newDist
        usedPoints[i] = True
        clusters2[idx].appendPoint(points[i])
print(len(clusters2))
f = open("lab8_output_2.txt", "w")
f.write(str(len(clusters2)) + "\n")
for c in clusters2:
    f.write(str(c.count) + getStrPoint(c.getMedium()) + getStrPoint(c.getDisp()) + "\n")
    plt.plot(c.x(), c.y(), 'b.', markersize = marker, color = c.color)
plt.savefig("static/img/img_lab8_2")

clusters3 = []
idxs = []
while len(idxs) < K:
    newIdx = random.randint(0, len(points) - 1)
    flag = True
    for j in range(len(idxs)):
        flag &= idxs[j] != newIdx
    if flag:
        idxs.append(newIdx)

for p in usedPoints:
    p = False

lastCenters = []
for idx in idxs:
    clusters3.append(Cluster(points[idx]))
    usedPoints[idx] = True
    lastCenters.append(points[idx])

count = 0
f = open("lab8_output_3.txt", "w")
while True:
    for i in range(len(points)):
        minDist = 0
        idx = -1
        for j in range(len(clusters3)):
            newDist = distance(clusters3[j].center, points[i])
            if idx == -1 or newDist < minDist:
                idx = j
                minDist = newDist
        #print(idx)
        clusters3[idx].appendPoint(points[i])

    flag = True
    for i in range(len(clusters3)):
        clusters3[i].redefineCenter()
        flag &= distance(lastCenters[i], clusters3[i].center) == 0
        lastCenters[i] = clusters3[i].center

    if not flag:
        print('bdc')
        plt.figure(400 + count)
        f.write(str(len(clusters3)) + "\n")
        for i in range(len(clusters3)):
            c = clusters3[i]
            f.write(str(c.count) + getStrPoint(c.getMedium()) + getStrPoint(c.getDisp()) + "\n")
            plt.plot(c.x(), c.y(), 'b.', markersize = marker, color = c.color)
        plt.savefig("static/img/img_lab8_3_" + str(count))
        count += 1
        for i in range(len(clusters3)):
            clusters3[i].redefineCenter()
            clusters3[i].flush()
    else:
        break
plt.figure(400)
#ansStack.append(copy(clusters3))
print(len(clusters3))
f.write(str(len(clusters3)) + "\n")
for i in range(len(clusters3)):
    c = clusters3[i]
    f.write(str(c.count) + getStrPoint(c.getMedium()) + getStrPoint(c.getDisp()) + "\n")
    plt.plot(c.x(), c.y(), 'b.', markersize = marker, color = c.color)
plt.savefig("static/img/img_lab8_3_" + str(count))

#plt.show()
    `)
}
