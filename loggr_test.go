// Tests for loggr - a minimalist logger
package loggr

import "testing"
import "bytes"
import "strings"

// Test log.Info()
func TestInfo(t *testing.T) {
	s := "Space"
	buf := bytes.NewBufferString(s)

	log := Log{
		Verbose: false,
		Writer:  buf,
	}

	input := "Earth"
	expected := "info: Earth\n"

	// Call log
	log.Info(input)

	// Get output
	actual := strings.Split(buf.String(), " - ")

	// Test
	if actual[1] != expected {
		t.Errorf("log.Info(%q), expected %q, got %q", input, expected, actual[1])
	}
}

// Test log.Error()
func TestError(t *testing.T) {
	s := "Space"
	buf := bytes.NewBufferString(s)

	log := Log{
		Verbose: false,
		Writer:  buf,
	}

	input := "Moon"
	expected := "error: Moon\n"

	// Call log
	log.Error(input)

	// Get output
	actual := strings.Split(buf.String(), " - ")

	// Test
	if actual[1] != expected {
		t.Errorf("log.Error(%q), expected %q, got %q", input, expected, actual[1])
	}
}

// Test log.Debug()
func TestDebug(t *testing.T) {
	s := "Space"
	buf := bytes.NewBufferString(s)

	log := Log{
		Verbose: true,
		Writer:  buf,
	}

	input := "Sun"
	expected := "debug: Sun\n"

	// Call log
	log.Debug(input)

	// Get output
	actual := strings.Split(buf.String(), " - ")

	// Test
	if actual[1] != expected {
		t.Errorf("log.Debug(%q), expected %q, got %q", input, expected, actual[1])
	}
}

//Test Log.Verbose = false (no debug)
func TestVerbose(t *testing.T) {
	s := "Space"
	buf := bytes.NewBufferString(s)

	log := Log{
		Verbose: false,
		Writer: buf,
	}

	input := "Sun"
	expected := "Space"

	// Call log
	log.Debug(input)

	// Get output
	actual := buf.String()

	// Test
	if actual != expected {
		t.Errorf("log.Debug(%q), expected %q, got %q", input, expected, actual[1])
	}
}

//Test Log.Verbose = false (no debug)
func TestNilWriter(t *testing.T) {
	s := "Space"
	buf := bytes.NewBufferString(s)

	log := Log{
		Verbose: false,
		// Writer: buf, --> no writer declared, will error
	}

	input := "Sun"
	expected := "Space"

	// Call log
	err := log.Info(input)

	// Get output
	actual := buf.String()

	// Test
	if actual != expected {
		t.Errorf("log.Info(%q), expected %q, got %q", input, expected, actual[1])
	}
	if err != ErrInvalidWriter {
		t.Errorf("log.Info(%q), expected %q, got %q", input, ErrInvalidWriter, err)
	}

}
