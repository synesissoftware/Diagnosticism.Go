package severity_test

import (
	. "github.com/synesissoftware/Diagnosticism.Go/severity"

	"github.com/stretchr/testify/require"

	"testing"
)

func Test_Severity_LEVELS_EXIST(t *testing.T) {

	require.Equal(t, 0, int(Unspecified))
	require.Equal(t, 1, int(Violation))
	require.Equal(t, 2, int(Alert))
	require.Equal(t, 3, int(Critical))
	require.Equal(t, 4, int(Failure))
	require.Equal(t, 5, int(Warning))
	require.Equal(t, 6, int(Notice))
	require.Equal(t, 7, int(Informational))
	require.Equal(t, 8, int(Debug0))
	require.Equal(t, 9, int(Debug1))
	require.Equal(t, 10, int(Debug2))
	require.Equal(t, 11, int(Debug3))
	require.Equal(t, 12, int(Debug4))
	require.Equal(t, 13, int(Debug5))
	require.Equal(t, 14, int(Trace))
}

func Test_Severity_String(t *testing.T) {

	require.Equal(t, "Violation", Violation.String())
	require.Equal(t, "Alert", Alert.String())
	require.Equal(t, "Critical", Critical.String())
	require.Equal(t, "Failure", Failure.String())
	require.Equal(t, "Warning", Warning.String())
	require.Equal(t, "Notice", Notice.String())
	require.Equal(t, "Informational", Informational.String())
	require.Equal(t, "Debug0", Debug0.String())
	require.Equal(t, "Debug1", Debug1.String())
	require.Equal(t, "Debug2", Debug2.String())
	require.Equal(t, "Debug3", Debug3.String())
	require.Equal(t, "Debug4", Debug4.String())
	require.Equal(t, "Debug5", Debug5.String())
	require.Equal(t, "Trace", Trace.String())

	require.Equal(t, "<Severity: 999>", Severity(999).String())
}

func Test_Severity_TranslateStockSeverity(t *testing.T) {

	require.Equal(t, "Violation", TranslateStockSeverity(Violation))
	require.Equal(t, "Alert", TranslateStockSeverity(Alert))
	require.Equal(t, "Critical", TranslateStockSeverity(Critical))
	require.Equal(t, "Failure", TranslateStockSeverity(Failure))
	require.Equal(t, "Warning", TranslateStockSeverity(Warning))
	require.Equal(t, "Notice", TranslateStockSeverity(Notice))
	require.Equal(t, "Informational", TranslateStockSeverity(Informational))
	require.Equal(t, "Debug0", TranslateStockSeverity(Debug0))
	require.Equal(t, "Debug1", TranslateStockSeverity(Debug1))
	require.Equal(t, "Debug2", TranslateStockSeverity(Debug2))
	require.Equal(t, "Debug3", TranslateStockSeverity(Debug3))
	require.Equal(t, "Debug4", TranslateStockSeverity(Debug4))
	require.Equal(t, "Debug5", TranslateStockSeverity(Debug5))
	require.Equal(t, "Trace", TranslateStockSeverity(Trace))

	require.Equal(t, "<Severity: 999>", TranslateStockSeverity(Severity(999)))
}

func Test_Severity_ColouredSeverityToString(t *testing.T) {

	require.Equal(t, "\033[1;93;41;5mViolation\033[0m", ColouredSeverityToString(Violation))
	require.Equal(t, "\033[1;96;41;5mAlert\033[0m", ColouredSeverityToString(Alert))
	require.Equal(t, "\033[1;97;41mCritical\033[0m", ColouredSeverityToString(Critical))
	require.Equal(t, "\033[1;31;43mFailure\033[0m", ColouredSeverityToString(Failure))
	require.Equal(t, "\033[1;34;43mWarning\033[0m", ColouredSeverityToString(Warning))
	require.Equal(t, "\033[1;97;100mNotice\033[0m", ColouredSeverityToString(Notice))
	require.Equal(t, "\033[1;37;100mInformational\033[0m", ColouredSeverityToString(Informational))
	require.Equal(t, "\033[1;37;44mDebug0\033[0m", ColouredSeverityToString(Debug0))
	require.Equal(t, "\033[1;37;44mDebug1\033[0m", ColouredSeverityToString(Debug1))
	require.Equal(t, "\033[1;37;44mDebug2\033[0m", ColouredSeverityToString(Debug2))
	require.Equal(t, "\033[1;37;44mDebug3\033[0m", ColouredSeverityToString(Debug3))
	require.Equal(t, "\033[1;37;44mDebug4\033[0m", ColouredSeverityToString(Debug4))
	require.Equal(t, "\033[1;37;44mDebug5\033[0m", ColouredSeverityToString(Debug5))
	require.Equal(t, "\033[0;37;44mTrace\033[0m", ColouredSeverityToString(Trace))

	require.Equal(t, "\033[1;31;47m<Severity: 999>\033[0m", ColouredSeverityToString(Severity(999)))
}
