package main

func main() {
	QuadA(5, 3)
}

func QuadA(x, y int) {

	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if (i == 1 || i == y) && (j == 1 || j == x) {
				print("o")
			} else if i == 1 || i == y {
				print("-")
			} else if j == 1 || j == x {
				print("|")
			} else if i != 1 && i != y && j != 1 && j != x {
				
				print(" ")
			}

		}
		println()

	}

}
