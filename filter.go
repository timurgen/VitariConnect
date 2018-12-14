package visma

import (
	"errors"
	"fmt"
)

type Compare string

const (
	GreaterThanOrEqualTo Compare = "GreaterThanOrEqualTo"
	LessThanOrEqualTo    Compare = "LessThanOrEqualTo"
	EqualTo              Compare = "EqualTo"
	NotEqualTo           Compare = "NotEqualTo"
	GreaterThan          Compare = "GreaterThan"
	LessThan             Compare = "LessThan"
)

type FilterRow struct {
	Name            string
	CompareFunction Compare
	Value1          string
	Operator        string
}

type Filter struct {
	Rows []FilterRow
}

func CreateFilterRow(name string, compareFunc Compare, value1 string, operator string) FilterRow {
	return FilterRow{name, compareFunc, value1, operator}
}

func AddRowToFilter(filter *Filter, row FilterRow) error {
	if len(filter.Rows) > 0 && row.Operator == "" {
		return errors.New("row must have operator value as it not first row in this filter")
	} else if len(filter.Rows) == 0 && row.Operator != "" {
		return errors.New("first row must not have operator")
	}
	filter.Rows = append(filter.Rows, row)
	return nil
}

func RenderFilter(f Filter) string {
	result := "<Filters>"
	for _, row := range f.Rows {
		rowStr := fmt.Sprintf("<%s Compare=\"%s\" Value1=\"%s\" Operator=\"%s\"/>", row.Name,
			row.CompareFunction, row.Value1, row.Operator)
		result += rowStr
	}
	result += "</Filters>"
	return result
}
