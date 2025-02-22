package diagnosticism_test

import (
	d "github.com/synesissoftware/Diagnosticism.Go"
	stegol "github.com/synesissoftware/STEGoL"

	"strings"
	"testing"
)

func Test_Version(t *testing.T) {

	var splits = strings.Split(d.VersionString, ".")
	var jnp = splits[0:3]

	var v_jnp = strings.Join(jnp, ".")

	stegol.CheckStringEqual(t, "0.6.0", v_jnp)
}
