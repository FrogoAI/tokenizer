package tokenizer

import (
	"testing"

	"github.com/FrogoAI/testutils"
)

func TestSplitBetweenTokens(t *testing.T) {
	testcases := []struct {
		name      string
		data      string
		arguments []string
		result    []string
	}{
		{
			name:      "split_between_different_tokens",
			data:      `some_string_which_we_should_split@should_not_be_visible;must_be_present`,
			arguments: []string{"@", ";"},
			result:    []string{"some_string_which_we_should_split", "must_be_present"},
		},
		{
			name:      "split_between_same_token_tokens",
			data:      `some_string_which_we_should_split;should_not_be_visible;must_be_present`,
			arguments: []string{";", ";"},
			result:    []string{"some_string_which_we_should_split", "must_be_present"},
		},
		{
			name:      "split_between_single_token",
			data:      `some_string_which_we_should_split;should_not_be_visible;must_be_present`,
			arguments: []string{";"},
			result:    []string{"some_string_which_we_should_split", "must_be_present"},
		},
		{
			name:      "return_fist_part_for_single_token",
			data:      `some_string_which_we_should_split;should_not_be_visible`,
			arguments: []string{";"},
			result:    []string{"some_string_which_we_should_split"},
		},
		{
			name:      "return_income_string_if_no_arguments",
			data:      `some_string_which_we_should_split;should_be_also_visible`,
			arguments: []string{},
			result:    []string{"some_string_which_we_should_split;should_be_also_visible"},
		},
		{
			name:      "return_income_string_if_no_match",
			data:      `some_string_which_we_should_split;should_be_also_visible`,
			arguments: []string{"@"},
			result:    []string{"some_string_which_we_should_split;should_be_also_visible"},
		},
		{
			name:      "return_empty_for_empty_input",
			data:      ``,
			arguments: []string{"@"},
			result:    []string{},
		},
		{
			name:      "if_both_token_are_empty",
			data:      `some_string_which_we_should_split`,
			arguments: []string{"", ""},
			result:    []string{"some_string_which_we_should_split"},
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			result := SplitBetweenTokens(test.data, test.arguments...)
			testutils.Equal(t, result, test.result)
		})
	}
}

func TestSanitizeEmail(t *testing.T) {
	testcases := []struct {
		name   string
		email  string
		result string
	}{
		{
			name:   "email_with_tag",
			email:  "testemail+example@gmail.com",
			result: `testemail@gmail.com`,
		},
		{
			name:   "email_with_two_tags",
			email:  "testemail+exa+mple@gmail.com",
			result: `testemail@gmail.com`,
		},
		{
			name:   "user_name_with_tag_without_domain",
			email:  "testemail+exa",
			result: `testemail`,
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			result := SanitizeEmail(test.email)
			testutils.Equal(t, result, test.result)
		})
	}
}

func TestByteSliceToString(t *testing.T) {
	testcases := map[string]struct {
		bytes []byte
		out   string
	}{
		"inStrs:empty":     {bytes: []byte{}, out: ""},
		"inStrs:nil":       {bytes: nil, out: ""},
		"inStrs:non_empty": {bytes: []byte{72, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 33}, out: "Hello world!"},
	}

	for k, c := range testcases {
		c := c

		t.Run(k, func(t *testing.T) {
			str := ByteSliceToString(c.bytes)
			testutils.Equal(t, str, c.out)
		})
	}
}

func TestByteSliceToStringNative(t *testing.T) {
	testcases := map[string]struct {
		bytes []byte
		out   string
	}{
		"inStrs:empty":     {bytes: []byte{}, out: ""},
		"inStrs:nil":       {bytes: nil, out: ""},
		"inStrs:non_empty": {bytes: []byte{72, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 33}, out: "Hello world!"},
	}

	for k, c := range testcases {
		c := c

		t.Run(k, func(t *testing.T) {
			str := ByteSliceToString(c.bytes)
			testutils.Equal(t, str, c.out)
		})
	}
}

func TestBetween(t *testing.T) {
	testcases := map[string]struct {
		data   string
		keys   []string
		result string
	}{
		"empty_data": {
			data:   "",
			keys:   []string{"[RESULT]"},
			result: "",
		},
		"not_key_in_data": {
			data:   "Some text",
			keys:   []string{"[RESULT]"},
			result: "",
		},
		"single_key_in_data": {
			data:   "Some [RESULT]text",
			keys:   []string{"[RESULT]"},
			result: "",
		},
		"key_is_present": {
			data:   "Some [RESULT]text[RESULT]",
			keys:   []string{"[RESULT]"},
			result: "text",
		},
		"trip_space_in_result": {
			data:   "Some [RESULT] text \n[RESULT]",
			keys:   []string{"[RESULT]"},
			result: "text",
		},
		"no_key": {
			data:   "Some [RESULT] text \n[RESULT]",
			keys:   []string{},
			result: "",
		},
		"between_two_keys": {
			data:   "Some < text\n > \n",
			keys:   []string{"<", ">"},
			result: "text",
		},
	}

	for name, test := range testcases {
		test := test

		t.Run(name, func(t *testing.T) {
			result := Between(test.data, test.keys...)
			testutils.Equal(t, test.result, result)
		})
	}
}
