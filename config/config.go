package config

/**
*@Author lyer
*@Date 2020/12/19 14:00
*@Describe
**/

type Server struct {
	Etcd   Etcd   `json:"etcd" yaml:"etcd"`
	System System `json:"system" yaml:"system"`
}
