package main

import (
	"regexp"
	"strings"
)

type funcSignature struct {
	VarName  string
	TypeName string
}

type rule struct {
	Type  string
	Value string
}

type validationRule struct {
	FieldName string
	FieldType string
	Rules     []rule
}

type templateData struct {
	Signature       funcSignature
	ValidationRules []validationRule
}

func buildValidationRules(fields []structField) (validationRules []validationRule) {
	if fields == nil {
		return nil
	}
	validationRules = make([]validationRule, 0)
	for _, field := range fields {
		validationRules = append(validationRules, validationRule{
			FieldName: field.Name,
			FieldType: field.Type,
			Rules:     parseRules(field.Tag),
		})
	}
	return validationRules
}

func parseRules(tag string) (rules []rule) {
	rules = make([]rule, 0)

	if tag == "" {
		return rules
	}

	vldPosition := strings.LastIndex(tag, "validate")
	if vldPosition == -1 {
		return rules
	}

	re, err := regexp.Compile(`validate:"(.*?)"`)
	if err != nil {
		return rules
	}
	match := re.FindStringSubmatch(tag)

	if match == nil {
		return rules
	}

	ruleStrs := strings.Split(match[1], "|")
	for _, str := range ruleStrs {
		colonPos := strings.Index(str, ":")
		if colonPos != -1 {
			rType := str[:colonPos]
			rValue := str[colonPos+1:]
			if rType != "" && rValue != "" {
				rules = append(rules, rule{
					Type:  rType,
					Value: rValue,
				})
			}
		}
	}

	return rules
}
