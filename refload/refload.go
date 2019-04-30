// Copyright 2019 The hjiangsse. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Represents NGST reference data structure using native Go struct types
package refload

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

//convert a string to another type
func convString(s string, t reflect.Type) (interface{}, error) {
	//create a new element, which type is v
	velm := reflect.New(t).Elem()

	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		if i, err := strconv.ParseInt(s, 10, 64); err != nil {
			return nil, err
		} else {
			velm.SetInt(i)
		}
		return velm.Interface(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64:
		if i, err := strconv.ParseUint(s, 10, 64); err != nil {
			return nil, err
		} else {
			velm.SetUint(i)
		}
		return velm.Interface(), nil
	case reflect.String:
		velm.SetString(s)
		return velm.Interface(), nil
	default:
		return nil, fmt.Errorf("%s can not convert to type %v", s, t.String())
	}
}

// Load reference data from file
func LoadRefDat(path string, sep string, s interface{}) error {
	//check if s is a pointer?
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("non-pointer %v\n", v.Type())
	}

	//check if s point to a slice
	ve := v.Elem()
	if ve.Kind() != reflect.Slice {
		return fmt.Errorf("pointed value non-slice %v\n", ve.Type())
	}

	//get the slice element type, and create a new element
	e := reflect.TypeOf(ve.Interface()).Elem()
	elem := reflect.New(e).Elem() //now elem can be setted
	numFields := elem.NumField()

	//read the file line by line
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	lineIdx := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//split current line with separator
		segs := strings.Split(strings.TrimSpace(scanner.Text()), sep)

		//line and slice type not match
		if len(segs) != numFields {
			return fmt.Errorf("[line:%d] line seg num %d != struct field num %d", lineIdx, len(segs), numFields)
		}

		for i := 0; i < numFields; i++ {
			v, err := convString(segs[i], elem.Field(i).Type())
			if err != nil {
				return err
			}
			elem.Field(i).Set(reflect.ValueOf(v))
		}
		//append new elem to the slice
		ve.Set(reflect.Append(ve, elem))
		lineIdx++
	}
	return nil
}
