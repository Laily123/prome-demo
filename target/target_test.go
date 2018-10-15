package target

import (
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	configFilePath := "/Users/Laily/go/src/prome-demo/test.json"
	c := ConfigFile{}
	c.path = configFilePath
	fileObj, err := os.OpenFile(c.path, os.O_RDWR, os.ModePerm)
	if err != nil {
		t.Error(err)
	}
	defer fileObj.Close()
	c.f = fileObj

	configs, err := c.readConfigFile()
	if err != nil {
		t.Error(err)
	}
	t.Log(configs)
}

func TestAddTarget(t *testing.T) {
	configFilePath := "/Users/Laily/go/src/prome-demo/test.json"
	addr := []string{"http://127.0.0.1"}
	targetID := "xxxx2"
	group := ""
	instance := "instance"
	err := AddTarget(configFilePath, addr, targetID, group, instance)
	t.Log(err)
}
