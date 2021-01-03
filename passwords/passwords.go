package passwords

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

type PassPolicy struct {
	LowerBound int
	UpperBound int
	Letter     string
	Password   string
}

func ParsePolicy(line string) *PassPolicy {
	colonIndex := strings.Index(line, ":")
	policyStr, pass := line[:colonIndex], line[colonIndex+2:]

	pattern := regexp.MustCompile("([[:digit:]]+)-([[:digit:]]+) ([[:alpha:]])")
	matches := pattern.FindStringSubmatch(policyStr)

	var lower, upper int
	var err error
	if lower, err = strconv.Atoi(matches[1]); err != nil {
		log.Fatal(err)
	}
	if upper, err = strconv.Atoi(matches[2]); err != nil {
		log.Fatal(err)
	}
	p := &PassPolicy{LowerBound: lower, UpperBound: upper, Letter: matches[3], Password: pass}

	return p
}

func (p *PassPolicy) OldValid() bool {
	counter := make(map[string]int)
	for _, ch := range p.Password {
		counter[string(ch)]++
	}
	countInPassword := counter[p.Letter]
	return countInPassword >= p.LowerBound && countInPassword <= p.UpperBound
}

func (p *PassPolicy) NewValid() bool {
	keyInFirstPos := string(p.Password[p.LowerBound-1]) == p.Letter
	keyInSecondPos := string(p.Password[p.UpperBound-1]) == p.Letter

	return (keyInFirstPos || keyInSecondPos) && !(keyInFirstPos && keyInSecondPos)
}
