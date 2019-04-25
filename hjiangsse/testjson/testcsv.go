package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const filePath = "./file.csv"
const separator = "|"

type lineElem struct {
	Num int32
	Age int
}

//convert a string to another type
func convString(s string, v reflect.Type) (interface{}, error) {
	switch v.Kind() {
	case reflect.Int:
		i, err := strconv.Atoi(s) //i is int
		if err != nil {
			return nil, err
		}
		return i, nil
	case reflect.Int8:
		i, err := strconv.ParseInt(s, 10, 8)
		if err != nil {
			return nil, err
		}
		return int8(i), nil
	case reflect.Int16:
		i, err := strconv.ParseInt(s, 10, 16)
		if err != nil {
			return nil, err
		}
		return int16(i), nil
	case reflect.Int32:
		i, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return nil, err
		}
		return int32(i), nil
	case reflect.Int64:
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}
		return i, nil

	case reflect.Uint:
		i, err := strconv.ParseUint(s, 10, 0) //i is uint
		if err != nil {
			return nil, err
		} else {
			return uint(i), nil
		}
	case reflect.Uint8:
		i, err := strconv.ParseUint(s, 10, 8)
		if err != nil {
			return nil, err
		}
		return uint8(i), nil
	case reflect.Uint16:
		i, err := strconv.ParseUint(s, 10, 16)
		if err != nil {
			return nil, err
		}
		return uint16(i), nil
	case reflect.Uint32:
		i, err := strconv.ParseUint(s, 10, 32)
		if err != nil {
			return nil, err
		}
		return uint32(i), nil
	case reflect.Uint64:
		i, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return nil, err
		}
		return i, nil
	case reflect.String:
		return s, nil
	default:
		return nil, fmt.Errorf("%s can not convert to type %T", s, v)
	}
	return nil, nil
}

func NewLoadLine(path string, sep string, s interface{}) error {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("non-pointer %v\n", v.Type())
	}

	ve := v.Elem()
	if ve.Kind() != reflect.Slice {
		return fmt.Errorf("pointed value non-slice %v\n", ve.Type())
	}

	//get the slice element type, and create a new element
	e := reflect.TypeOf(ve.Interface()).Elem()
	elem := reflect.New(e).Elem()
	numFields := elem.NumField()

	//read the file line by line
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	lineIdex := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		segs := strings.Split(line, sep)

		if len(segs) != numFields {
			return fmt.Errorf("[line:%d] seg num %d != field num %d", lineIdex, len(segs), numFields)
		}

		for i := 0; i < numFields; i++ {
			v, err := convString(segs[i], elem.Field(i).Type())
			if err != nil {
				return err
			}
			elem.Field(i).Set(reflect.ValueOf(v))
		}
		ve.Set(reflect.Append(ve, elem))
		lineIdex++
	}

	return nil
}

func main() {
	var pslice []lineElem
	err := NewLoadLine(filePath, separator, &pslice)
	if err != nil {
		panic(err)
	}
	fmt.Println(pslice)
}
