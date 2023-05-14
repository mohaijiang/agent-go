package idl

import (
	"reflect"
	"strings"
)

// lowerFirstCharacter returns a copy of the string with the first character in lower case.
// e.g. "UserName" -> "userName"
func lowerFirstCharacter(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}

type tag struct {
	// Name is the name of the field in the struct.
	name        string
	omitEmpty   bool
	variantType bool
}

func parseTags(field reflect.StructField) tag {
	icTag := field.Tag.Get("ic")
	if icTag != "" {
		var t tag
		tags := strings.Split(icTag, ",")
		if len(tags) != 0 {
			t.name = tags[0]
			for _, option := range tags[1:] {
				switch option {
				case "omitempty":
					t.omitEmpty = true
				case "variant":
					t.variantType = true
				default:
					// ignore unknown options
				}
			}
		}
		return t
	}
	return tag{
		name: lowerFirstCharacter(field.Name),
	}
}
