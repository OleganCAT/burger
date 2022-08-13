package main

import (
	"fmt"
	"math/rand"
)

func breadToSandwich(b []string) string {
	var result string

	for {
		n := rand.Intn(len(b))
		if n == 0 {
			result = b[0]
		} else if n == 1 {
			result = b[1]
		} else {
			result = b[2]
		}
		return result
	}
}

func meatToSandwich(m []string) string {
	var result string

	for {
		n := rand.Intn(len(m))
		if n == 0 {
			result = m[0]
		} else if n == 1 {
			result = m[1]
		} else if n == 2 {
			result = m[2]
		} else {
			result = m[3]
		}
		return result
	}
}

func vegetablesToSandwich(v []string) string {
	var result string

	for {
		n := rand.Intn(len(v))
		if n == 0 {
			result = v[0]
		} else if n == 1 {
			result = v[1]
		} else {
			result = v[2]
		}
		return result
	}
}

func sousesToSandwich(s []string) string {
	var result string

	for {
		n := rand.Intn(len(s))
		if n == 0 {
			result = s[0]
		} else if n == 1 {
			result = s[1]
		} else {
			result = s[2]
		}
		return result
	}
}

func printToSandwich(b, m, v, s string) string {
	return fmt.Sprintf("Наш бургер будет состоять из %v, %v, %v и %v", b, m, v, s)
}

func main() {
	bread := []string{"чёрной булочки", "белой булочки", "ржаной булочки"}
	meat := []string{"свиной котлеты", "говяжой котлеты", "соевой котлеты", "куриной котлеты"}
	vegetables := []string{"огурцов", "помидоров", "листьев салата"}
	souses := []string{"томатного соуса", "горчичного соуса", "кисло-сладкого соуса"}

	resultBread := breadToSandwich(bread)
	resultMeat := meatToSandwich(meat)
	resultVegetable := vegetablesToSandwich(vegetables)
	resultSouses := sousesToSandwich(souses)

	res := printToSandwich(resultBread, resultMeat, resultVegetable, resultSouses)
	fmt.Println(res)

}
