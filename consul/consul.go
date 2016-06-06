package consul

import (
	"bytes"
	"encoding/json"
	"github.com/andresvia/oasproj/project"
	"net/http"
)

type Service struct {
	Name    string `json:"Service"`
	Address string
	Tags     []string
	Port    int64
}

type Catalog struct {
	Datacenter string
	Node       string
	Address    string
	Service    Service
}

func (cat *Catalog) Register() (err error) {
	client := &http.Client{}
	var payload []byte
	payload, err = json.Marshal(&cat)
	// println(string(payload))
	body := bytes.NewReader(payload)
	var req *http.Request
	if req, err = http.NewRequest("PUT", defaultConsulUrl+"/v1/catalog/register", body); err == nil {
		_, err = client.Do(req)
	}
	return
}

var defaultConsulUrl = "http://localhost:8500"

func New(p project.Project) (c Catalog, err error) {
	s := Service{
		Name:    p.Project_name,
		Address: "0.0.0.0",
	}
	if p.Programming_language != "" {
		s.Tags = append(s.Tags, "lang:" + p.Programming_language)
	}
	if p.Project_framework != "" {
		s.Tags = append(s.Tags, "framework:"+p.Project_framework)
	}
	if p.Programming_language != "" {
		s.Tags = append(s.Tags, "orgunit:"+p.Organizational_unit)
	}

	c = Catalog{
		Datacenter: "dc1",
		Service:    s,
		Node:       "oasproj-register",
		Address:    "0.0.0.0",
	}
	return
}
