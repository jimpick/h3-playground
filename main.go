package main

import (
	"encoding/base32"
	"encoding/hex"
	"fmt"
	"log"
	"strings"

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

	// str := "ONXW2ZJAMRQXIYJAO5UXI2BAAAQGC3TEEDX3XPY="
	str := strings.ToUpper("uxirkffr")
	data, err := base32.StdEncoding.DecodeString(str)
	if err != nil {
		log.Fatal(err)
	}
	hex2 := strings.ReplaceAll(fmt.Sprintf("8%-14s", hex.EncodeToString(data)), " ", "f")
	fmt.Printf("hex: %v\n", hex2)

	// hex3 := "8a5d11514b1ffff"
	hex3 := hex2

	cell2 := h3.Cell(h3.IndexFromString(hex3))
	fmt.Printf("cell2 %s\n", cell2)
	resolution := cell2.Resolution()
	fmt.Printf("cell2 resolution %v\n", resolution)
	base := cell2.BaseCellNumber()
	for i := resolution; i > 0; i-- {
		parent := cell2.Parent(i)
		childPos := parent.ChildPos(i - 1)
		fmt.Printf("  %d: Parent %s %d: %d\n", i, parent.String(), parent.Resolution(), childPos)
	}
	fmt.Printf("cell2 base %v\n", base)
}

// 3.4.5.4.2.4.2.1.2.4.46.h3.test.hex.camp
