package main

import (
	"fmt"

	"github.com/uber/h3-go/v4"
)

func main() {
	/*
		latLng := h3.NewLatLng(37.775938728915946, -122.41795063018799)
		resolution := 9 // between 0 (biggest cell) and 15 (smallest cell)

		cell := h3.LatLngToCell(latLng, resolution)

		fmt.Printf("%s\n", cell)
	*/
	// Output:
	// 8928308280fffff

	// uxirkffr.test.hex.camp
	hex := "8a5d11514b1ffff"

	cell2 := h3.Cell(h3.IndexFromString(hex))
	fmt.Printf("cell2 %s\n", cell2)
	resolution := cell2.Resolution()
	fmt.Printf("cell2 resolution %v\n", resolution)
	base := cell2.BaseCellNumber()
	fmt.Printf("cell2 base %v\n", base)
	for i := resolution; i > 0; i-- {
		parent := cell2.Parent(i)
		childPos := parent.ChildPos(i - 1)
		fmt.Printf("  %d: Parent %s %d: %d\n", i, parent.String(), parent.Resolution(), childPos)
	}
}

// 3.4.5.4.2.4.2.1.2.4.46.h3.test.hex.camp
