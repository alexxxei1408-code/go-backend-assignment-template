package main

import "testing"

func TestFormatLog_Basic(t *testing.T) {
	records := []LogRecord{
		{Level: "info", Message: "started"},
		{Level: "warn", Message: "low disk"},
		{Level: "error", Message: "failed"},
	}

	got := FormatLog(records)

	want := "INFO: started\nWARN: low disk\nERROR: failed\n"
	if got != want {
		t.Fatalf("FormatLog:\n got: %q\nwant: %q", got, want)
	}
}

func TestFormatLog_Empty(t *testing.T) {
	if got := FormatLog(nil); got != "" {
		t.Fatalf("FormatLog(nil): expected empty string, got %q", got)
	}
	if got := FormatLog([]LogRecord{}); got != "" {
		t.Fatalf("FormatLog(empty slice): expected empty string, got %q", got)
	}
}


