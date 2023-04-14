package main

import (
	"fmt"
)

func main() {
	fmt.Println("Вы выехали из Москвы и опаздываете на важную встречу в Рязань.")
	fmt.Println("У вас есть 2 часа, чтобы добраться, а это 200км. пути.")
	fmt.Println("Укажите, с какой скоростью вы едете? (км/ч)")

	var carSpeed int
	fmt.Scan(&carSpeed)

	distancePassed := carSpeed * 2
	distanceLeft := 200 - distancePassed
	timeLeftMin := int(((float64)(distancePassed-200) / (float64)(carSpeed)) * 60)

	if distancePassed < 200 {
		fmt.Println("За 2 часа вы проехали", distancePassed, "километров.")
		fmt.Println("А до Рязани еще", distanceLeft, "километров!!")
		fmt.Println("Вы опоздали.")
	}

	if distancePassed > 200 {
		fmt.Println("Вы успешно добрались на встречу вовремя!")
		fmt.Println("К тому же, успели выпить кофе, так как сэкономили", timeLeftMin, "минут!")
	}

	if distancePassed == 200 {
		fmt.Println("Вы успешно добрались на встречу вовремя! Прям впритык!")
	}
}

/*В строке timeLeftMin я сжульничал и погуглил код. Так как я понимал, что там образуется
ошибка из-за не правильного типа данных, целые числа и т.д., но сам еще не сталкивался
с числамя с запятой (float), то решил подсмотреть, как было бы правильно*/
