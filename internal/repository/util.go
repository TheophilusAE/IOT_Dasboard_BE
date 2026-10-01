package repository

import "strconv"

func placeholder(n int) string {
	return "$" + strconv.Itoa(n)
}
