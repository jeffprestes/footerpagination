package main

import "fmt"

func main() {

	printFooter(4, 10, 2, 2)
	printFooter(4, 5, 1, 0)
	printFooter(4, 20, 1, 0)

}

func printFooter(current_page int, total_pages int, boundaries int, around int) {

	var gapFilledBeforeCurrentPage, gapFilledAfterCurrentPage bool

	fmt.Print("\n\r")
	for x := 1; x <= total_pages; x++ {
		if isWithinBoundaries(x, current_page, total_pages, boundaries) {
			fmt.Printf("%d ", x)
			continue

		} else if isWithinAroundRange(x, current_page, around) {
			fmt.Printf("%d ", x)
			continue

		} else {
			ok := shouldPrintGap(x, gapFilledBeforeCurrentPage, gapFilledAfterCurrentPage, current_page)
			if ok {
				if x < current_page {
					gapFilledBeforeCurrentPage = true
				} else if x > current_page {
					gapFilledAfterCurrentPage = true
				}
				fmt.Print("... ")
			}
			continue
		}
	}
	fmt.Print("\n\r\n\r")
}

func isWithinAroundRange(x int, current_page int, around int) (result bool) {
	if around == 0 {
		return
	} else if x <= (current_page+around) || (x > (current_page-around) && x < current_page) {
		result = true
	}
	return
}

func isWithinBoundaries(x int, current_page int, total_pages int, boundaries int) (result bool) {
	boundaries--
	if x == 1 || x == current_page || x == total_pages || x == (1+boundaries) || x == (total_pages-boundaries) {
		result = true
	}
	return
}

func shouldPrintGap(x int, gapFilledBeforeCurrentPage bool, gapFilledAfterCurrentPage bool, current_page int) (result bool) {
	if (gapFilledBeforeCurrentPage && x < current_page) || (gapFilledAfterCurrentPage && x > current_page) {
		return
	}
	result = true
	return
}
