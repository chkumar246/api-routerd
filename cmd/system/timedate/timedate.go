// SPDX-License-Identifier: Apache-2.0

package timedate

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/RestGW/api-routerd/cmd/share"

	log "github.com/sirupsen/logrus"
)

const (
	dbusInterface = "org.freedesktop.timedate1"
	dbusPath      = "/org/freedesktop/timedate1"
)

var timeInfo = map[string]string{
	"Timezone":        "",
	"LocalRTC":        "",
	"CanNTP":          "",
	"NTP":             "",
	"NTPSynchronized": "",
	"TimeUSec":        "",
	"RTCTimeUSec":     "",
}

var timeDateMethod = map[string]string{
	"SetTime":       "",
	"SetTimezone":   "",
	"SetLocalRTC":   "",
	"SetNTP":        "",
	"ListTimezones": "",
}

//TimeDate JSON message
type TimeDate struct {
	Property string `json:"property"`
	Value    string `json:"value"`
}

//SetTimeDate set timedate property
func (t *TimeDate) SetTimeDate() error {
	conn, err := share.GetSystemBusPrivateConn()
	if err != nil {
		log.Errorf("Failed to get systemd bus connection: %v", err)
		return err
	}
	defer conn.Close()

	_, k := timeDateMethod[t.Property]
	if !k {
		return fmt.Errorf("Failed to set timedate:  %s not found", t.Property)
	}

	h := conn.Object(dbusInterface, dbusPath)

	if t.Value == "SetNTP" {

		b, err := share.ParseBool(t.Value)
		if err != nil {
			return err
		}

		r := h.Call(dbusInterface+"."+t.Property, 0, b, false).Err
		if r != nil {
			log.Errorf("Failed to set SetNTP: %s", r)
			return r
		}
	} else {

		r := h.Call(dbusInterface+"."+t.Property, 0, t.Value, false).Err
		if r != nil {
			log.Errorf("Failed to set timedate property: %s", r)
			return r
		}
	}

	return nil
}

//GetTimeDate gets property from timedated
func GetTimeDate(rw http.ResponseWriter, property string) error {
	conn, err := share.GetSystemBusPrivateConn()
	if err != nil {
		log.Errorf("Failed to get dbus connection: %v", err)
		return err
	}
	defer conn.Close()

	h := conn.Object(dbusInterface, dbusPath)
	for k := range timeInfo {
		p, perr := h.GetProperty("org.freedesktop.timedate1." + k)
		if perr != nil {
			log.Errorf("Failed to get org.freedesktop.timedate1.%s", k)
			continue
		}

		switch k {
		case "Timezone":
			v, b := p.Value().(string)
			if !b {
				continue
			}

			timeInfo[k] = v
			break
		case "LocalRTC":
			v, b := p.Value().(bool)
			if !b {
				continue
			}

			timeInfo[k] = strconv.FormatBool(v)

			break

		case "CanNTP":
			v, b := p.Value().(bool)
			if !b {
				continue
			}

			timeInfo[k] = strconv.FormatBool(v)

			break
		case "NTP":
			v, b := p.Value().(bool)
			if !b {
				continue
			}

			timeInfo[k] = strconv.FormatBool(v)

			break
		case "NTPSynchronized":
			v, b := p.Value().(bool)
			if !b {
				continue
			}

			timeInfo[k] = strconv.FormatBool(v)

			break
		case "TimeUSec":
			v, b := p.Value().(uint64)
			if !b {
				continue
			}

			t := time.Unix(0, int64(v))
			timeInfo[k] = t.String()

		case "RTCTimeUSec":
			v, b := p.Value().(uint64)
			if !b {
				continue
			}

			t := time.Unix(0, int64(v))

			timeInfo[k] = t.String()
			break
		}
	}

	if property == "" {
		return share.JSONResponse(timeInfo, rw)
	}

	t := TimeDate{
		Property: property,
		Value:    timeInfo[property],
	}

	return share.JSONResponse(t, rw)
}
