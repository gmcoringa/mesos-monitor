package mesos

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const PATH = "/metrics/snapshot"

type Mesos struct {
	Client 		http.Client
	URL			string
}

type MesosCollector interface {
	Collect() map[string]interface{}
}

func NewMesos(host string, connectionTimeoutMS uint64)(MesosCollector){
	url := host + PATH
	timeout := time.Duration(connectionTimeoutMS) * time.Millisecond
	client := http.Client{Timeout: timeout}
	
	return &Mesos{Client: client, URL: url}
}

func (m *Mesos) Collect() (map[string]interface{}) {

	response, err := m.Client.Get(m.URL)
	if err != nil {
		fmt.Println("Failed to collect metrics from Mesos server: ", err)
		return nil
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Failed to parse response: ", err)
		return nil
	}

	var f interface{}
	err2 := json.Unmarshal(body, &f)
	if err2 != nil {
		fmt.Println("Failed to decode json: ", err2)
		return nil
	}
	
	return f.(map[string]interface{})
}
