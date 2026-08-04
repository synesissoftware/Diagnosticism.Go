package diagnosticism_test

import (
	d "github.com/synesissoftware/Diagnosticism.Go"

	"regexp"
	"testing"
)

// Optional Windows drive letter prefix (e.g. "D:"), then path segments that may
// not contain ':' until the first deliberate separator used by FileLine*.
const (
	reWinDrive = `(?:[A-Za-z]:)?`
)

func Test_File(t *testing.T) {

	expected := `^` + reWinDrive + `[^:]*debug_test[^:]*$`
	actual := d.File()

	match, _ := regexp.MatchString(expected, actual)
	if !match {

		t.Errorf("result of calling `File()` - '%s' - did not match expected format '%s'", actual, expected)
	}
}

func Test_FileLine(t *testing.T) {

	expected := `^` + reWinDrive + `[^:]*debug_test[^:]*:\d+$`
	actual := d.FileLine()

	match, _ := regexp.MatchString(expected, actual)
	if !match {

		t.Errorf("result of calling `FileLine()` - '%s' - did not match expected format '%s'", actual, expected)
	}
}

func Test_Function(t *testing.T) {

	expected := `^.*Test_Function$`
	actual := d.Function()

	match, _ := regexp.MatchString(expected, actual)
	if !match {

		t.Errorf("result of calling `Function()` - '%s' - did not match expected format '%s'", actual, expected)
	}
}

func Test_FileLineFunction(t *testing.T) {

	expected := `^` + reWinDrive + `[^:]*debug_test[^:]*:\d+:.*Test_FileLineFunction$`
	actual := d.FileLineFunction()

	match, _ := regexp.MatchString(expected, actual)
	if !match {

		t.Errorf("result of calling `FileLineFunction()` - '%s' - did not match expected format '%s'", actual, expected)
	}
}

func Test_Line(t *testing.T) {

	expected := 67
	actual := d.Line()

	if expected != actual {

		t.Errorf("result of calling `Line()` - %d - did not match expected value %d", actual, expected)
	}
}

func Test_LineFunction(t *testing.T) {

	expected := `^\d+:.*Test_LineFunction$`
	actual := d.LineFunction()

	match, _ := regexp.MatchString(expected, actual)
	if !match {

		t.Errorf("result of calling `LineFunction()` - '%s' - did not match expected format '%s'", actual, expected)
	}
}

func Test_GetFileLineFor(t *testing.T) {

	expected := `^` + reWinDrive + `[^:]*debug_test[^:]*:\d+$`
	actual, _ := d.GetFileLineFor(0)

	match, _ := regexp.MatchString(expected, actual)
	if !match {

		t.Errorf("result of calling `FileLine()` - '%s' - did not match expected format '%s'", actual, expected)
	}
}

func Test_GetFileLineFunctionFor(t *testing.T) {

	expected := `^` + reWinDrive + `[^:]*debug_test[^:]*:\d+:.*Test_GetFileLineFunctionFor$`
	actual, _ := d.GetFileLineFunctionFor(0)

	match, _ := regexp.MatchString(expected, actual)
	if !match {

		t.Errorf("result of calling `GetFileLineFunctionFor()` - '%s' - did not match expected format '%s'", actual, expected)
	}
}
