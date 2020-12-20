package config

/**
*@Author lyer
*@Date 2020/12/19 14:00
*@Describe
**/

type Etcd struct {
	Endpoints []string `json:"endpoints" yaml:"endpoints"`
	Timeout int `json:"timeout" yaml:"timeout"`
}