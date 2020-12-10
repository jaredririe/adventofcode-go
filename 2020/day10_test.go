package main

import (
	"testing"
)

func TestConnectAllAdapters(t *testing.T) {

	cases := []struct {
		input    []int
		expected int
	}{
		{
			input: []int{
				16,
				10,
				15,
				5,
				1,
				11,
				7,
				19,
				6,
				12,
				4,
			},
			expected: 35,
		},
		{
			input: []int{
				16,
				10,
				15,
				5,
				1,
				11,
				19,
				12,
				4,
			},
			expected: -1,
		},
		{
			input:    day10Input,
			expected: 2475,
		},
	}

	for i, c := range cases {
		actual := ConnectAllAdapters(c.input)
		if actual != c.expected {
			t.Errorf("[case %d] expected: %d, actual: %d", i, c.expected, actual)
		}
	}
}

func TestCountValidCombinations(t *testing.T) {

	cases := []struct {
		input    []int
		expected int
	}{
		{
			input: []int{
				16,
				10,
				15,
				5,
				1,
				11,
				7,
				19,
				6,
				12,
				4,
			},
			expected: 8,
		},
		{
			input: []int{
				28,
				33,
				18,
				42,
				31,
				14,
				46,
				20,
				48,
				47,
				24,
				23,
				49,
				45,
				19,
				38,
				39,
				11,
				1,
				32,
				25,
				35,
				8,
				17,
				7,
				9,
				4,
				2,
				34,
				10,
				3,
			},
			expected: 19208,
		},
		{
			input:    day10Input,
			expected: 442136281481216,
		},
	}

	for i, c := range cases {
		actual := CountValidCombinations(c.input)
		if actual != c.expected {
			t.Errorf("[case %d] expected: %d, actual: %d", i, c.expected, actual)
		}
	}
}

var day10Input = []int{
	48,
	171,
	156,
	51,
	26,
	6,
	80,
	62,
	65,
	82,
	130,
	97,
	49,
	31,
	142,
	83,
	75,
	20,
	154,
	119,
	56,
	114,
	92,
	33,
	140,
	74,
	118,
	1,
	96,
	44,
	128,
	134,
	121,
	64,
	158,
	27,
	17,
	101,
	59,
	12,
	89,
	88,
	145,
	167,
	11,
	3,
	39,
	43,
	105,
	16,
	170,
	63,
	111,
	2,
	108,
	21,
	146,
	77,
	45,
	52,
	32,
	127,
	147,
	76,
	58,
	37,
	86,
	129,
	57,
	133,
	120,
	163,
	138,
	161,
	139,
	71,
	9,
	141,
	168,
	164,
	124,
	157,
	95,
	25,
	38,
	69,
	87,
	155,
	135,
	15,
	102,
	70,
	34,
	42,
	24,
	50,
	68,
	169,
	10,
	55,
	117,
	30,
	81,
	151,
	100,
	162,
	148,
}
