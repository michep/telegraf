package graphite

import (
	"fmt"
	"regexp"
	"strings"
)

type replacer struct {
	template *regexp.Regexp
	format   string
}

type GraphiteReParser struct {
	Separator            string
	DefaultTags          map[string]string
	MeasurementGroupName string
	templates            []replacer
}

func NewGraphiteReParser(separator, measurementgroupname string, templates []string, defaultTags map[string]string) (*GraphiteReParser, error) {
	var (
		err error
		re  *regexp.Regexp
	)

	if separator == "" {
		separator = DefaultSeparator
	}
	if measurementgroupname == "" {
		measurementgroupname = "metric"
	}
	p := &GraphiteReParser{
		Separator:            separator,
		MeasurementGroupName: measurementgroupname,
	}

	if defaultTags != nil {
		p.DefaultTags = defaultTags
	}

	for _, template := range templates {
		parts := strings.Split(template, " => ")
		if len(parts) != 2 {
			return p, fmt.Errorf("exec input parser config is error: no => in template")
		}
		re, err = regexp.Compile(parts[0])
		if err == nil {
			repl := replacer{template: re, format: parts[1]}
			p.templates = append(p.templates, repl)
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
	for _, repl := range p.templates {
		matches = repl.template.FindStringSubmatch(name)
		if len(matches) > 0 {
			tags = p.parseTags(matches, repl.template)
			metricname = repl.template.ReplaceAllString(name, repl.format)
			found = true
			break
		}
	}
	if !found {
		return "", make(map[string]string), "", nil
	}
	return metricname, tags, "", nil
}

func (p *GraphiteReParser) parseTags(matches []string, regexp *regexp.Regexp) (dynamic map[string]string) {
	dynamic = make(map[string]string)
	for idx, subname := range regexp.SubexpNames() {
		if subname != "" && !strings.HasPrefix(subname, p.MeasurementGroupName) {
			dynamic[subname] = p.concat(dynamic[subname], matches[idx], p.Separator)
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
