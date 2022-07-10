package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

/*
Ученые разработали новый материал неизвестной прочности. Они знают, что материал разбивается при падении с высоты от 1 метра до 5 000 метров. Но не знают, с какой именно высоты. Чтобы определить прочность, ученые поднимают предмет на некоторую высоту и сбрасывают его оттуда. Их задача — определить, начиная с какой именно высоты предмет начнет разбиваться.
Специальная платформа, с помощью которой они осуществляют эксперимент, скидывает предмет только с дискретных высот (1, 2, 3 ... 4999, 5000 метров — платформа не может скинуть предмет, например, с 2,5 метров. Точности в 1 метр ученым вполне достаточно). При падении с высоты "n" метров предмет уничтожается. Если же его сбрасывали с высоты ниже "n", то его можно использовать в повторных экспериментах.
Нужно АБСОЛЮТНО ТОЧНО найти ту высоту, начиная с которой предметы разрушаются. Сделать это нужно за МИНИМАЛЬНО возможное число экспериментов. У ученых при этом всего 2 предмета, но они абсолютно одинаковые. Каким образом этого можно достигнуть? Сколько экспериментов при этом максимально потребуется?


*/

func main() {

	rand.Seed(time.Now().UnixNano())
	heightBroken := 1 + rand.Intn(5000-1)

	count, height := experiment(5000, heightBroken)
	fmt.Printf("Колличество экспериментов: %d  при высоте разрушения предмета: %d метров ", count, height)
}

func experiment(heightMax, heightBroken int) (counterExperiment, heightCast int) {

	heightTest := steps(heightMax)
	countStep := heightTest
	counterExperiment = 1
	for heightBroken > heightTest {
		counterExperiment++
		countStep -= 1
		heightTest += countStep
	}
	heightTest -= countStep

	for heightBroken > heightTest {
		counterExperiment++
		heightTest++
	}

	return counterExperiment, heightTest
}

func steps(height int) (stepMax int) {

	// Using the discriminant formula , we solve the equation:
	// x^x + x -2height = 0
	stepMax = (-1 - int(math.Sqrt(float64(1+4*2*height)))) / 2
	return int(math.Abs(float64(stepMax)))
}
