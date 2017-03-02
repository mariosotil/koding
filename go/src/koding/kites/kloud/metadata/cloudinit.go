package metadata

import "errors"

const header = "#cloud-config"

var ErrNotCloudInit = errors.New("not a cloud-init content")

type CloudInit map[string]interface{}

func (ci CloudInit) MergeIn(mixin CloudInit) {

}

func (ci CloudInit) String() string {
	return ""
}

type CloudConfig struct {
}

func NewCloudInit(cfg *CloudConfig) CloudInit {
	return CloudInit{}
}

func ParseCloudInit(p []byte) (CloudInit, error) {
	return nil, nil
}
