package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal"
)

func main() {
	fmt.Println("Смирнов Глеб Денисович 1/280")
	fmt.Println(internal.TaskA(1.1, 3.6, 0.5))
	fmt.Println(internal.TaskB([]float64{2.35, 1.46, 1.36, 1.28, 1.2}))
}
