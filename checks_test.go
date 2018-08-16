/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   checks_test.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: flourme <flourme@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/08/16 01:56:21 by flourme           #+#    #+#             */
/*   Updated: 2018/08/16 01:56:21 by flourme          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestCheckHeader(t *testing.T) {

	tests := []struct {
		path        string
		description string
		shouldFail  bool
		shouldSkip  bool
	}{
		{"tests/header-tests/good-one", "Testing a good header [should pass]", false, false},
		{"tests/header-tests/missplaced", "Testing a missplaced header [should fail]", true, false},
		{"tests/header-tests/miss-login", "Testing a bad login header [should fail]", true, false},
		{"tests/header-tests/miss-login", "Testing a bad login header but skip CheckHeader [should pass]", false, true},
	}

	for _, test := range tests {
		t.Log(test.description)
		if test.shouldSkip {
			Cfg.Options.BypassHeader = true
		}

		content, _ := ioutil.ReadFile(test.path)
		s := strings.Split(string(content), "\n")
		normErrors, err := CheckHeader(s)
		if test.shouldFail {
			if err == nil || len(normErrors) == 0 {
				t.Error("This test should fail, here's what happen :", err, normErrors)
			}
		} else {
			if err != nil || len(normErrors) != 0 {
				t.Error("This test shouldn't fail, here's what happen : ", err, normErrors)
			}
		}
	}
}
