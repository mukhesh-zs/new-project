package main

import (
	"log"
	"testing"
)

func TestPostData(t *testing.T) {
	testcases := []struct {
		desc      string
		input     employee_info
		expoutput employee_info
	}{
		{"mukhesh details", employee_info{2, "mukhesh", "SDE", 0, 10000}, employee_info{2, "mukhesh", "SDE", 0, 10000}},
	}
	for i, v := range testcases {
		emp, err := PostData(v.input)
		if err != nil {
			log.Print(err)
		}
		if emp.id != v.expoutput.id {
			t.Error("test failed")
			t.Errorf("%d", i)
		}
	}

}
func TestGetData(t *testing.T) {
	testcases := []struct {
		desc      string
		input     int64
		expoutput employee_info
	}{
		{desc: "mukhesh info", input: 1, expoutput: employee_info{1, "mukhesh", "SDE", 0, 10000}},
	}
	for _, val := range testcases {
		output, err := GetData(val.input)
		if err != nil {
			log.Print(err)
		}
		if output != val.expoutput {
			t.Error("test failed")
		}
	}
}
