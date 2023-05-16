package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/luizgustavoob/second-process/internal/core/consts"
	"github.com/luizgustavoob/second-process/internal/core/entities"
)

type (
	Client interface {
		Do(req *http.Request) (*http.Response, error)
	}

	Registry struct {
		client Client
	}
)

func NewRegistry(client Client) *Registry {
	return &Registry{
		client: client,
	}
}

func (r *Registry) Register() error {
	addr, err := IPAddress()
	if err != nil {
		log.Printf("error getting ip address: %s", err.Error())
		return err
	}

	srv := entities.RemoteService{
		ProcessName: consts.ProcessName,
		Address:     fmt.Sprintf("http://%s:%s", addr, consts.ServerPort),
	}

	js, _ := json.Marshal(srv)
	req, err := http.NewRequest(http.MethodPost, consts.RegistryUrl, bytes.NewBuffer(js))
	if err != nil {
		log.Printf("error creating register request: %s", err.Error())
		return err
	}

	_, err = r.client.Do(req)
	if err != nil {
		log.Printf("service %s not registered: %s", consts.ProcessName, err.Error())
		return err
	}

	return nil
}

func (r *Registry) Unregister() {
	endpoint := fmt.Sprintf("%s/%s", consts.RegistryUrl, consts.ProcessName)
	req, err := http.NewRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		log.Printf("error creating unregister request: %s", err.Error())
		return
	}

	_, err = r.client.Do(req)
	if err != nil {
		log.Printf("service %s not unregistered: %s", consts.ProcessName, err.Error())
		return
	}
}
