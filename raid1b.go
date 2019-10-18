package piscine

import (
	"github.com/01-edu/z01"
)

func Raid1b(x, y int) {
	if x > 0 && y > 0 {
		if x == 1 && y == 1 {
			z01.PrintRune('/')
			z01.PrintRune('\n')
		} else if x > 1 {
			z01.PrintRune('/')
			for i := 1; i <= x-2; i++ {
				z01.PrintRune('*')
			}
			z01.PrintRune('\\')
			z01.PrintRune('\n')
			if y == 2 {
				z01.PrintRune('/')
				for i := 1; i <= x-2; i++ {
					z01.PrintRune('*')
				}
				z01.PrintRune('\\')
				z01.PrintRune('\n')
			} else if y > 2 {
				//z01.PrintRune('\n')
				for l := 3; l <= y; l++ {
					z01.PrintRune('*')
					for j := 1; j <= x-2; j++ {
						z01.PrintRune(' ')
					}
					z01.PrintRune('*')
					z01.PrintRune('\n')
				}
				z01.PrintRune('\\')
				for i := 1; i <= x-2; i++ {
					z01.PrintRune('*')
				}
				z01.PrintRune('/')
				z01.PrintRune('\n')
			}
			// z01.PrintRune('\n')
		} else if y > 1 {
			z01.PrintRune('/')
			if x == 1 {
				for k := 1; k <= y-2; k++ {
					z01.PrintRune('\n')
					z01.PrintRune('*')
				}
			}
			z01.PrintRune('\n')
			z01.PrintRune('\\')
			z01.PrintRune('\n')
		}
	} else {
	}
}

// func main() {
// 	Raid1b(5, 1)
// 	Raid1b(1, 1)
// 	Raid1b(1, 5)
// }
