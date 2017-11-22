package server

import (
	"fmt"

	"github.com/rancher/go-rancher/v2"
)

func NewRancherClient(url string) (*client.RancherClient, error) {
	opts := &client.ClientOpts{
		Url: "http://10.20.3.10:8181/v1/schemas",
	}
	return client.NewRancherClient(opts)
}

func GetRancherHostPublicKey(rClient *client.RancherClient, hostUUID string) (string, error) {
	hosts, err := rClient.Host.List(&client.ListOpts{
		Filters: map[string]interface{}{
			"uuid": hostUUID,
		},
	})
	if err != nil {
		return "", err
	}

	if len(hosts.Data) >= 0 {
		return hosts.Data[0].Info.(map[string]interface{})["hostKey"].(map[string]interface{})["data"].(string), nil
	}

	return "", fmt.Errorf("Host: %s not found", hostUUID)
}
