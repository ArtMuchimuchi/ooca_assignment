package main

import (
	"errors"
	"testing"
)

func TestDoublesDiscount(t *testing.T) {
	testCases := []struct {
		name     string
		input    Order
		expected float32
		err      error
	}{
		{"double discount (no doubble)", Order{RedSet: 1,
			GreenSet:  0,
			BlueSet:   0,
			YellowSet: 0,
			PinkSet:   0,
			PurpleSet: 4,
			OrangeSet: 0,
			IsMember:  true}, 0.0, nil},
		{"double discount 1", Order{RedSet: 0,
			GreenSet:  0,
			BlueSet:   0,
			YellowSet: 0,
			PinkSet:   0,
			PurpleSet: 4,
			OrangeSet: 2,
			IsMember:  true}, 5.0, nil},
		{"double discount 2", Order{RedSet: 1,
			GreenSet:  2,
			BlueSet:   0,
			YellowSet: 0,
			PinkSet:   2,
			PurpleSet: 4,
			OrangeSet: 2,
			IsMember:  true}, 5.0, nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := doublesDiscount(tc.input)
			if result != tc.expected || err != tc.err {
				t.Errorf("Test %s failed. Expected %f, got %f", tc.name, tc.expected, result)
			}
		})
	}
}

func TestMemberDiscount(t *testing.T) {
	testCases := []struct {
		name     string
		input    Order
		expected float32
		err      error
	}{
		{"member discount (no member)", Order{RedSet: 1,
			GreenSet:  0,
			BlueSet:   0,
			YellowSet: 0,
			PinkSet:   0,
			PurpleSet: 4,
			OrangeSet: 0,
			IsMember:  false}, 0.0, nil},
		{"member discount", Order{RedSet: 0,
			GreenSet:  0,
			BlueSet:   0,
			YellowSet: 0,
			PinkSet:   0,
			PurpleSet: 4,
			OrangeSet: 2,
			IsMember:  true}, 10.0, nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := memberDiscount(tc.input)
			if result != tc.expected || err != tc.err {
				t.Errorf("Test %s failed. Expected %f, got %f", tc.name, tc.expected, result)
			}
		})
	}
}

func TestCalDiscount(t *testing.T) {
	testCases := []struct {
		name          string
		totalAmout    float32
		discount      float32
		afterDiscount float32
		err           error
	}{
		{"calculate discount (normal) 1", 520, 15, 442, nil},
		{"calculate discount (normal) 2", 120, 50, 60, nil},
		{"calculate discount (negative)", -520, -15, 0, errors.New("total amount or discount can not be negative.")},
		{"calculate discount (exceed discount)", 300, 200, 0, errors.New("total discount can not exceed 100.")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := calDiscount(tc.totalAmout, tc.discount)
			if result != tc.afterDiscount || !((err != nil) == (tc.err != nil)) {
				t.Errorf("Test %s failed. Expected %f, got %f\n", tc.name, tc.afterDiscount, result)
				t.Errorf("Expected error %v, got %v", tc.err, err)
			}
		})
	}
}

func TestCalTotalAmount(t *testing.T) {
	testCases := []struct {
		name     string
		input    Order
		expected float32
		err      error
	}{
		{"total amount 1", Order{RedSet: 1,
			GreenSet:  1,
			BlueSet:   1,
			YellowSet: 1,
			PinkSet:   1,
			PurpleSet: 1,
			OrangeSet: 1,
			IsMember:  true}, 460.0, nil},
		{"total amount 2", Order{RedSet: 0,
			GreenSet:  0,
			BlueSet:   0,
			YellowSet: 0,
			PinkSet:   0,
			PurpleSet: 0,
			OrangeSet: 0,
			IsMember:  true}, 0.0, nil},
		{"total amount 3", Order{RedSet: 1,
			GreenSet:  2,
			BlueSet:   0,
			YellowSet: 0,
			PinkSet:   2,
			PurpleSet: 1,
			OrangeSet: 1,
			IsMember:  true}, 500.0, nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := calTotalAmount(tc.input)
			if result != tc.expected || err != tc.err {
				t.Errorf("Test %s failed. Expected %f, got %f", tc.name, tc.expected, result)
			}
		})
	}
}
