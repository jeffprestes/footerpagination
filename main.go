package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println(printFooter(4, 5, 1, 0))
	fmt.Println(printFooter(4, 10, 2, 2))
}

func printFooter(currentPage int, totalPages int, boundaries int, around int) string {

	var strTemp strings.Builder

	var gapFilledBeforeCurrentPage, gapFilledAfterCurrentPage bool

	for x := 1; x <= totalPages; x++ {
		if isWithinBoundaries(x, currentPage, totalPages, boundaries) || isWithinAroundRange(x, currentPage, around) {
			strTemp.WriteString(strconv.Itoa(x))
			strTemp.WriteString(" ")
			continue

		} else {
			ok := shouldPrintGap(x, gapFilledBeforeCurrentPage, gapFilledAfterCurrentPage, currentPage)
			if ok {
				if x < currentPage {
					gapFilledBeforeCurrentPage = true
				} else if x > currentPage {
					gapFilledAfterCurrentPage = true
				}
				strTemp.WriteString("... ")
			}
			continue
		}
	}
	return strings.TrimSpace(strTemp.String())
}

func isWithinAroundRange(x int, currentPage int, around int) (result bool) {
	if around == 0 {
		return
	} else if x <= (currentPage+around) || (x > (currentPage-around) && x < currentPage) {
		result = true
	}
	return
}

func isWithinBoundaries(x int, currentPage int, totalPages int, boundaries int) (result bool) {
	boundaries--
	if x == 1 || x == currentPage || x == totalPages || x == (1+boundaries) || x == (totalPages-boundaries) {
		result = true
	}
	return
}

func shouldPrintGap(x int, gapFilledBeforeCurrentPage bool, gapFilledAfterCurrentPage bool, currentPage int) (result bool) {
	if (gapFilledBeforeCurrentPage && x < currentPage) || (gapFilledAfterCurrentPage && x > currentPage) {
		return
	}
	result = true
	return
}
