package libutils

import (
	"fmt"
	"os"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestGet(t *testing.T) {
	err := InitConfig("../config/app.yaml")
	if err != nil {
		path, err := os.Getwd()
		if err != nil {
			t.Fatalf("InitConfig fail. Get current working directory fail")
		}
		fmt.Printf("current directory: %v\n", path)
		t.Fatalf("InitConfig fail\n")
	}

	list_test := []struct {
		desc string
		key  string
		want string
	}{
		{"get_config_1_level", "test", "test_1_level"},
		{"get_config_2_level", "test2.test_at_level_2", "test_2_level"},
		{"get_config_1_level_fail", "not_exist", ""},
		{"get_config_2_level_fail_1", "not_exist.not_exist_2", ""},
		{"get_config_2_level_fail_2", "test.not_exist_2", ""},
	}

	for _, test := range list_test {
		t.Run(test.desc, func(t *testing.T) {
			got := Get(test.key)
			assert.Equal(t, test.want, got, "Incorrect value")
		})
	}
}
