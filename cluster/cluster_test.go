package cluster

import (
	"fmt"
	"testing"
)

func TestNewCluster(t *testing.T) {
	file := "../cluster.yml"
	cluster, err := NewCluster(file)
	if err != nil {
		t.Errorf("Parse Cluster Failed:%s\n", err.Error())
	}
	fmt.Printf("Cluster Parsed:\n%+v", cluster)
}

func TestContainerPort(t *testing.T) {
	var port ContainerPort
	port = "~20000:456"
	hp, cp, err := port.GetPorts()
	if err != nil {
		t.Errorf("%d, %d, %s\n", hp, cp, err.Error())
	}
	t.Logf("%d, %d\n", hp, cp)
}
