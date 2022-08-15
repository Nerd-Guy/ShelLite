package shellite

import (
	"fmt"
	"strings"
)

type Table struct {
	Columns []string
	Rows    [][]string
}

func NewTable() *Table {
	var table Table = Table{}
	return &table
}

func (table *Table) AddColumn(name string) {
	table.Columns = append(table.Columns, name)
}

func (table *Table) AddRow(data ...string) {
	table.Rows = append(table.Rows, data)
}

func (table *Table) Print() {
	MaxDistances := make(map[string]int)
	for _, v := range table.Rows {
		for i, v2 := range v {
			if len(v2) > MaxDistances[table.Columns[i]] {
				MaxDistances[table.Columns[i]] = len(v2) + 3
			}
		}
	}
	for _, v := range table.Columns {
		dist := MaxDistances[v] - len(v)
		if dist < 0 {
			dist = -dist
		}
		fmt.Printf("%s%s", v, strings.Repeat(" ", dist))
	}
	fmt.Println()
	for _, v := range table.Rows {
		for i, v2 := range v {
			fmt.Printf("%s%s", v2, strings.Repeat(" ", MaxDistances[table.Columns[i]]-len(v2)))
		}
		fmt.Println()
	}
}
