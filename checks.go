package main

import (
	"fmt"
	"regexp"
)

type NormError struct {
	line   int
	column int
	msg    string
}

// headerRegExp contains all the regular Expressions for CheckHeader
// Each RegExp should check each line of the header
var headerRegExp = []string{
	`\/\* \*{74} \*\/`,
	`\/\* {76}\*\/`,
	`\/\* {56}:{3} {6}:{8} {3}\*\/`,
	`\/\* {3}\S* *:\+: {6}:\+: {4}:\+: {3}\*\/`,
	`\/\* {52}\+:\+ \+:\+ {9}\+:\+ {5}\*\/`,
	`\/\* {3}By: \S{1,8} <\S*@\S*\.\S*> *\+\#\+  \+:\+ {7}\+#\+ {8}\*\/`,
	`\/\* {48}\+#\+#\+#\+#\+#\+   \+#\+ {11}\*\/`,
	`\/\* {3}Created: \d{4}\/\d{2}\/\d{2} \d{2}:\d{2}:\d{2} by \S{1,8} *#\+# {4}#\+# {13}\*\/`,
	`\/\* {3}Updated: \d{4}\/\d{2}\/\d{2} \d{2}:\d{2}:\d{2} by \S{1,8} *###   #{8}.fr {7}\*\/`,
	`\/\* {76}\*\/`,
	`\/\* \*{74} \*\/`,
}

// CheckHeader checks if the 42 Header is correctly on top of files.
// It checks also if all the lines have exactly 80 characters, because
// some items (like logins) have variable lenght.
// Can be skipped with --bypass-header flag
//
// Checks:
//	* Minimum height in a file for header
//	* Line width for every header's line
//	* Every line pattern
//
// Returns:
// 	- []NormError : All the NormError spotted
//	- error: The last error spotted
func CheckHeader(file []string) ([]NormError, error) {
	var headerSlice []string

	var r *regexp.Regexp
	var err error
	normErrors := []NormError{}

	if Cfg.Options.BypassHeader {
		return nil, nil
	}

	if len(file) >= 11 {
		headerSlice = file[:11]
	} else {
		normErrors = append(normErrors, NormError{
			line:   len(file),
			column: 1,
			msg:    `File has not the minimum line height for header, is it really here ?`,
		})
		err = fmt.Errorf(`File has not the minimum line height for header, is it really here ?`)
	}

	for index, line := range headerSlice {
		r, _ = regexp.Compile(headerRegExp[index])
		if r.MatchString(line) != true {
			normErrors = append(normErrors, NormError{
				line:   index + 1,
				column: 1,
				msg:    `Header pattern mismatching`,
			})
			err = fmt.Errorf(`Header pattern mismatching`)
		}

		if len(line) != 80 {
			normErrors = append(normErrors, NormError{
				line:   index + 1,
				column: len(line),
				msg:    "Line width exceed",
			})
			err = fmt.Errorf("Line width exceed")
		}
	}
	return normErrors, err
}

/*                                                                              */
