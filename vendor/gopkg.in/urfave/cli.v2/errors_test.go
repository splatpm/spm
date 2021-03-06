package cli

import (
	"errors"
	"os"
	"testing"
)

func TestHandleExitCoder_nil(t *testing.T) {
	exitCode := 0
	called := false

	OsExiter = func(rc int) {
		exitCode = rc
		called = true
	}

	defer func() { OsExiter = os.Exit }()

	HandleExitCoder(nil)

	expect(t, exitCode, 0)
	expect(t, called, false)
}

func TestHandleExitCoder_ExitCoder(t *testing.T) {
	exitCode := 0
	called := false

	OsExiter = func(rc int) {
		exitCode = rc
		called = true
	}

	defer func() { OsExiter = os.Exit }()

	HandleExitCoder(Exit("galactic perimeter breach", 9))

	expect(t, exitCode, 9)
	expect(t, called, true)
}

func TestHandleExitCoder_MultiErrorWithExitCoder(t *testing.T) {
	exitCode := 0
	called := false

	OsExiter = func(rc int) {
		exitCode = rc
		called = true
	}

	defer func() { OsExiter = os.Exit }()

	exitErr := Exit("galactic perimeter breach", 9)
	err := newMultiError(errors.New("wowsa"), errors.New("egad"), exitErr)
	HandleExitCoder(err)

	expect(t, exitCode, 9)
	expect(t, called, true)
}
