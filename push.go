package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
)

var (
	pushCounter = 0
	httpClient  = &http.Client{}
)

func initPush() {
	go func() {
		pushCounter += rand.Intn(10)

	}()
}

func push(value int) error {
	jobName := "service"
	instance := "test"
	url := fmt.Sprintf("http://mix1:9023/metrics/job/%s/instance/%s", jobName, instance)
	data := fmt.Sprintf("prome_demo_counter %d\n", value) // \n 一定要有
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return err
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(res))
	return nil
}
