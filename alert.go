package main

type Alert struct {
	Receiver          string            `json:"receiver" binding:"required"`
	Status            string            `json:"status" binding:"required"`
	ExternalURL       string            `json:"externalURL" binding:"required"`
	Version           string            `json:"version" binding:"required"`
	GroupKey          string            `json:"groupKey" binding:"required"`
	GroupLabels       map[string]string `json:"groupLabels" binding:"required"`
	CommonLabels      map[string]string `json:"commonLabels" binding:"required"`
	CommonAnnotations map[string]string `json:"commonAnnotations" binding:"required"`
	Alerts            []SingleAlert     `json:"alerts" binding:"required"`
}

func (a *Alert) IsResolved() bool {
	return a.Status == "resolved"
}

type SingleAlert struct {
	Status       string            `json:"status" binding:"required"`
	GeneratorURL string            `json:"generatorURL" binding:"required"`
	Labels       map[string]string `json:"labels" binding:"required"`
	Annotations  map[string]string `json:"annotations" binding:"required"`
}

func (a *SingleAlert) IsResolved() bool {
	return a.Status == "resolved"
}

type GroupLabelsStruct struct {
	Alertname string `json:"alertname" binding:"required"`
}
