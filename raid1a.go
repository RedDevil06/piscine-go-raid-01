package piscine

import (
	"github.com/01-edu/z01"
)

func Raid1a(x, y int) {
	if x == 1 && y == 1 {
		z01.PrintRune('o')
	} else if x > 1 {
		z01.PrintRune('o')
		for i := 1; i <= x-2; i++ {
			z01.PrintRune('-')
		}
		z01.PrintRune('o')
		z01.PrintRune('\n')
		if y > 1 {
			z01.PrintRune('|')
			for j := 1; j <= x-2; j++ {
				z01.PrintRune(' ')
			}
			z01.PrintRune('|')
			z01.PrintRune('\n')
		}
		z01.PrintRune('o')
		for i := 1; i <= x-2; i++ {
			z01.PrintRune('-')
		}
		z01.PrintRune('o')
		// z01.PrintRune('\n')
	} else if y > 1 {
		z01.PrintRune('o')
		if x == 1 {
			for k := 1; k <= y-2; k++ {
				z01.PrintRune('\n')
				z01.PrintRune('|')
			}
		}
		z01.PrintRune('\n')
		z01.PrintRune('o')
	}
	z01.PrintRune('\n')
}

// func main() {
// 	Raid1a(5, 3)
// 	Raid1a(1, 1)
// 	Raid1a(1, 5)
// }