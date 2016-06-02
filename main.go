/* -.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.

* File Name : main.go

* Purpose :

* Creation Date : 05-28-2016

* Last Modified : Thu 02 Jun 2016 11:42:41 AM PDT

* Created By : Kiyor

_._._._._._._._._._._._._._._._._._._._._.*/

package logic

import (
	"bytes"
	"fmt"
	"log"
	"regexp"
	"text/template"
)

var (
	funcs = template.FuncMap{
		"re": reMatch,
	}
)

func reMatch(intf, r interface{}) (bool, error) {
	re, err := regexp.Compile(fmt.Sprintf("%v", r))
	if err != nil {
		return false, err
	}
	if re.MatchString(fmt.Sprintf("%v", intf)) {
		return true, nil
	}
	return false, nil
}

func Input(in string, intf interface{}) bool {
	return input(fmt.Sprintf(`{{if %s }}true{{else}}false{{end}}`, in), intf)
}

func input(in string, intf interface{}) bool {
	tmpl, err := template.New("main").Funcs(funcs).Parse(in)
	if err != nil {
		log.Fatal(err)
	}

	var b bytes.Buffer

	if err := tmpl.Execute(&b, intf); err != nil {
		log.Fatal(err)
	}

	if b.String() == "true" {
		return true
	}
	return false
}

func Verify(in string, intf interface{}) error {
	return verify(fmt.Sprintf(`{{if %s }}true{{else}}false{{end}}`, in), intf)
}

func verify(in string, intf interface{}) error {
	tmpl, err := template.New("main").Funcs(funcs).Parse(in)
	if err != nil {
		return err
	}

	var b bytes.Buffer

	if err := tmpl.Execute(&b, intf); err != nil {
		return err
	}
	return nil
}
