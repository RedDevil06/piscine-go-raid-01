package piscine

import (
	"github.com/01-edu/z01"
)

func Raid1d(x, y int) {
	if x > 0 && y > 0 {
		if x == 1 && y == 1 {
			z01.PrintRune('A')
			z01.PrintRune('\n')
		} else if x > 1 {
			z01.PrintRune('A')
			for i := 1; i <= x-2; i++ {
				z01.PrintRune('B')
			}
			z01.PrintRune('C')
			z01.PrintRune('\n')
			if y == 2 {
				z01.PrintRune('A')
				for i := 1; i <= x-2; i++ {
					z01.PrintRune('B')
				}
				z01.PrintRune('C')
				z01.PrintRune('\n')
			} else if y > 2 {
				//z01.PrintRune('\n')
				for l := 3; l <= y; l++ {
					z01.PrintRune('B')
					for j := 1; j <= x-2; j++ {
						z01.PrintRune(' ')
					}
					z01.PrintRune('B')
					z01.PrintRune('\n')
				}
				z01.PrintRune('A')
				for i := 1; i <= x-2; i++ {
					z01.PrintRune('B')
				}
				z01.PrintRune('C')
				z01.PrintRune('\n')
			}
			// z01.PrintRune('\n')
		} else if y > 1 {
			z01.PrintRune('A')
			if x == 1 {
				for k := 1; k <= y-2; k++ {
					z01.PrintRune('\n')
					z01.PrintRune('B')
				}
			}
			z01.PrintRune('\n')
			z01.PrintRune('C')
			z01.PrintRune('\n')
		}
	} else {
	}
}

// func main() {
// 	Raid1d(5, 3)
// 	Raid1d(5, 1)
// 	Raid1d(1, 1)
// 	Raid1d(1, 5)
// }
