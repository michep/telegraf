package graphite

import (
	"fmt"
	"regexp"
	"strings"
)

type GraphiteReParser struct {
	Separator   string
	DefaultTags map[string]string
	templates   []*regexp.Regexp
}

func NewGraphiteReParser(separator string, templates []string, defaultTags map[string]string) (*GraphiteReParser, error) {
	var (
		err error
		re  *regexp.Regexp
	)

	if separator == "" {
		separator = DefaultSeparator
	}
	p := &GraphiteReParser{
		Separator: separator,
	}

	if defaultTags != nil {
		p.DefaultTags = defaultTags
	}

	for _, template := range templates {
		re, err = regexp.Compile(template)
		if err == nil {
			p.templates = append(p.templates, re)
		} else {
			return p, fmt.Errorf("exec input parser config is error: %s ", err.Error())
		}
	}

	return p, nil
}

func (p *GraphiteReParser) ApplyTemplate(line string) (string, map[string]string, string, error) {
	var (
		found      bool
		matches    []string
		metricname string
		tags       map[string]string
	)
	// Break line into fields (name, value, timestamp), only name is used
	fields := strings.Fields(line)
	if len(fields) == 0 {
		return "", make(map[string]string), "", nil
	}
	name := fields[0]
	for _, regex := range p.templates {
		matches = regex.FindStringSubmatch(name)
		if len(matches) > 0 {
			metricname, tags = p.parseName(matches, regex.SubexpNames())
			found = true
			break
		}
	}
	if !found {
		return "", make(map[string]string), "", nil
	}
	return metricname, tags, "", nil
}

func (p *GraphiteReParser) parseName(matches, subnames []string) (name string, dynamic map[string]string) {
	dynamic = make(map[string]string)
	for idx, subname := range subnames {
		if subname != "" {
			if subname == "measurement" {
				name = p.concat(name, matches[idx], p.Separator)
			} else {
				dynamic[subname] = p.concat(dynamic[subname], matches[idx], p.Separator)
			}
		}
	}
	return
}

func (p *GraphiteReParser) concat(val1, val2, delim string) string {
	if val1 == "" {
		return val2
	} else {
		return val1 + delim + val2
	}
}
