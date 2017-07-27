package robot

import "fmt"

//Table struct has Width (Y representation) and Length (X representation)
type Table struct {
	Width, Length int
}

//NewTable is to initantiate Table object
func NewTable(width int, length int) (Table, error) {
	if width < 0 || length < 0 {
		return Table{}, fmt.Errorf(errFailToInitiate)
	}
	return Table{Width: width, Length: length}, nil
}
