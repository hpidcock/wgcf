package api

import (
	"bytes"
	"encoding/json"
	"github.com/hpidcock/wgcf/cloudflare/structs/request"
	"github.com/hpidcock/wgcf/cloudflare/structs/resp"
	"github.com/hpidcock/wgcf/cloudflare/url"
	"github.com/hpidcock/wgcf/cloudflare/util"
	"github.com/hpidcock/wgcf/config"
)

func GetDeviceData(ctx *config.Context) (*resp.DeviceData, error) {
	var result resp.DeviceData
	if err := util.NewAuthenticatedRequest("GET", url.GetDeviceUrl(ctx.DeviceId), nil, ctx.AccessToken, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func EnableWarp(ctx *config.Context) (bool, error) {
	data := request.DeviceWarpData{WarpEnabled: true}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return false, err
	}
	var result resp.DeviceData
	if err := util.NewAuthenticatedRequest("PATCH", url.GetDeviceUrl(ctx.DeviceId),
		bytes.NewBuffer(dataBytes), ctx.AccessToken, &result); err != nil {
		return false, err
	}
	return result.WarpEnabled, nil
}
