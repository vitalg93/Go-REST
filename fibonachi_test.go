package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFibonachi(t *testing.T) {
	// test table:
	tt := []struct {
		name   string
		x      string
		y      string
		answer []int
		err    string
	}{
		{name: "case 1: x=0, y=5", x: "0", y: "5", answer: []int{0, 1, 1, 2, 3, 5}},
		{name: "case 2: x = y: x=5, y=5", x: "5", y: "5", answer: []int{5}},
		{name: "case 3: x > y: x=3, y=2", x: "3", y: "2", answer: []int{}},
		{name: "case 4: incorrect x value - x=a, y=5", x: "a", y: "5", err: "can't convert x into int"},
		{name: "case 5: incorrect y value - x=0, y=b", x: "0", y: "b", err: "can't convert y into int"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "localhost:8080/fibonachi?x="+tc.x+"&y="+tc.y, nil)
			if err != nil {
				t.Fatalf("could not created request: %v", err)
			}
			req.Header.Add("Content-Type", "application/json")

			rec := httptest.NewRecorder()
			getFibonachi(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}

			if tc.err != "" {
				if res.StatusCode != http.StatusBadRequest {
					t.Errorf("expected status Bad Request; got %v", res.StatusCode)
				}
				if msg := string(bytes.TrimSpace(b)); msg != tc.err {
					t.Errorf("expected message %q; got %q", tc.err, msg)
				}
				return
			}

			if res.StatusCode != http.StatusOK {
				t.Errorf("expected status OK; got %v", res.Status)
			}

			data := bytes.TrimSpace(b)
			var task Task
			err = json.Unmarshal(data, &task)
			if err != nil {
				t.Errorf("can't encode JSON; got %v", err)
			}

			if !equalSlices(task.Answer, tc.answer) {
				t.Fatalf("expected %v; got %v", tc.answer, task.Answer)
			}
		})
	}
}

// An auxiliary method for comparing slices. Returns true if the slices contain the same elements in the same order
func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
