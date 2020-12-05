package main

import (
	"strconv"
)

// F = 0, B = 1
// FBFBBFF = 0101100 = row 44
//
// L = 0, R = 1
// RLR = 101 = column 5
//
// seat ID: multiply the row by 8, then add the column.
// In this example, the seat has ID 44 * 8 + 5 = 357

func VacantSeats(input []string) []int64 {

	seatIDs := parseSeats(input)

	var highestSeatID int64
	seatMap := make(map[int64]bool)

	for _, seatID := range seatIDs {
		seatMap[seatID] = true
		if seatID > highestSeatID {
			highestSeatID = seatID
		}
	}

	var vacantSeats []int64

	for i := int64(0); i < highestSeatID; i++ {
		if _, ok := seatMap[i]; !ok {
			vacantSeats = append(vacantSeats, i)
		}
	}

	return vacantSeats
}

func HighestSeatID(input []string) int64 {

	seatIDs := parseSeats(input)

	var highestSeatID int64
	for _, seatID := range seatIDs {
		if seatID > highestSeatID {
			highestSeatID = seatID
		}
	}

	return highestSeatID
}

func parseSeats(input []string) []int64 {

	var seatIDs []int64

	for _, seat := range input {

		rowKey := seat[0:7]
		columnKey := seat[7:10]

		var rowBinary, columnBinary string

		for _, r := range rowKey {
			if r == rune('F') {
				rowBinary += "0"
			} else if r == rune('B') {
				rowBinary += "1"
			}
		}

		for _, r := range columnKey {
			if r == rune('L') {
				columnBinary += "0"
			} else if r == rune('R') {
				columnBinary += "1"
			}
		}

		row, _ := strconv.ParseInt(rowBinary, 2, 64)
		column, _ := strconv.ParseInt(columnBinary, 2, 64)

		seatID := row*8 + column
		seatIDs = append(seatIDs, seatID)
	}

	return seatIDs
}
