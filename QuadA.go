package main

func main() {
	QuadA(5, 3)
}

func QuadA(x, y int) {

	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if (i == 1 || i == y) && (j == 1 || j == x) {
				// Print 'o' at the corners
				print("o")
			} else if i == 1 || i == y {
				// Print '-' at the top and bottom borders
				print("-")
			} else if j == 1 || j == x {
				// Print '|' at the left and right borders
				print("|")
			} else if i != 1 && i != y && j != 1 && j != x {
				// Print space inside the rectangle
				print(" ")
			}

		}
		println()

	}

}
