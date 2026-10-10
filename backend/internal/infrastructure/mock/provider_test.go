package mock

import "testing"

func TestListedFileIDsCanBeRetrieved(t *testing.T) {
	listedFiles, err := GetFiles()
	if err != nil {
		t.Fatalf("GetFiles() error = %v", err)
	}

	for _, listedFile := range listedFiles {
		_, exists, err := GetFile(listedFile.ID)
		if err != nil {
			t.Fatalf("GetFile(%q) error = %v", listedFile.ID, err)
		}
		if !exists {
			t.Errorf("GetFile(%q) does not exist", listedFile.ID)
		}
	}
}
