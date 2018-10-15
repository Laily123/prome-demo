package target

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

var (
	ConfigDir = "/Users/Laily/go/src/prome-demo"
)

type requestParms struct {
	Addr     string `json:"addr"`
	TargetID string `json:"target_id"`
	JobName  string `json:"job_name"`
	Instance string `json:"instance"`
	Group    string `json:"group"`
}

type respStruct struct {
	Err    int         `json:"err"`
	ErrMsg string      `json:"err_msg"`
	Data   interface{} `json:"data"`
}

func AddTargetHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		makeResponse(w, "", err.Error())
		return
	}
	params := &requestParms{}
	err = json.Unmarshal(data, params)
	if err != nil {
		makeResponse(w, "", err.Error())
		return
	}
	res := checkParams(params)
	if res != "" {
		makeResponse(w, "", res)
		return
	}

	configFile := params.JobName + ".json"
	path := filepath.Join(ConfigDir, configFile)
	err = AddTarget(path, []string{params.Addr}, params.TargetID, params.Group, params.Instance)
	if err != nil {
		makeResponse(w, "", err.Error())
		return
	}
	makeResponse(w, "add target succ", "")
	return
}

func checkParams(params *requestParms) string {
	if params.Addr == "" {
		return "addr is empty"
	}
	if params.TargetID == "" {
		return "target_id is empty"
	}
	if params.JobName == "" {
		return "job_name is empty"
	}
	if params.Instance == "" {
		return "instance is empty"
	}
	return ""
}

func makeResponse(w http.ResponseWriter, data interface{}, errMsg string) error {
	errCode := 0
	if errMsg != "" {
		errCode = 1
	}
	resp := &respStruct{
		Err:    errCode,
		ErrMsg: errMsg,
		Data:   data,
	}
	respBytes, _ := json.Marshal(resp)
	_, err := w.Write(respBytes)
	return err
}
