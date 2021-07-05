package cmd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"nina/fixtures"
	"os"
	"strings"
	"testing"
)

func TestListTimers(t *testing.T) {
	timers := fixtures.Timers()
	jsonValue, _ := json.Marshal(timers)

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/timers" {
			t.Fatalf("Incorrect request URI. expected /timers. got: %s", r.RequestURI)
		}

		w.Write(jsonValue)
	}))
	defer svr.Close()

	os.Setenv("NOKO_ACCESS_TOKEN", "noko1234")
	os.Setenv("NOKO_BASE_URL", svr.URL)

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	args := []string{"timers", "list"}
	cmd := rootCmd
	cmd.SetArgs(args)
	cmd.Execute()

	w.Close()

	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	text := string(out)

	expected := `
Project 1                       1h12,  running: This is foo. This is bar
Project 2                       1h49,   paused: This is fix. This is bax
	`

	if strings.TrimSpace(text) != strings.TrimSpace(expected) {
		t.Fatalf("Incorrect output in listCmd. expected:\n%s\n. got\n%s\n", expected, text)
	}
}
