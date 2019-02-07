package filesystem

import (
	"testing"
)

// TestModePermissionMaskIsExpected is a sanity check that ModePermissionsMask
// is equivalent to 0777 on all platforms (which it should be on POSIX platforms
// under the POSIX standard and on Windows platforms based on the os package's
// (immutable) FileMode definition).
func TestModePermissionMaskIsExpected(t *testing.T) {
	if ModePermissionsMask != Mode(0777) {
		t.Error("ModePermissionsMask value not equal to expected:", ModePermissionsMask, "!=", Mode(0777))
	}
}