package commands

import "testing"

func TestPatternCorrectnes(t *testing.T) {
	parts := convertTimePattern.FindStringSubmatch("15:40 GMT in America/San_Francisco")

	if len(parts) != 4 {
		t.Error("expected length of parts to be", 4, "got", len(parts))
	}

	if parts[1] != "15:40" {
		t.Error("expected", "15:40", "got", parts[1])
	}

	if parts[2] != "GMT" {
		t.Error("expected", "GMT", "got", parts[2])
	}

	if parts[3] != "America/San_Francisco" {
		t.Error("expected", "America/San_Francisco", "got", parts[3])
	}

	if convertTimePattern.MatchString("3pm in 'MURICA") {
		t.Error("expected", "3pm in 'MURICA", "not to match")
	}
}

func TestHandler(t *testing.T) {
	err, retMsg:= convertTimeHandler("15:40 Europe/Berlin in America/Los_Angeles", nil)
	if err != nil {
		t.Fatal("Errored with", err)
	}

	if retMsg != "06:40" {
		t.Error("Expected", "06:40", "got", retMsg)
	}
}
