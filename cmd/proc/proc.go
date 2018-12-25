// SPDX-License-Identifier: Apache-2.0

package proc

import (
	"api-routerd/cmd/share"
	"encoding/json"
	"errors"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/load"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
)

const ProcMiscPath = "/proc/misc"
const ProcNetArpPath = "/proc/net/arp"

type NetArp struct {
	IPAddress string `json:"ip_address"`
	HWType    string `json:"hw_type"`
	Flags     string `json:"flags"`
	HWAddress string `json:"hw_address"`
	Mask      string `json:"mask"`
	Device    string `json:"device"`
}

func GetVersion(rw http.ResponseWriter) (error) {
	infostat, err := host.Info()
	if err != nil {
		return err
	}

	j, err := json.Marshal(infostat)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return errors.New("Json encoding: Version")
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(j)

	return nil
}

func GetNetStat(rw http.ResponseWriter, protocol string) (error) {
	conn, err := net.Connections(protocol)
	if err != nil {
		return err
	}

	j, err := json.Marshal(conn)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return errors.New("Json encoding netstat")
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(j)

	return nil
}

func GetNetDev(rw http.ResponseWriter) (error) {
	netdev, err := net.IOCounters(true)
	if err != nil {
		return err
	}

	j, err := json.Marshal(netdev)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return errors.New("Json encoding NetDev")
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(j)

	return nil
}

func GetInterfaceStat(rw http.ResponseWriter) (error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return err
	}

	j, err := json.Marshal(interfaces)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return errors.New("Json encoding interface stat")
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(j)

	return nil
}

func GetSwapMemoryStat(rw http.ResponseWriter) (error) {
	swap, err := mem.SwapMemory()
	if err != nil {
		return err
	}

	j, err := json.Marshal(swap)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return errors.New("Json encoding memory stat")
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(j)

	return nil
}

func GetVirtualMemoryStat(rw http.ResponseWriter) (error) {
	virt, err := mem.VirtualMemory()
	if err != nil {
		return err
	}

	j, err := json.Marshal(virt)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return errors.New("Json encoding VM stat")
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(j)

	return nil
}

func GetCPUInfo(rw http.ResponseWriter) (error) {
	cpus, err := cpu.Info()
	if err != nil {
		return err
	}

	j, err := json.Marshal(cpus)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return errors.New("Json encoding CPU Info")
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(j)

	return nil
}

func GetCPUTimeStat(rw http.ResponseWriter) (error) {
	cpus, err := cpu.Times(true)
	if err != nil {
		return err
	}

	j, err := json.Marshal(cpus)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return errors.New("Json encoding CPU stat")
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(j)

	return nil
}

func GetAvgStat(rw http.ResponseWriter) (error) {
	avgstat, r := load.Avg()
	if r != nil {
		return r
	}

	j, err := json.Marshal(avgstat)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return errors.New("Json encoding avg stat")
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(j)

	return nil
}

func GetMisc(rw http.ResponseWriter) (error) {
	lines, err := share.ReadFullFile(ProcMiscPath)
	if err != nil {
		log.Fatal("Failed to read: %s", ProcMiscPath)
		return errors.New("Failed to read misc")
	}

	miscMap := make(map[int]string)
	for _, line := range lines {
		fields := strings.Fields(line)

		deviceNum, err := strconv.Atoi(fields[0])
		if err != nil {
			continue
		}
		miscMap[deviceNum] = fields[1]
	}

	j, err := json.Marshal(miscMap)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return errors.New("Json encoding")
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(j)

	return nil
}

func GetNetArp(rw http.ResponseWriter) (error) {
	lines, err := share.ReadFullFile(ProcNetArpPath)
	if err != nil {
		log.Fatal("Failed to read: %s", ProcNetArpPath)
		return errors.New("Failed to read ProcNetArpPath")
	}

	netarp := make([]NetArp, len(lines)-1)
	for i, line := range lines {
		if i == 0 {
			continue
		}

		fields := strings.Fields(line)

		if len(fields) < 6 {
			continue
		}

		arp := NetArp{}
		for i, f := range fields {
			switch i {
			case 0:
				arp.IPAddress = f
			case 1:
				arp.HWType = f
			case 2:
				arp.Flags = f
			case 3:
				arp.HWAddress = f
			case 4:
				arp.Mask = f
			case 5:
				arp.Device = f
			}
		}
		netarp[i-1] = arp
	}

	j, err := json.Marshal(netarp)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return errors.New("Json encoding ARP")
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(j)

	return nil
}
