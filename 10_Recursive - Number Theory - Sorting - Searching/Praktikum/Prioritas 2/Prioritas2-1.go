package main

import "fmt"

func playingDomino(cards [][]int, deck []int) interface{} {
	// Membuat variabel temp untuk menyimpan nilai kartu yang bisa dikeluarkan
	temp := make([][]int, 0)
	//Perulangan untuk melihat kartu mana saja yang dapat dikeluarkan
	for i := range cards {
		if cards[i][0] == deck[0] || cards[i][0] == deck[1] {
			temp = append(temp, cards[i])
		} else if cards[i][1] == deck[0] || cards[i][1] == deck[1] {
			temp = append(temp, cards[i])
		}
	}
	//len(temp)==0 itu berarti tidak ada kartu yang dapat keluar
	if len(temp) == 0 {
		return "tutup kartu"
	}
	//Variabel max untuk menyimpan nilai kartu yang terbesar
	max := make([]int, 2)
	//Perulangan untuk melihat kartu mana yang terbesar
	for i := range temp {
		if temp[i][0]+temp[i][1] >= max[0]+max[1] {
			max[0] = temp[i][0]
			max[1] = temp[i][1]
		}
	}
	return max
}

func main() {
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 4}, {2, 1}, {3, 3}}, []int{4, 3}))
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 3}, {3, 4}, {2, 1}}, []int{3, 6}))
	fmt.Println(playingDomino([][]int{{6, 6}, {2, 4}, {3, 6}}, []int{5, 1}))
}
