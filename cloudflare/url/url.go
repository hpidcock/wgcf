package url

import (
	"github.com/hpidcock/wgcf/util"
)

var apiVersion = "v0a884"
var apiEndpoint = util.JoinUrls("https://api.cloudflareclient.com", apiVersion)

var RegUrl = util.JoinUrls(apiEndpoint, "reg")
var ClientConfigUrl = util.JoinUrls(apiEndpoint, "client_config")

var TermsUrl = "https://www.cloudflare.com/application/terms/"

// returns DeviceData
func GetDeviceUrl(deviceId string) string {
	return util.JoinUrls(RegUrl, deviceId)
}

// returns AccountData
func GetAccountUrl(deviceId string) string {
	return util.JoinUrls(GetDeviceUrl(deviceId), "account")
}

// returns BoundDevicesData
func GetBoundDevicesUrl(deviceId string) string {
	return util.JoinUrls(GetAccountUrl(deviceId), "devices")
}

// creates and returns a new license key that replaces the old key
func GetRecreateLicenseKeyUrl(deviceId string) string {
	return util.JoinUrls(GetAccountUrl(deviceId), "license")
}

// allows operating on devices bound to this device's account
func GetBoundDeviceUrl(deviceId string, targetDeviceId string) string {
	return util.JoinUrls(GetAccountUrl(deviceId), "reg", targetDeviceId)
}
