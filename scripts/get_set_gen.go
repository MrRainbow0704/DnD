package main

import (
	"flag"
	"fmt"
	"strings"
)

func title(s string) string {
	ss := strings.Split(s, "_")
	t := make([]string, len(ss))
	for i, w := range ss {
		t[i] = strings.ToTitle(w[:1]) + strings.ToLower(w[1:])
	}
	return strings.Join(t, "")
}

func main() {
	var table = flag.String("t", "", "Table name")
	var columns = flag.String("c", "", "Columns names comma separated")
	flag.Parse()
	cols := strings.Split(*columns, ",")
	out := make([]string, len(cols))
	for i, col := range cols {
		col = strings.TrimSpace(col)
		s := make([]string, 5)
		s[0] = "-- name: Set" + title((*table)[0:len(*table)-1]) + title(col) + " :one"
		s[1] = "UPDATE " + *table
		s[2] = "SET " + col + " = sqlc.arg(" + col + ")"
		s[3] = "WHERE (id = sqlc.arg(id))"
		s[4] = "RETURNING *;"
		out[i] = strings.Join(s, "\n")
	}
	fmt.Print(strings.Join(out, "\n\n"))
}
