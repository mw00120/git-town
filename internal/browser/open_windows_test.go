//go:build windows

package browser

import (
	"testing"

	"github.com/git-town/git-town/v24/internal/browser/browserdomain"
	. "github.com/git-town/git-town/v24/pkg/prelude"
	"github.com/shoenig/test/must"
)

func TestOpenOnWindows(t *testing.T) {
	t.Parallel()

	t.Run("uses rundll32 for the default browser", func(t *testing.T) {
		t.Parallel()

		runner := &recordingRunner{}

		Open("https://bitbucket.org/org/repo/pull-requests/new?source=branch&dest=org%2Frepo%3Amain", runner, None[browserdomain.BrowserExecutable](), true)

		must.Eq(t, [][]string{{"rundll32", "url.dll,FileProtocolHandler", "https://bitbucket.org/org/repo/pull-requests/new?source=branch&dest=org%2Frepo%3Amain"}}, runner.calls)
	})

	t.Run("uses the configured browser executable", func(t *testing.T) {
		t.Parallel()

		runner := &recordingRunner{}

		Open("https://example.com", runner, Some(browserdomain.BrowserExecutable("firefox")), true)

		must.Eq(t, [][]string{{"firefox", "https://example.com"}}, runner.calls)
	})
}

type recordingRunner struct {
	calls [][]string
}

func (self *recordingRunner) Run(executable string, args ...string) error {
	self.calls = append(self.calls, append([]string{executable}, args...))
	return nil
}

func (self *recordingRunner) RunWithEnv(_ []string, executable string, args ...string) error {
	return self.Run(executable, args...)
}
