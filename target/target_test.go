package target

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	configFilePath := "/Users/Laily/go/src/prome-demo/test.json"
	c := ConfigFile{}
	c.path = configFilePath
	configs, err := c.readConfigFile()
	if err != nil {
		t.Error(err)
	}
	t.Log(configs)
}

func TestAddTarget(t *testing.T) {
	configFilePath := "/Users/Laily/go/src/prome-demo/test.json"
	addr := []string{"http://127.0.0.1"}
	targetID := "xxxx1"
	group := ""
	err := AddTarget(configFilePath, addr, targetID, group)
	t.Log(err)
}
