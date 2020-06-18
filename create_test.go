package main

import (
	"testing"

	"github.com/thealamu/cvee/section"
)

type mockSection struct {
	Name string
}

func (m *mockSection) Edit() error {
	m.Name = "MOCK"
	return nil
}

func TestEditSections(t *testing.T) {
	m := &mockSection{}
	mockSections := []section.Section{
		m,
	}
	//Verify editSections calls the Edit method of sections
	if err := editSections(mockSections); err != nil {
		t.Errorf("editSections returns error, no error expected")
		return
	}
	if m.Name != "MOCK" {
		t.Errorf("editSections does not call Edit method of sections")
	}
}
