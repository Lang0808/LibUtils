package libutils

import (
	"testing"
)

func TestCreateFile(t *testing.T) {
	list_test := []struct {
		desc      string
		file_path string
	}{
		{"create_file_in_root", "/remove_me_please.log"}, // fail case: permission denied
		{"create_file_in_folder1_root", "/tmp/remove_me_please.log"},
		{"create_file_in_folder2_root", "/tmp/exist_folder/remove_me_please.log"},
		{"create_file_current_dir", "remove_me_please.log"},
		{"create_file_current_dir-folder_1", "tmp/remove_me_please.log"},
		{"create_file_current_dir-folder_2", "tmp/tmp2/remove_me_please.log"},
	}

	for _, test := range list_test {
		t.Run(test.desc, func(t *testing.T) {
			err := CreateFile(test.file_path)
			if err != nil {
				t.Fatalf("%v; Error: %v;\n", test.desc, err)
			}
			if !IsExist(test.file_path) {
				t.Fatalf("%v; file_path not exist;\n", test.desc)
			}
		})
	}
}
