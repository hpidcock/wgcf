package api

import (
	"bytes"
	"encoding/json"

	"github.com/hpidcock/wgcf/cloudflare/structs/request"
	"github.com/hpidcock/wgcf/cloudflare/structs/resp"
	"github.com/hpidcock/wgcf/cloudflare/url"
	"github.com/hpidcock/wgcf/cloudflare/util"
	"github.com/hpidcock/wgcf/wireguard"
)

func Register(encodedPrivateKey string, deviceModel string) (*resp.RegistrationData, error) {
	timestamp := util.GetTimestamp()
	privateKey, err := wireguard.NewKey(encodedPrivateKey)
	if err != nil {
		return nil, err
	}
	publicKey := privateKey.Public()
	data := request.RegistrationData{
		PublicKey: publicKey.String(),
		InstallID: "",
		FcmToken:  "",
		Tos:       timestamp,
		Model:     deviceModel,
		Type:      "Android",
		Locale:    "en_US",
	}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var result resp.RegistrationData
	if err := util.NewUnauthenticatedRequest("POST", url.RegUrl, bytes.NewBuffer(dataBytes), &result); err != nil {
		return nil, err
	}
	return &result, nil
}
