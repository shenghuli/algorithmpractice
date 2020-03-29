package main

import (
  "fmt"
  "math" 
)


func main() {
	//剪绳子
	x := 18
	fmt.Println(MaxProduct(x))
}

//题目：给你一根长度为n的绳子，请把绳子剪成m段（m、n都是整数，n>1且m>1),每段绳子的长度标记为k[0],k[1],...
//k[m].请问k[0]*k[1]*...*k[m]可能的最大乘机是多少？例如当绳子的长度为8时，我们把他剪成的长度分别为2、3、3
//的三段，此时得到的最大乘积为18

//分析 f(n) = max(f(i)*f(n-i) 0<i<n
//MaxProduct 计算最大乘积,动态规划解法
func MaxProduct(lenght int) int {
	//绳子必须剪一刀
	if lenght < 2 {
		return 0
	}
	if lenght == 2 {
		return 1
	}
	if lenght == 3 {
		return 2
	}
	products := make([]int, lenght+1)
	//初始化，迭代的基础中，可以不剪
	products[0] = 0
	products[1] = 1
	products[2] = 2
	products[3] = 3
	for i := 4; i <= lenght; i++ {
		//求解长度为i的最大乘积
		max := 0
		for j := 1; j <= i/2; j++ {
			product := products[j] * products[i-j]
			if max < product {
				max = product
			}
		}
		products[i] = max
	}
	return products[lenght]
}

//MaxProductGreedy
func MaxProductGreedy(length int) int {
	if length < 2 { //无法剪
		return 0
	}
	if length == 2 {
		return 1
	}
	if length == 3 {
		return 2
	}
	//通过数学分析，尽可能剪多的3
	timesOf3 := length / 3
	//剪完3，剩1，此时剪成两个2，乘积更大
	if length%3 == 1 {
		timesOf3--
	}
	timesOf2 := (length - timesOf3*3) / 2
	return int(math.Pow(3, float64(timesOf3))) * int(math.Pow(2, float64(timesOf2)))
}

