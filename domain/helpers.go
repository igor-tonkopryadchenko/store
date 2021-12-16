package domain

import "fmt"

func (x *YearWeek) HfWeek() string {
	return fmt.Sprintf("%d-W%d", x.Year, x.Week)
}
