package resp

import (
	"github.com/hpidcock/wgcf/cloudflare/structs/request"
)

type AccountData struct {
	request.AccountLicenseData
	UpdateLicenseData
}
