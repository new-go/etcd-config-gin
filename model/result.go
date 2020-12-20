package model

/**
*@Author lyer
*@Date 2020/12/19 15:09
*@Describe
**/

type Result struct {
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg,omitempty"`
	Code string      `json:"code,omitempty"`
}
