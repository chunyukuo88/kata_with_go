package main

import "fmt"

func getDayMap() map[int][]string {
	dayMap := map[int][]string{
		1:  {"first", "a Partridge in a Pear Tree"},
		2:  {"second", "two Turtle Doves"},
		3:  {"third", "three French Hens"},
		4:  {"fourth", "four Calling Birds"},
		5:  {"fifth", "five Gold Rings"},
		6:  {"sixth", "six Geese-a-Laying"},
		7:  {"seventh", "seven Swans-a-Swimming"},
		8:  {"eighth", "eight Maids-a-Milking"},
		9:  {"ninth", "nine Ladies Dancing"},
		10: {"tenth", "ten Lords-a-Leaping"},
		11: {"eleventh", "eleven Pipers Piping"},
		12: {"twelfth", "twelve Drummers Drumming"},
	}
	return dayMap
}

func Verse(i int) string {
	dayMap := getDayMap()
	return "On the " + dayMap[i][0] + " day of Christmas my true love gave to me: "
}

func concatenateGifts(i int) string {
	dayMap := getDayMap()
	var gifts string
	for index := i; index > 0; index-- {
		if index == 1 {
			gifts += (dayMap[index][1] + ". ")
		} else {
			gifts += (dayMap[index][1] + ", ")
		}
	}
	return gifts
}

func Song() string {
	dayMap := getDayMap()
	var song string
	for i := range dayMap {
		song += Verse(i) + concatenateGifts(i)
	}
	return song
}

func main() {
	song := Song()
	fmt.Println(song)
}
