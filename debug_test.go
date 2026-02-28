package diagnosticism_test

import (
	d "github.com/synesissoftware/Diagnosticism.Go"

	"regexp"
	"testing"
)

func Test_File(t *testing.T) {

	// NOTE: this regex will likely not work on Windows

	expected := `^[^:]+debug_test[^:]+$`
	actual := d.File()

	match, _ := regexp.MatchString(expected, actual)
	if !match {

		t.Errorf("result of calling `File()` - '%s' - did not match expected format '%s'", actual, expected)
	}
}

func Test_FileLine(t *testing.T) {

	// NOTE: this regex will likely not work on Windows

	expected := `^[^:]+debug_test[^:]+:\d+$`
	actual := d.FileLine()

	match, _ := regexp.MatchString(expected, actual)
	if !match {

		t.Errorf("result of calling `FileLine()` - '%s' - did not match expected format '%s'", actual, expected)
	}
}

func Test_Function(t *testing.T) {

	// NOTE: this regex will likely not work on Windows

	expected := `^.*Test_Function$`
	actual := d.Function()

	match, _ := regexp.MatchString(expected, actual)
	if !match {

		t.Errorf("result of calling `Function()` - '%s' - did not match expected format '%s'", actual, expected)
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

func Test_Line(t *testing.T) {

	// NOTE: this regex will likely not work on Windows

	expected := 71
	actual := d.Line()

	if expected != actual {

		t.Errorf("result of calling `Line()` - %d - did not match expected value %d", actual, expected)
	}
}

func Test_LineFunction(t *testing.T) {

	// NOTE: this regex will likely not work on Windows

	expected := `^\d+:.*Test_LineFunction$`
	actual := d.LineFunction()

	match, _ := regexp.MatchString(expected, actual)
	if !match {

		t.Errorf("result of calling `LineFunction()` - '%s' - did not match expected format '%s'", actual, expected)
	}
}

func Test_GetFileLineFor(t *testing.T) {

	// NOTE: this regex will likely not work on Windows

	expected := `^[^:]+debug_test[^:]+:\d+$`
	actual, _ := d.GetFileLineFor(0)

	match, _ := regexp.MatchString(expected, actual)
	if !match {

		t.Errorf("result of calling `FileLine()` - '%s' - did not match expected format '%s'", actual, expected)
	}
}

func Test_GetFileLineFunctionFor(t *testing.T) {

	expected := `^[^:]+debug_test[^:]+:\d+:.*Test_GetFileLineFunctionFor$`
	actual, _ := d.GetFileLineFunctionFor(0)

	match, _ := regexp.MatchString(expected, actual)
	if !match {

		t.Errorf("result of calling `GetFileLineFunctionFor()` - '%s' - did not match expected format '%s'", actual, expected)
	}
}
