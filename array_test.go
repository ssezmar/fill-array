package main

import (
	"testing"
	"reflect"
)

func TestNewIntArray(t *testing.T) {
	testTable := []struct {
		name			string
		sizeX, sizeY 	int 
		wantErr 		bool
	}{
		{
			name: "Ok",
			sizeX: 2,
			sizeY: 2,
		},
		{
			name: "SizeX is 0",
			sizeX: 0,
			wantErr: true,
		},
		{
			name: "SizeX is negative",
			sizeX: -1,
			wantErr: true,
		},
		{
			name: "SizeY is 0",
			sizeX: 2,
			sizeY: 0,
			wantErr: true,
		},
		{
			name: "SizeY is negative",
			sizeX: 2,
			sizeY: -1,
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := NewIntArray(testCase.sizeX, testCase.sizeY)
			if err != nil && !testCase.wantErr {
				t.Fatalf("Unexpected error: %s", err.Error())
			}

			if err == nil && testCase.wantErr {
				t.Fatal("Error was expected, but got nil")
			}

		})
	}

}

func TestIntArrayFill(t *testing.T) {
	testTable := []struct {
		name 			string
		sizeX, sizeY 	int
		values 			[]int
		wantErr 		bool
	}{
		{
			name: "Ok", 
			sizeX: 2,
			sizeY: 2,
			values: []int {
				1, 2, 3, 5,
			},
		},
		{
			name: "Values Array has less elements", 
			sizeX: 2,
			sizeY: 2,
			values: []int{1, 2, 3},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			IntArray, _ := NewIntArray(testCase.sizeX, testCase.sizeY)
			err := IntArray.Fill(testCase.values)
			if err != nil && !testCase.wantErr {
				t.Fatal("Unexpected Fill Error")
			}

			if err == nil && testCase.wantErr {
				t.Fatal("Initialization Error was expected, but got nil")
			}
		})
	}

}

func TestIntArrayGet(t *testing.T) {
	testTable := []struct {
		name	string
		sizeX, sizeY int
		values []int
		result [][]int
		wantGetErr bool
	}{
		{
			name: "Ok", 
			sizeX: 2,
			sizeY: 2,
			values: []int{1, 2, 3, 4},
			result: [][]int{
				{1, 2},
				{3, 4},
			},
		},
		{
			name: "Empty Array", 
			sizeX: 2,
			sizeY: 2,
			wantGetErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T){
			arr, _ := NewIntArray(testCase.sizeX, testCase.sizeY)

			if testCase.wantGetErr {
				_, err := arr.Get()
				if err == nil {
					t.Fatal("Expected error, but got nil")
				}

				t.Skip("Ok")
			}

			err := arr.Fill(testCase.values)
			if err != nil {
				t.Fatal("Unexpected Fill error")
			}

			res, err := arr.Get()
			if err != nil {
				t.Fatal("Unexpected Get error")
			}

			if !reflect.DeepEqual(res, testCase.result) {
				t.Fatalf("Incorrect Result, Wnt %v, Got %v", testCase.result, res)
			}

		})
	}

}