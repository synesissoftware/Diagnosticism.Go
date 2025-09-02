package diagnosticism_test

import (
	. "github.com/synesissoftware/Diagnosticism.Go"

	"github.com/stretchr/testify/require"

	"testing"
)

const (
	Expected_VersionMajor uint16 = 0
	Expected_VersionMinor uint16 = 10
	Expected_VersionPatch uint16 = 0
	Expected_VersionAB    uint16 = 0xFFFF
)

func Test_Version_Elements(t *testing.T) {
	require.Equal(t, Expected_VersionMajor, VersionMajor)
	require.Equal(t, Expected_VersionMinor, VersionMinor)
	require.Equal(t, Expected_VersionPatch, VersionPatch)
	require.Equal(t, Expected_VersionAB, VersionAB)
}

func Test_Version(t *testing.T) {
	require.Equal(t, uint64(0x0000_000A_0000_FFFF), Version)
}

func Test_Version_String(t *testing.T) {
	require.Equal(t, "0.10.0", VersionString())
}
