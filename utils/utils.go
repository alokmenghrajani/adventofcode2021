package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

const MaxInt = int(^uint(0) >> 1)
const MaxUint = ^uint(0)
const MinInt = ^MaxInt

func IntMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func SliceMinMax(slice []int) (int, int) {
	min := MaxInt
	max := MinInt
	for _, v := range slice {
		min = IntMin(min, v)
		max = IntMax(max, v)
	}
	return min, max
}

// Like strconv.Atoi but returns a default value on error
func Atoi(s string, fallback int) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		return fallback
	}
	return v
}

func MustAtoi(s string) int {
	v, err := strconv.Atoi(s)
	PanicOnErr(err)
	return v
}

// Returns key from map[T]int which has the max value
func MapFindMax(m interface{}) interface{} {
	var maxK interface{} = nil
	var maxV = MinInt
	iter := reflect.ValueOf(m).MapRange()
	for iter.Next() {
		k := iter.Key()
		v := int(iter.Value().Int())
		if v > maxV {
			maxV = v
			maxK = k.Interface()
		}
	}
	return maxK
}

// Returns key from map[T]int which has the min value
func MapFindMin(m interface{}) interface{} {
	var minK interface{} = nil
	var minV = MaxInt
	iter := reflect.ValueOf(m).MapRange()
	for iter.Next() {
		k := iter.Key()
		v := int(iter.Value().Int())
		if v < minV {
			minV = v
			minK = k.Interface()
		}
	}
	return minK
}

func Readfile(year, day int) string {
	filename := ""
	if year == 2021 {
		filename = fmt.Sprintf("day%02d/input.txt", day)
	} else {
		filename = fmt.Sprintf("%d/year%dday%02d/input.txt", year, year, day)
	}
	file, err := os.Open(filename)
	PanicOnErr(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	contents, err := ioutil.ReadAll(reader)
	PanicOnErr(err)

	return strings.TrimSuffix(string(contents), "\n")
}

func ParseToStruct(re *regexp.Regexp, input string, target interface{}) bool {
	m := re.FindStringSubmatch(input)
	if m == nil {
		return false
	}

	var useOffset bool

	for i, name := range re.SubexpNames() {
		if i == 0 {
			continue
		}
		var field reflect.Value
		if name == "" {
			// use offset
			if i == 1 {
				useOffset = true
			} else if !useOffset {
				panic("can't mix named and unnamed subexpressions")
			}
			field = reflect.ValueOf(target).Elem().Field(i - 1)
		} else {
			// use name
			if i == 1 {
				useOffset = false
			} else if useOffset {
				panic("can't mix named and unnamed subexpressions")
			}
			field = reflect.ValueOf(target).Elem().FieldByName(name)
		}
		if field.Kind() == reflect.String {
			field.SetString(m[i])
		} else if field.Kind() == reflect.Int {
			v, err := strconv.Atoi(m[i])
			PanicOnErr(err)
			field.SetInt(int64(v))
		} else if field.Kind() == reflect.Uint8 {
			if len(m[i]) != 1 {
				panic(fmt.Sprintf("expecting 1 char, got: %s", m[i]))
			}
			field.SetUint(uint64(m[i][0]))
		} else {
			panic(fmt.Sprintf("unknown kind: %s", field.Kind()))
		}
	}
	return true
}

func MustParseToStruct(re *regexp.Regexp, input string, target interface{}) {
	if !ParseToStruct(re, input, target) {
		panic(fmt.Errorf("failed to parse: %s", input))
	}
}

func CharToLower(c byte) byte {
	return strings.ToLower(string(c))[0]
}

func CharToUpper(c byte) byte {
	return strings.ToUpper(string(c))[0]
}

func Contains(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Gcd(x, y int) int {
	if x <= 0 || y <= 0 {
		panic(fmt.Errorf("invalid input: %d, %d", x, y))
	}
	if x == y {
		return x
	}
	if x > y {
		return Gcd(x-y, y)
	} else {
		return Gcd(x, y-x)
	}
}

func Sign(x int) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	} else {
		return 0
	}
}
