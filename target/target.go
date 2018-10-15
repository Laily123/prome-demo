package target

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
	"time"
)

const (
	tag_target_id    = "target_id"
	tag_target_group = "target_group"
	tag_instance     = "instance"
)

// AddTarget 添加 scrape target
// default metrics path is "/metrics"

// configFileName: file_sd_config file name
// addr: scrape metrics address, like "xx.com:8080"
// targetID: unique for targets. can use hostname+port
// group: add a group for target
func AddTarget(configFileName string, addr []string, targetID, group, instance string) error {
	c := &ConfigFile{}
	c.path = configFileName
	fileObj, err := os.OpenFile(c.path, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer fileObj.Close()
	c.f = fileObj

	configs, err := c.readConfigFile()
	if err != nil {
		return err
	}
	// 重复的 targe 直接认为添加成功
	for _, v := range configs {
		if id, ok := v.Labels[tag_target_id]; ok {
			if id == targetID {
				return nil
			}
		}
	}
	target := &targetStruct{
		Targets: addr,
		Labels: map[string]string{
			tag_target_id:    targetID,
			tag_target_group: group,
			tag_instance:     instance,
		},
	}
	configs = append(configs, target)

	err = c.writeConfigFile(configs)
	return err
}

type ConfigFile struct {
	f    *os.File
	path string
}

func (c *ConfigFile) readConfigFile() ([]*targetStruct, error) {
	configs := make([]*targetStruct, 0)
	// 文件加锁，重试三次
	var lockerr error
	for i := 0; i < 3; i++ {
		if lockerr != nil {
			time.Sleep(50 * time.Millisecond)
		}
		lockerr = syscall.Flock(int(c.f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
		if lockerr == nil {
			break
		}
	}
	if lockerr != nil {
		return nil, fmt.Errorf("can not modify config file: %s\n", lockerr)
	}

	fileData, err := ioutil.ReadAll(c.f)
	if err != nil {
		return nil, err
	}
	if string(fileData) != "[]" {
		err = json.Unmarshal(fileData, &configs)
		if err != nil {
			return nil, err
		}
	}
	return configs, nil
}

func (c *ConfigFile) writeConfigFile(configs []*targetStruct) error {
	data, _ := json.MarshalIndent(configs, "", "  ")
	c.f.Truncate(0)
	c.f.Seek(0, 0)
	_, err := c.f.Write(data)
	if err != nil {
		return err
	}
	return nil
}
