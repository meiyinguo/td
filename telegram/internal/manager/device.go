package manager

import (
	"github.com/gotd/td/constant"
	"github.com/gotd/td/tg"
	"strconv"
	"time"
)

// DeviceConfig is config which send when Telegram connection session created.
type DeviceConfig struct {
	// Device model.
	DeviceModel string
	// Operating system version.
	SystemVersion string
	// Application version.
	AppVersion string
	// Code for the language used on the device's OS, ISO 639-1 standard.
	SystemLangCode string
	// Language pack to use.
	LangPack string
	// Code for the language used on the client, ISO 639-1 standard.
	LangCode string
	// Info about an MTProto proxy.
	Proxy tg.InputClientProxy
	// Additional initConnection parameters. For now, only the tz_offset field is supported,
	// for specifying timezone offset in seconds.
	Params tg.JSONValueClass
	//
	AndroidAppFingerPrint string
}

var (
	AndroidAppVersion     = constant.AndroidAppVersion
	AndroidAppFingerPrint = constant.AndroidAppFingerPrint
)

// SetDefaults sets default values.
func (c *DeviceConfig) SetDefaults() {
	const notAvailable = "n/a"

	// Strings must be non-empty, so set notAvailable if default value is empty.
	set := func(to *string, value string) {
		if value != "" {
			*to = value
		} else {
			*to = notAvailable
		}
	}

	if c.DeviceModel == "" {
		set(&c.DeviceModel, "Samsung SM-G9880")
	}
	if c.SystemVersion == "" {
		set(&c.SystemVersion, "SDK 30")
	}
	if c.AppVersion == "" {
		set(&c.AppVersion, AndroidAppVersion)
	}
	if c.SystemLangCode == "" {
		c.SystemLangCode = "en-us"
	}
	if c.LangCode == "" {
		c.LangCode = "en"
	}
	if c.AndroidAppFingerPrint == "" {
		c.AndroidAppFingerPrint = AndroidAppFingerPrint
	}
	if c.LangPack == "" {
		c.LangPack = "android"
	}
	if c.Params == nil {
		deviceToken := "__FIREBASE_GENERATING_SINCE_" + strconv.Itoa(int(time.Now().Unix())) + "__"
		c.Params = &tg.JSONObject{
			Value: []tg.JSONObjectValue{
				{Key: "device_token", Value: &tg.JSONString{Value: deviceToken}},
				{Key: "data", Value: &tg.JSONString{Value: c.AndroidAppFingerPrint}},
				{Key: "installer", Value: &tg.JSONString{Value: ""}},
				{Key: "package_id", Value: &tg.JSONString{Value: "org.telegram.messenger"}},
				{Key: "tz_offset", Value: &tg.JSONNumber{Value: 0}},
				{Key: "perf_cat", Value: &tg.JSONNumber{Value: 1}},
			},
		}
	}

}
