/* -.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.

* File Name : main_test.go

* Purpose :

* Creation Date : 05-28-2016

* Last Modified : Thu 02 Jun 2016 11:42:27 AM PDT

* Created By : Kiyor

_._._._._._._._._._._._._._._._._._._._._.*/

package logic

import (
	"testing"
)

func Test_input(t *testing.T) {
	m1 := make(map[string]bool)
	m1[`{{ if eq "a" "a" }}true{{else}}false{{end}}`] = true
	m1[`{{ if eq "a" "b" }}true{{else}}false{{end}}`] = false

	for k, v := range m1 {
		if input(k, nil) != v {
			t.Fatal(k, v)
		}
	}

	m2 := make(map[string]bool)
	m2[`eq "a" "a"`] = true
	m2[`eq "a" "b"`] = false
	m2[`gt 1 0 `] = true
	m2[`and (or (gt 1 0) (gt 1 2)) (eq 1 1)`] = true
	m2[`true`] = true
	m2[`false`] = false

	for k, v := range m2 {
		if Input(k, nil) != v {
			t.Fatal(k, v, Input(k, nil))
		}
	}

	m3 := make(map[string]bool)
	t3 := make(map[string]interface{})
	t3["a"] = 1
	t3["b"] = "hello"
	m3[`and (or (gt .a 0) (gt .a 2)) (eq .a 1)`] = true
	m3[`re .b ".ello"`] = true
	m3[`not (re .b ".ello")`] = false
	for k, v := range m3 {
		if Input(k, t3) != v {
			t.Fatal(k, v, Input(k, t3))
		}
	}

	m4 := make(map[string]bool)
	type T4 struct {
		A int
	}
	t4 := &T4{A: 1}
	m4[`and (or (gt .A 0) (gt .A 2)) (eq .A 1)`] = true
	for k, v := range m4 {
		if Input(k, t4) != v {
			t.Fatal(k, v, Input(k, t4))
		}
	}

}
