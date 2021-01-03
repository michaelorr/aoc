package days

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"orr.co/adventofcode/data"
)

func Day4() {
	Day4Part1()
}

func Day4Part1() {
	var valid int
	for line := range data.AsChunks(data.Day(4)) {
		p := NewPassport(line)

		if p.IsValid() {
			p.Print()
			valid++
		}

	}
	fmt.Println(valid)
}

func (p *Passport) Print() {
	fmt.Printf("\tpid: %s\n", *p.pid)
}

type Passport struct {
	byr *string
	iyr *string
	eyr *string
	hgt *string
	hcl *string
	ecl *string
	pid *string
	cid *string
}

func (p *Passport) IsValid() bool {
	var s string
	var i int

	// BYR
	if p.byr == nil {
		return false
	}
	s = string(*p.byr)
	i, _ = strconv.Atoi(s)
	if !regexp.MustCompile("[[:digit:]]{4}").MatchString(s) || i < 1920 || i > 2002 {
		return false
	}

	//iyr
	if p.iyr == nil {
		return false
	}
	s = string(*p.iyr)
	i, _ = strconv.Atoi(s)
	if !regexp.MustCompile("[[:digit:]]{4}").MatchString(s) || i < 2010 || i > 2020 {
		return false
	}

	//eyr
	if p.eyr == nil {
		return false
	}
	s = string(*p.eyr)
	i, _ = strconv.Atoi(s)
	if !regexp.MustCompile("[[:digit:]]{4}").MatchString(s) || i < 2020 || i > 2030 {
		return false
	}

	//hgt
	if p.hgt == nil {
		return false
	}
	s = string(*p.hgt)
	cm_regex := regexp.MustCompile("([[:digit:]]*)cm")
	in_regex := regexp.MustCompile("([[:digit:]]*)in")
	cm := cm_regex.MatchString(s)
	in := in_regex.MatchString(s)

	if !(cm || in) {
		return false
	}

	if cm {
		num, _ := strconv.Atoi(cm_regex.FindStringSubmatch(s)[1])
		if num < 150 || num > 193 {
			return false
		}
	}
	if in {
		num, _ := strconv.Atoi(in_regex.FindStringSubmatch(s)[1])
		if num < 59 || num > 76 {
			return false
		}
	}

	//hcl
	if p.hcl == nil {
		return false
	}
	s = *p.hcl
	if !regexp.MustCompile("#[[:xdigit:]]{6}").MatchString(s) {
		return false
	}

	//ecl
	ecl := make(map[string]bool)
	ecl["amb"] = true
	ecl["blu"] = true
	ecl["brn"] = true
	ecl["gry"] = true
	ecl["grn"] = true
	ecl["hzl"] = true
	ecl["oth"] = true

	if p.ecl == nil {
		return false
	}
	if _, ok := ecl[*p.ecl]; !ok {
		return false
	}
	//pid
	if p.pid == nil || len(*p.pid) != 9 {
		return false
	}

	if !regexp.MustCompile("[[:digit:]]{9}").MatchString(*p.pid) {
		return false
	}

	return true
}

func NewPassport(input string) *Passport {
	p := &Passport{}
	for _, field := range strings.Fields(input) {
		kv := strings.Split(field, ":")
		switch key, val := kv[0], kv[1]; key {
		case "byr":
			p.byr = &val
		case "iyr":
			p.iyr = &val
		case "eyr":
			p.eyr = &val
		case "hgt":
			p.hgt = &val
		case "hcl":
			p.hcl = &val
		case "ecl":
			p.ecl = &val
		case "pid":
			p.pid = &val
		case "cid":
			p.cid = &val
		}
	}

	return p
}
