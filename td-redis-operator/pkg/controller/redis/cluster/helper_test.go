package cluster

import (
	"td-redis-operator/pkg/apis/cache/v1alpha1"
	"testing"

	corev1 "k8s.io/api/core/v1"
)

func TestController_createRedisCluster(t *testing.T) {
	type args struct {
		pods []*corev1.Pod
		mp   *v1alpha1.RedisCluster
	}
	tests := []struct {
		name    string
		c       *Controller
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.createRedisCluster(tt.args.pods, tt.args.mp); (err != nil) != tt.wantErr {
				t.Errorf("Controller.createRedisCluster() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
