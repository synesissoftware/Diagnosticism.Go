package diagnosticism_test

import (
	d "github.com/synesissoftware/Diagnosticism.Go"

	"regexp"
	"testing"
)

func Test_FileLine(t *testing.T) {

	// NOTE: this regex will likely not work on Windows

	expected := `^[^:]+debug_test[^:]+:\d+$`
	actual := d.FileLine()

	match, _ := regexp.MatchString(expected, actual)
	if !match {

		t.Errorf("result of calling `FileLine()` - '%s' - did not match expected format '%s'", actual, expected)
	}
}

func Test_FileLineFunction(t *testing.T) {

	// NOTE: this regex will likely not work on Windows

	expected := `^[^:]+debug_test[^:]+:\d+:.*Test_FileLineFunction$`
	actual := d.FileLineFunction()

	match, _ := regexp.MatchString(expected, actual)
	if !match {

		t.Errorf("result of calling `FileLineFunction()` - '%s' - did not match expected format '%s'", actual, expected)
	}
}
