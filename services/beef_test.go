package services_test

import (
	"pieFireDire/models"
	"pieFireDire/services"
	"reflect"
	"testing"
)

func TestBeef(t *testing.T) {

	testCase := []struct {
		name   string
		input  string
		expect map[string]int
	}{
		{
			name:  "case1",
			input: "Fatback t-bone t-bone, pastrami  ..   t-bone.  pork, meatloaf jowl enim.  Bresaola t-bone.",
			expect: map[string]int{
				"bresaola": 1, "enim": 1, "fatback": 1, "jowl": 1, "meatloaf": 1, "pastrami": 1, "pork": 1, "t-bone": 4,
			},
		},
		{
			name:  "case2",
			input: "Fatback t-bone t-bone, pastrami  ..   t-bone.  pork, meatloaf jowl enim.  Bresaola t-bone. Bresaola",
			expect: map[string]int{
				"enim": 1, "fatback": 1, "jowl": 1, "meatloaf": 1, "pastrami": 1, "pork": 1, "t-bone": 4, "bresaola": 2,
			},
		},
		{
			name:  "case3",
			input: "Fatback t-bone t-bone, pastrami  ..   t-bone.  pork, meatloaf jowl enim.  Bresaola t-bone. Bresaola",
			expect: map[string]int{
				"enim": 1, "fatback": 1, "jowl": 1, "meatloaf": 1, "pastrami": 1, "pork": 1, "t-bone": 4, "bresaola": 2,
			},
		},
	}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			bf := services.NewBeefService()
			bf.Count(v.input)
			if !reflect.DeepEqual(v.expect, bf.(*services.Beef).CounterList) {
				t.Errorf("Case: %v: expected %v, actual %v", v.name, v.expect, bf.(*services.Beef).CounterList)
			}
		})
	}

}

func TestBeefReadMultiLine(t *testing.T) {
	input1 := "Fatback t-bone t-bone, pastrami  ..   t-bone.  pork, "
	input2 := "meatloaf jowl enim.  Bresaola t-bone. Bresaola"
	expect := map[string]int{
		"enim": 1, "fatback": 1, "jowl": 1, "meatloaf": 1, "pastrami": 1, "pork": 1, "t-bone": 4, "bresaola": 2}
	bf := services.NewBeefService()
	bf.Count(input1)
	bf.Count(input2)
	if !reflect.DeepEqual(expect, bf.Get()) {
		t.Errorf("Case: %v: expected %v, actual %v", "test input multi line", expect, bf.(*services.Beef).CounterList)
	}
}

func TestGet(t *testing.T) {
	input := "Fatback t-bone t-bone, pastrami  ..   t-bone.  pork, meatloaf jowl enim.  Bresaola t-bone. Bresaola"
	expected := map[string]int{
		"enim": 1, "fatback": 1, "jowl": 1, "meatloaf": 1, "pastrami": 1, "pork": 1, "t-bone": 4, "bresaola": 2,
	}
	bf := services.NewBeefService()
	bf.Count(input)
	got := bf.Get()
	if !reflect.DeepEqual(expected, bf.Get()) {
		t.Errorf("Case: Get result, expected %v got %v", expected, got)
	}
}

func TestBeefReadfileError(t *testing.T) {
	bf := services.NewBeefService()
	err := bf.Read("./path/error")
	if err != models.ErrOpenFile {
		t.Errorf("Case: Test error read file, expected %v got %v", models.ErrOpenFile, err)
	}
}

func TestBeefReadfileSuccess(t *testing.T) {
	bf := services.NewBeefService()
	err := bf.Read("./testFile/mock_file_beef.txt")
	if err != nil {
		t.Error(err)
	}

	got := bf.Get()
	expected := map[string]int{
		"bresaola": 1, "enim": 1, "fatback": 1, "jowl": 1, "meatloaf": 1, "pastrami": 1, "pork": 1, "t-bone": 4,
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Case: %v: expected %v, actual %v", "BeefReadfileSuccess", expected, got)
	}
}
