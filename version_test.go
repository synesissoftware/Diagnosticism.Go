package diagnosticism

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	Expected_VersionMajor uint16 = 0
	Expected_VersionMinor uint16 = 6
	Expected_VersionPatch uint16 = 1
	Expected_VersionAB    uint16 = 0xFFFF
)

func Test_Version_Elements(t *testing.T) {
	require.Equal(t, Expected_VersionMajor, VersionMajor)
	require.Equal(t, Expected_VersionMinor, VersionMinor)
	require.Equal(t, Expected_VersionPatch, VersionPatch)
	require.Equal(t, Expected_VersionAB, VersionAB)
}

func Test_Version(t *testing.T) {
	require.Equal(t, uint64(0x0000_0006_0001_FFFF), Version)
}

func Test_Version_String(t *testing.T) {
	require.Equal(t, "0.6.1", VersionString())
}
