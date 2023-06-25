package libutils

import (
	"testing"
)

func TestError(t *testing.T) {
	project_name := "test_project"
	list_test := []struct {
		desc          string
		function_name string
		package_name  string
		message       string
		level         string
	}{
		{"Error1", "func1", "package1", "ErrorMessage1", "Error"},
		{"Info1", "func1", "package1", "InfoMessage1", "Info"},
		{"Error2", "func2", "package3", "ErrorMessage2 2 2", "Error"},
		{"Info2", "func2", "package2", "Info Info Info", "Info"},
	}

	InitLogger(project_name)
	for _, test := range list_test {
		t.Run(test.desc, func(t *testing.T) {
			if test.level == "Error" {
				Error(test.package_name, test.function_name, test.message)
			} else {
				Info_msg(test.package_name, test.function_name, test.message)
			}
		})
	}
}
