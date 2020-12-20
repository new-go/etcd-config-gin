package router

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/biningo/etcd-config-gin/global"
	"github.com/biningo/etcd-config-gin/model"
	"github.com/gin-gonic/gin"
	"go.etcd.io/etcd/clientv3"
)

/**
*@Author lyer
*@Date 2020/12/19 14:48
*@Describe
**/

var (
	ErrEtcdNotFound = errors.New("etcd not found")
	ErrJsonBind     = errors.New("file upload error")
	ErrJsonMarshal  = errors.New("json marshal error")
	ErrEtcdPut      = errors.New("etcd put error")
	ErrFileUpload   = errors.New("file upload fail")
)

func GetConfig(ctx *gin.Context) {
	key := ctx.Query("key")
	resp, err := global.G_ETCD.Get(context.TODO(), key)
	if err != nil {
		ctx.JSON(406, model.Result{Msg: err.Error()})
		return
	}
	kvs := resp.Kvs
	if len(kvs) <= 0 {
		ctx.JSON(406, model.Result{Msg: ErrEtcdNotFound.Error()})
		return
	}
	kv := kvs[0]
	result := make(map[string]interface{})
	if err := json.Unmarshal(kv.Value, &result); err != nil {
		ctx.JSON(406, model.Result{Msg: ErrJsonMarshal.Error()})
		return
	}
	ctx.JSON(200, model.Result{Data: result})
}

func PutConfig(ctx *gin.Context) {
	cfgBody := make(map[string]interface{})
	if err := ctx.BindJSON(&cfgBody); err != nil {
		ctx.JSON(406, model.Result{Msg: ErrJsonBind.Error()})
		return
	}

	k := cfgBody["key"].(string)
	v, err := json.Marshal(cfgBody["val"])
	if err != nil {
		ctx.JSON(406, model.Result{Msg: ErrJsonMarshal.Error()})
		return
	}

	resp, err := global.G_ETCD.Put(context.TODO(), k, string(v))
	if err != nil {
		ctx.JSON(406, model.Result{Msg: ErrEtcdPut.Error()})
		return
	}
	ctx.JSON(200, model.Result{Data: resp, Msg: "success"})
}

func DelConfig(ctx *gin.Context) {
	key := ctx.Query("key")
	resp, err := global.G_ETCD.Delete(context.TODO(), key)
	if err != nil {
		ctx.JSON(406, model.Result{Data: resp, Msg: err.Error()})
		return
	}
	ctx.JSON(200, model.Result{Data: resp, Msg: "success"})
}

func ListConfig(ctx *gin.Context) {
	prefix := ctx.Query("prefix")
	resp, err := global.G_ETCD.Get(context.TODO(), prefix, clientv3.WithPrefix())
	if err != nil {
		ctx.JSON(406, model.Result{Msg: err.Error()})
		return
	}

	kvs := resp.Kvs
	if len(kvs) <= 0 {
		ctx.JSON(406, model.Result{Msg: ErrEtcdNotFound.Error()})
		return
	}

	arr := []string{}
	for _, kv := range resp.Kvs {
		arr = append(arr, string(kv.Key))
	}
	ctx.JSON(200, model.Result{Data: arr})
}

func UploadConfig(ctx *gin.Context) {
	f, err := ctx.FormFile("config")
	if err != nil {
		ctx.JSON(406, model.Result{Msg: ErrFileUpload.Error()})
		return
	}
	r, err := f.Open()
	if err != nil {
		ctx.JSON(406, model.Result{Msg: ErrFileUpload.Error()})
		return
	}
	content := make([]byte, f.Size)
	_, err = r.Read(content)
	resp, err := global.G_ETCD.Put(context.TODO(), "/config/"+f.Filename, string(content))
	if err != nil {
		ctx.JSON(406, model.Result{Msg: ErrEtcdPut.Error()})
		return
	}
	ctx.JSON(200, model.Result{Msg: "success", Data: resp})
}

func DownloadConfig(ctx *gin.Context) {
	key := ctx.Query("key")
	//fileType := ctx.Query("fileType")
	resp, err := global.G_ETCD.Get(context.TODO(), key)
	if err != nil {
		ctx.JSON(406, model.Result{Msg: err.Error()})
		return
	}
	kv := resp.Kvs[0]
	ctx.Header("Content-Disposition", "attachment; filename="+string(kv.Key))
	ctx.Writer.Write(kv.Value)
}
