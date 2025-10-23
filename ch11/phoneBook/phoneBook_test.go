package main

import (
	"os"
	"path"
	"slices"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestReadCSVFile(t *testing.T) {
	t1 := path.Join(os.TempDir(), "test.csv")
	f, err := os.Create(t1)
	if err != nil {
		t.Error("os.Create() failed:", err)
		return
	}

	defer t.Cleanup(func() {
		err = os.Remove(t1)
		if err != nil {
			t.Error("os.Remove() failed:", err)
		}
	})

	entriesCount := 3
	entries := make([]Entry, entriesCount)
	payload := ""
	for i := range entriesCount {
		strI := strconv.Itoa(i + 1)

		entry := Entry{
			Name:       "Name" + strI,
			Surname:    "Surname" + strI,
			Tel:        strings.Repeat(strI, 5),
			LastAccess: strconv.FormatInt(time.Now().Add(time.Duration(i)*time.Second).Unix(), 10),
		}
		entries[i] = entry

		payload += strings.Join([]string{
			entry.Name,
			entry.Surname,
			entry.Tel,
			entry.LastAccess,
		}, ",") + "\n"
	}

	_, err = f.Write([]byte(payload))
	if err != nil {
		t.Error("f.Write() failed:", err)
		return
	}

	err = f.Close()
	if err != nil {
		t.Error("f.Close() failed:", err)
		return
	}

	err = readCSVFile(t1)
	if err != nil {
		t.Error("readCSVFile failed:", err)
		return
	}

	if !slices.Equal(entries, data) {
		t.Error("Different slices", entries, data)
	}
}
