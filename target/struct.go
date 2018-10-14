package target

type targetStruct struct {
	Targets []string          `json:"targets"`
	Labels  map[string]string `json:"labels"`
}
