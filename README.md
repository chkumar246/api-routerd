# api-routerd
![N|Solid](https://ibin.co/4R6Hzr2H7l4A.png)

A RestAPI MicroService GateWay for Linux

[![Build Status CircleCI](https://circleci.com/gh/RestGW/api-routerd.svg?style=svg)](https://circleci.com/gh/RestGW/api-routerd)
[![Build Status](https://travis-ci.org/RestGW/api-routerd.svg?branch=master)](https://travis-ci.org/RestGW/api-routerd)
[![HitCount](http://hits.dwyl.io/ssahani/RestGW/api-routerd.svg)](http://hits.dwyl.io/ssahani/RestGW/api-routerd)
[![CodeFactor](https://www.codefactor.io/repository/github/restgw/api-routerd/badge)](https://www.codefactor.io/repository/github/restgw/api-routerd)
[![codebeat badge](https://codebeat.co/badges/1bdd48c6-4cc1-4255-a11b-9807473e9c3d)](https://codebeat.co/projects/github-com-restgw-api-routerd-master)
[![Go Report Card](https://goreportcard.com/badge/github.com/RestGW/api-routerd)](https://goreportcard.com/report/github.com/RestGW/api-routerd)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/5094d66b86f44522b3de7381a0bba5a1)](https://app.codacy.com/app/ssahani/api-routerd?utm_source=github.com&utm_medium=referral&utm_content=RestGW/api-routerd&utm_campaign=Badge_Grade_Dashboard)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Coverage Status](https://coveralls.io/repos/github/RestGW/api-routerd/badge.svg?branch=master)](https://coveralls.io/github/RestGW/api-routerd?branch=master)


api-routerd is a cloud-enabled, mobile-ready, a super light weight remote management tool which uses REST API for real time configuration and performance as well as health monitoring for systems (containers) and applications. It provides fast API based monitoring without affecting the system it's running on.

- Platform independent REST APIs can be accessed via any application (curl, chrome, PostMan ...) from any OS (Linux, IOS, Android, Windows ...)
- An [Iphone App Demo](https://www.linkedin.com/feed/update/urn:li:activity:6486243669560127488) using REST APIS
- Minimal data transfer using JSONoffline-storage, AngularJS powered HTML5 Markdown editor.

# Features!

- systemd socket activation support
- systemd
  - systemd informations
  - services (start, stop, restart, status)
  - service properties for example CPUShares
  - See service logs.
- networkd config
  - .network
  - .netdev
  - .link
- configure hostnamed
- configure users using login (list-sessions, list-users and terminate-user etc)
- configure timdate
- configure nameserver /etc/resolv.conf
- configure timesynd
- configure journald.conf
- configure system.conf
- configure coredump.conf
- configure systemd-resolved.conf
- configure kernel modules (modprobe, lsmod, rmmod)
- configure network (netlink)
  - Link: mtu, up, down
  - Create bridge and enslave links
  - Adddress: Set, Get, Delete
  - Gateway: Default Gateway Add and Delete
- configure group add/delete/modify
- configure users add/delete/modify (requires newuser)
- configure sysctl add/delete/modify and apply
- see information from /proc such as netstat, netdev, memory and much more
- configure /proc/sys/net (core/ipv4/ipv6), VM
- See ethtool information
- See sudoers and sshd conf

### api-routerd json API

 Refer spreadsheet [API](https://docs.google.com/spreadsheets/d/e/2PACX-1vTl2Vmp-BdTE5Vgi_PiW-qKPJnbLxdSso9kT2GAkAxCu_iWrw3_PZLlEuyXz0lbFgd7DoofXlmmb3dP/pubhtml
)

### Tech

api-routerd uses a number of open source projects to work properly:

* [logrus](github.com/sirupsen/logrus)
* [gorilla mux](github.com/gorilla/mux) 
* [netlink](github.com/vishvananda/netlink)
* [gopsutil](github.com/shirou/gopsutil)
* [coreos dbus](github.com/coreos/go-systemd/dbus)
* [dbus](github.com/godbus/dbus) 
* [ethtool](github.com/safchain/ethtool)
* [BurntSushi toml](github.com/BurntSushi/toml)
* [go-ini](github.com/go-ini/ini)


And of course Dillinger itself is open source with a [public repository][git-repo-url]
 on GitHub.

### Development

Want to contribute? Great!

### Installation

First configure your GOPATH. If you have already done this skip this step.

```sh
# keep in ~/.bashrc
```

```sh
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
export OS_OUTPUT_GOPATH=1
```

Clone inside src dir of $GOPATH. In my case

```sh
$ pwd
/home/sus/go/src
```

### Installation

```sh
$ go get github.com/RestGW/api-routerd
$ git clone https://github.com/RestGW/api-routerd
$ cd api-routerd
$ pwd
/home/sus/go/src/api-routerd
$ go build -v

$ sudo ./api-routerd
INFO[0000] api-routerd: v0.1 (built go1.11.4)
INFO[0000] Start Server at 0.0.0.0:8080
INFO[0000] Starting api-routerd in plain text mode

```

### How to configure IP and Port ?

Conf dir: ```/etc/api-routerd/```
Conf File: ```api-routerd.toml```

```sh
$ cat /etc/api-routerd/api-routerd.toml
[Network]
IPAddress="0.0.0.0"
Port="8080"
```

### How to configure users ?

Add user name and authentication string in space separated lines

```sh
# cat /etc/api-routerd/api-routerd-auth.conf
Susant secret
Max bbbb
Joy ccccc
```

### How to configure TLS ?

Generate private key (.key)

```sh
# Key considerations for algorithm "RSA" ≥ 2048-bit
$ openssl genrsa -out server.key 2048
Generating RSA private key, 2048 bit long modulus (2 primes)
.......................+++++
.+++++
e is 65537 (0x010001)

openssl genrsa -out server.key 2048
```

Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)

```sh
$ openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) [XX]:

```

Place ```server.crt``` and ```server.key``` in the dir /etc/api-routerd/tls

```sh
[root@Zeus tls]# ls
server.crt  server.key
[root@Zeus tls]# pwd
/etc/api-routerd/tls

```

Now start using https

```sh
$ curl --header "X-Session-Token: secret" --request GET https://localhost:8080/api/network/ethtool/vmnet8/get-link-features -k --tlsv1.2

{"esp-hw-offload":false,"esp-tx-csum-hw-offload":false,"fcoe-mtu":false,"highdma":false,"hw-tc-offload":false,"l2-fwd-offload":false,"loopback":false,"netns-local":false,"rx-all":false,"rx-checksum":false,"rx-fcs":false,"rx-gro":true,"rx-gro-hw":false,"rx-hashing":false,"rx-lro":false,"rx-ntuple-filter":false,"rx-udp_tunnel-port-offload":false,"rx-vlan-filter":false,"rx-vlan-hw-parse":false,"rx-vlan-stag-filter":false,"rx-vlan-stag-hw-parse":false,"tls-hw-record":false,"tls-hw-rx-offload":false,"tls-hw-tx-offload":false,"tx-checksum-fcoe-crc":false,"tx-checksum-ip-generic":false,"tx-checksum-ipv4":false,"tx-checksum-ipv6":false,"tx-checksum-sctp":false,"tx-esp-segmentation":false,"tx-fcoe-segmentation":false,"tx-generic-segmentation":false,"tx-gre-csum-segmentation":false,"tx-gre-segmentation":false,"tx-gso-partial":false,"tx-gso-robust":false,"tx-ipxip4-segmentation":false,"tx-ipxip6-segmentation":false,"tx-lockless":false,"tx-nocache-copy":false,"tx-scatter-gather":false,"tx-scatter-gather-fraglist":false,"tx-sctp-segmentation":false,"tx-tcp-ecn-segmentation":false,"tx-tcp-mangleid-segmentation":false,"tx-tcp-segmentation":false,"tx-tcp6-segmentation":false,"tx-udp-segmentation":false,"tx-udp_tnl-csum-segmentation":false,"tx-udp_tnl-segmentation":false,"tx-vlan-hw-insert":false,"tx-vlan-stag-hw-insert":false,"vlan-challenged":false}
```

```sh
$ curl --request GET --header "X-Session-Token: secret" https://localhost:8080/api/proc/misc --tlsv1.2 -k
{"130":"watchdog","144":"nvram","165":"vmmon","183":"hw_random","184":"microcode","227":"mcelog","228":"hpet","229":"fuse","231":"snapshot","232":"kvm","235":"autofs","236":"device-mapper","53":"vboxnetctl","54":"vsock","55":"vmci","56":"vboxdrvu","57":"vboxdrv","58":"rfkill","59":"memory_bandwidth","60":"network_throughput","61":"network_latency","62":"cpu_dma_latency","63":"vga_arbiter"}

$ curl --request GET --header "X-Session-Token: secret" https://localhost:8080/api/proc/net/arp --tlsv1.2 -k
[{"ip_address":"192.168.225.1","hw_type":"0x1","flags":"0x2","hw_address":"1a:89:20:38:68:8f","mask":"*","device":"wlp4s0"}]

$ curl --request GET --header "X-Session-Token: secret" https://localhost:8080/api/proc/modules --tlsv1.2 -k
```

Use case: systemd

See systemd information

```sh
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/service/systemd/version
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/service/systemd/state
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/service/systemd/features
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/service/systemd/virtualization
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/service/systemd/architecture
````

Get all units

```sh
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/service/systemd/units
```

Get all properties of a unit

```sh
$ curl --request GET --header "X-Session-Token: secret "http://localhost:8080/api/service/systemd/sshd.service/get
```

Status of a unit

```sh
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/service/systemd/sshd.service/status
```

Set and get propetties

```sh
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/service/systemd/sshd.service/get/LimitNOFILESoft
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/service/systemd/sshd.service/get/LimitNOFILE
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/service/systemd/sshd.service/get
$ curl --header "X-Session-Token: secret" --header "Content-Type: application/json" --request PUT --data '{"value":"1100"}' http://localhost:8080/api/service/systemd/sshd.service/set/CPUShares
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/service/systemd/sshd.service/get/CPUShares
$ curl --header "Content-Type: application/json" --request POST --data '{"action":"start","unit":"sshd.service"}' --header "X-Session-Token: secret" http://localhost:8080/api/service/systemd
$ curl --header "Content-Type: application/json" --request POST --data '{"action":"stop","unit":"sshd.service"}' --header "X-Session-Token: secret" http://localhost:8080/api/service/systemd
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/service/systemd/sshd.service/status
{"property":"active","unit":"sshd.service"}
```

Send a signal to the service

```sh
$curl --header "Content-Type: application/json" --request POST --data '{"action":"kill","unit":"sshd.service", "value":"9"}' --header "X-Session-Token: secret" http://localhost:8080/api/service/systemd
```

Get all unittype properties such as Service, Mount, Socket

```sh
$ curl --request GET --header "X-Session-Token: secret http://localhost:8080/api/service/systemd/sshd.service/gettype/Service
```

systemd system.conf ```/etc/systemd/system.conf```

Get
```sh
$ curl --header "Content-Type: application/json" --request POST --data '{ "DefaultLimitNOFILE":"1024"}' --header "X-Session-Token: secret" http://localhost:8080/api/service/systemd/conf
{"CPUAffinity":"","CapabilityBoundingSe":"","CrashChangeVT":"","CrashReboot":"","CrashShell":"","CtrlAltDelBurstAction":"","DefaultBlockIOAccounting":"","DefaultCPUAccounting":"","DefaultEnvironment":"","DefaultIOAccounting":"","DefaultIPAccounting":"","DefaultLimitAS":"","DefaultLimitCORE":"","DefaultLimitCPU":"","DefaultLimitDATA":"","DefaultLimitFSIZE":"","DefaultLimitLOCKS":"","DefaultLimitMEMLOCK":"","DefaultLimitMSGQUEUE":"","DefaultLimitNICE":"","DefaultLimitNOFILE":"1024","DefaultLimitNPROC":"","DefaultLimitRSS":"","DefaultLimitRTPRIO":"","DefaultLimitRTTIME":"","DefaultLimitSIGPENDING":"","DefaultLimitSTACK":"","DefaultMemoryAccounting":"","DefaultRestartSec":"","DefaultStandardError":"","DefaultStandardOutput":"","DefaultStartLimitBurst":"","DefaultStartLimitIntervalSec":"","DefaultTasksAccounting":"","DefaultTasksMax":"","DefaultTimeoutStartSec":"","DefaultTimeoutStopSec":"","DefaultTimerAccuracySec":"","DumpCore":"","IPAddressAllow":"","IPAddressDeny":"","JoinControllers":"","LogColor":"","LogLevel":"","LogLocation":"","LogTarget":"","RuntimeWatchdogSec":"","ShowStatus":"","ShutdownWatchdogSec":"","SystemCallArchitectures":"","TimerSlackNSec":""}
```

Set

```sh
$ curl --header "Content-Type: application/json" --request POST --data '{ "DefaultLimitNOFILE":"1024"}' --header "X-Session-Token: secret" http://localhost:8080/api/service/systemd/conf/update
```

Reset

```sh
$ curl --header "Content-Type: application/json" --request POST --data '{ "DefaultLimitNOFILE":""}' --header "X-Session-Token: secret" http://localhost:8080/api/service/systemd/conf/update

```

Configure journald

```sh
$ curl --header "Content-Type: application/json" --request GET --header "X-Session-Token: secret" http://localhost:8080/api/system/journal/conf
```

Use case:
- command: "GET"
  - netdev
  - version
  - vm
  - netstat
  - proto-counter-stat
  - proto-pid-stat
  - interface-stat
  - swap-memory
  - virtual-memory
  - cpuinfo
  - cputimestat
  - avgstat
  - misc
  - arp
  - modules
  - userstat
  - temperaturestat

/proc examples:

```sh
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/proc/netdev
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/proc/version
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/proc/netstat/tcp
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/proc/proto-pid-stat/2881/tcp

$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/proc/cpuinfo
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/proc/cputimestat
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/proc/virtual-memory
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/proc/swap-memory
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/proc/userstat
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/proc/temperaturestat
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/proc/proto-counter-stat
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/proc/misc
$ curl --request GET --header "X-Session-Token: secret" https://localhost:8080/api/proc/net/arp
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/proc/modules
```

Information by pid

GET
```sh
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/proc/process/2769/pid-connections/
[{"fd":270,"family":2,"type":1,"localaddr":{"ip":"192.168.225.101","port":45354},"remoteaddr":{"ip":"74.125.24.102","port":443},"status":"ESTABLISHED","uids":[1000,1000,1000,1000],"pid":2769},{"fd":196,"family":2,"type":1,"localaddr":{"ip":"192.168.225.101","port":49138},"remoteaddr":{"ip":"172.217.194.94","port":443},"status":"ESTABLISHED","uids":[1000,1000,1000,1000],"pid":27

$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/proc/process/2769/pid-rlimit/
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/proc/process/2769/pid-rlimit-usage/
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/proc/process/2769/pid-status/
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/proc/process/2769/pid-username/
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/proc/process/2769/pid-open-files/
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/proc/process/2769/pid-fds/
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/proc/process/2769/pid-name/
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/proc/process/2769/pid-memory-percent/
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/proc/process/2769/pid-memory-maps/
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/proc/process/2769/pid-memory-info/
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/proc/process/2769/pid-io-counters/
```

```sh
# curl --header --request GET  --header "X-Session-Token: secret" http://localhost:8080/api/proc/misc
{"130":"watchdog","144":"nvram","165":"vmmon","183":"hw_random","184":"microcode","227":"mcelog","228":"hpet","229":"fuse","231":"snapshot","232":"kvm","235":"autofs","236":"device-mapper","53":"vboxnetctl","54":"vsock","55":"vmci","56":"vboxdrvu","57":"vboxdrv","58":"rfkill","59":"memory_bandwidth","60":"network_throughput","61":"network_latency","62":"cpu_dma_latency","63":"vga_arbiter"}
```

proc vm: property any file name in /proc/sys/vm/

```sh
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/proc/sys/vm/swappiness
{"property":"swappiness","value":"60"}
$ curl --header "Content-Type: application/json" --request PUT --data '{"value":"70"}' --header "X-Session-Token: secret" http://localhost:8080/api/proc/sys/vm/swappiness
{"property":"swappiness","value":"70"}

```

```sh
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/proc/version
{"hostname":"Zeus","uptime":17747,"bootTime":1545381768,"procs":360,"os":"linux","platform":"fedora","platformFamily":"fedora","platformVersion":"29","kernelVersion":"4.19.2-300.fc29.x86_64","virtualizationSystem":"kvm","virtualizationRole":"host","hostid":"27f7c64c-3148-11b2-a85c-ec64a5733ce1"}

```

set and get any value from ```/proc/sys/net```.
supported: IPv4, IPv6 and core

```sh
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/proc/sys/net/ipv4/enp0s31f6/forwarding
{"path":"ipv4","property":"forwarding","value":"0","link":"enp0s31f6"}
$  curl --header "Content-Type: application/json" --request PUT --data '{"value":"1"}' --header "X-Session-Token: secret" http://localhost:8080/api/proc/sys/net/ipv4/enp0s31f6/forwarding
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/proc/sys/net/ipv4/enp0s31f6/forwarding
{"path":"ipv4","property":"forwarding","value":"1","link":"enp0s31f6"}

```

###### Use case configure link

Add address

```sh
$ curl --header "Content-Type: application/json" --request POST --data '{"action":"add-address", "address":"192.168.1.131/24", "link":"dummy"}' --header "X-Session-Token: secret" http://localhost:8080/api/network/address/add
$ ip addr show dummy
9: dummy: <BROADCAST,NOARP> mtu 1500 qdisc noop state DOWN group default qlen 1000
    link/ether ea:d0:e3:be:ea:25 brd ff:ff:ff:ff:ff:ff
    inet 192.168.1.131/24 brd 192.168.1.255 scope global dummy
       valid_lft forever preferred_lft forever

```

Set link up/down

```sh
$ curl --header "Content-Type: application/json" --request PUT --data '{"action":"set-link-up", "link":"dummy"}' --header "X-Session-Token: secret" http://localhost:8080/api/network/link/set

$ curl --header "Content-Type: application/json" --request PUT --data '{"action":"set-link-down", "link":"dummy"}' --header "X-Session-Token: secret" http://localhost:8080/api/network/link/set
$ ip addr show dummy
9: dummy: <BROADCAST,NOARP> mtu 1500 qdisc noqueue state DOWN group default qlen 1000
    link/ether ea:d0:e3:be:ea:25 brd ff:ff:ff:ff:ff:ff
    inet 192.168.1.131/24 brd 192.168.1.255 scope global dummy
       valid_lft forever preferred_lft forever
```

Set MTU

 ```sh
$curl --header "Content-Type: application/json" --request PUT --data '{"action":"set-link-mtu", "link":"dummy", "mtu":"1280"}' http://localhost:8080/api/network/link/set
$ ip addr show dummy
9: dummy: <BROADCAST,NOARP> mtu 1280 qdisc noqueue state DOWN group default qlen 1000
    link/ether ea:d0:e3:be:ea:25 brd ff:ff:ff:ff:ff:ff
    inet 192.168.1.131/24 brd 192.168.1.255 scope global dummy
       valid_lft forever preferred_lft forever
```

Set Default GateWay

```sh
 $ curl --header "Content-Type: application/json" --request POST --data '{"action":"add-default-gw", "link":"dummy", "gateway":"192.168.1.1/24", "onlink":"true"}' --header "X-Session-Token: secret" http://localhost:8080/api/network/route/add
$ ip route
default via 192.168.1.1 dev dummy onlink
```

Create a bridge and enslave two links

```sh
$ curl --header "Content-Type: application/json" --request POST --data '{"action":"add-link-bridge", "link":"test-br", "enslave":["dummy","dummy1"]}' --header "X-Session-Token: secret" http://localhost:8080/api/network/link/add

# ip link
9: dummy: <BROADCAST,NOARP> mtu 12801 qdisc noop master test-br state DOWN mode DEFAULT group default qlen 1000
    link/ether f2:58:ea:f3:83:1e brd ff:ff:ff:ff:ff:ff
10: dummy1: <BROADCAST,NOARP> mtu 1500 qdisc noop master test-br state DOWN mode DEFAULT group default qlen 1000
    link/ether 12:00:9a:65:36:6d brd ff:ff:ff:ff:ff:ff
11: test-br: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN mode DEFAULT group default
    link/ether 12:00:9a:65:36:6d brd ff:ff:ff:ff:ff:ff

```

Delete a link

```sh
$ curl --header "Content-Type: application/json" --request DELETE --data '{"action":"delete-link", "link":"test-br"}' --header "X-Session-Token: secret" http://localhost:8080/api/network/link/delete
```

###### Use Case: networkd

```sh
$ curl --header "Content-Type: application/json" --request POST --data '{"Match": [{"Name":"eth0"}], "DHCP":"yes", "LLDP":"yes","Addresses": [{"Address":"192.168.1est1"},{"Address":"192.168.1.4", "Label":"test3", "Peer":"192.168.1.5"}], "Routes": [{"Gateway":"192.168.1.1","GatewayOnlink":"true"},{"Destination":"192.168.1.10","Table":"10"}]}' --header "X-Session-Token: secret" http://localhost:8080/api/network/networkd/network

$ cat /etc/systemd/network/25-eth0.network
[Match]
Name=eth0

[Network]
DHCP=yes
LLDP=yes


[Address]
Address=192.168.1.2
Label=test1

[Address]
Address=192.168.1.4
Peer=192.168.1.5
Label=test3


[Route]
Gateway=192.168.1.1
GatewayOnlink=yes

[Route]
Destination=192.168.1.10
Table=10

```

RoutingPolicyRule

```sh
$  curl --header "Content-Type: application/json" --request POST --data '{"Match": [{"Name":"eth0"}], "DHCP":"yes", "LLDP":"yes","RoutingPolicyRule": [{"TypeOfService":"1"},{"From":"192.168.1.4", "Table":"3", "Priority":"5"}], "Routes": [{"Gateway":"192.168.1.1","GatewayOnlink":"true"},{"Destination":"192.168.1.10","Table":"10"}]}' --header "X-Session-Token: secret" http://localhost:8080/api/network/networkd/network


$ cat 25-eth0.network
[Match]
Name=eth0

[Network]
DHCP=yes
LLDP=yes



[Route]
Gateway=192.168.1.1
GatewayOnlink=yes

[Route]
Destination=192.168.1.10
Table=10


[RoutingPolicyRule]
TypeOfService=1
From=192.168.1.4
Table=3
Priority=5

```

[DHCP] Section

```sh
$  curl --header "Content-Type: application/json" --request POST --data '{"Match": [{"Name":"eth0"}], "DHCP":"yes", "LLDP":"yes","DHCPSection": [{"UseDNS":"yes"},{"UseMTU":"yes", "CriticalConnection":"yes", "UseRoutes":"yes"}], "Routes": [{"Gateway":"192.168.1.1","GatewayOnlink":"true"},{"Destination":"192.168.1.10","Table":"10"}]}' --header "X-Session-Token: secret" http://localhost:8080/api/network/networkd/network

```

networkd NetDev

```sh
$ curl --header "Content-Type: application/json" --request POST --data '{"Name":"bond-test", "Description":"testing bond", "Kind":"bond", "Mode":"balance-rr"}' --header "X-Session-Token: secret" http://localhost:8080/api/network/networkd/netdev

# cat /etc/systemd/network/25-bond-test.netdev
[NetDev]
Name=bond-test
Description=testing bond
Kind=bond

[Bond]
Mode=balance-rr

```

networkd Link

```sh
$ curl --header "Content-Type: application/json" --request POST --data '{"Match": [{"MAC":"00:50:56:c0:00:08"}], "Name":"test","Description":"testing link", "WakeOnLan":"phy", "TCPSegmentationOffload":"yes"}' --header "X-Session-Token: secret" http://localhost:8080/api/network/networkd/link

sus@Zeus api-routerd]$
# pwd
/etc/systemd/network
# cat 00-test.link
[Match]
MACAddress=00:50:56:c0:00:08

[Link]
Description=testing link
Name=test
WakeOnLan=phy
TCPSegmentationOffload=yes
```

Bridge

```sh

$ curl --header "Content-Type: application/json" --request POST --data '{"Name":"bridge-test", "Description":"testing bridge", "Kind":"bridge", "HelloTimeSec":"30"}' --header "X-Session-Token: secret" http://localhost:8080/api/network/networkd/netdev
$ cat /etc/systemd/network/25-bridge-test.netdev
[NetDev]
Name=bridge-test
Description=testing bridge
Kind=bridge

[Bridge]
HelloTimeSec =30

$ curl --header "Content-Type: application/json" --request POST --data '{"Name":"eth0", "Description":"etho bridge enslave", "Bridge":"bridge-test"}' http://localhost:8080/api/network/networkd/network
$ cat /etc/systemd/network/25-eth0.network
[Match]
Name=eth0

[Network]
Bridge=bridge-test
```

Tunnel

```sh
$ curl --header "Content-Type: application/json" --request POST --data '{"Name":"tunnel-test", "Description":"testing tunnel", "Kind":"tunnel", "Local":"192.168.1.2", "Remote":"192.168.1.2", "Independent":"true"}' --header "X-Session-Token: secret" http://localhost:8080/api/network/networkd/netdev
$ cat /etc/systemd/network/
00-.link               00-test.link           25-bond-test.netdev    25-tunnel-test.netdev
$ cat /etc/systemd/network/25-tunnel-test.netdev
[NetDev]
Name=tunnel-test
Description=testing tunnel
Kind=tunnel

[Tunnel]
Local=192.168.1.2
Remote=192.168.1.2
Independent=true
```
Vxlan

```sh
$ curl --header "Content-Type: application/json" --request POST --data '{"Name":"vxlan-test", "Description":"testing vxlan", "Kind":"vxlan", "Local":"192.168.1.2", "Remote":"192.168.1.2", "Id":"21"}' --header "X-Session-Token: secret" http://localhost:8080/api/network/networkd/netdev

# cat 25-vxlan-test.netdev
[NetDev]
Name=vxlan-test
Description=testing vxlan
Kind=vxlan

[VXLAN]
Id=21
Local=192.168.1.2
Remote=192.168.1.2
```

Veth

```sh
$ curl --header "Content-Type: application/json" --request POST --data '{"Name":"veth-test", "Description":"testing veth", "Kind":"veth", "Peer":"test"}' --header "X-Session-Token: secret" http://localhost:8080/api/network/networkd/netdev
# cat 25-veth-test.netdev
[NetDev]
Name=veth-test
Description=testing veth
Kind=veth

[Peer]
Name=test
```

###### Hostname

Example: Get and Set Hostname

```sh
$ curl --request GET --header "X-Session-Token: secret" http://localhost:8080/api/system/hostname
{"Chassis":"laptop","Deployment":"","HomeURL":"https://fedoraproject.org/","Hostname":"Zeus","IconName":"computer-laptop","KernelName":"Linux","KernelRelease":"4.19.2-300.fc29.x86_64","KernelVersion":"#1 SMP Wed Nov 14 19:05:24 UTC 2018","Location":"","OperatingSystemCPEName":"cpe:/o:fedoraproject:fedora:29","OperatingSystemPrettyName":"Fedora 29 (Twenty Nine)","PrettyHostname":"","StaticHostname":"Zeus"}

$ curl --header "Content-Type: application/json" --request PUT --data '{"property":"SetStaticHostname","value":"Zeus1"}' --header "X-Session-Token: secret" http://localhost:8080/api/system/hostname/set
$ hostname
Zeus1
$ curl --header "Content-Type: application/json" --request PUT --data '{"property":"SetStaticHostname","value":"Zeus"}' --header "X-Session-Token: secret" http://localhost:8080/api/system/hostname/set
$ hostname
Zeus
```

Supported Property for hostname

```sh
        "Hostname"
        "StaticHostname"
        "PrettyHostname"
        "IconName"
        "Chassis"
        "Deployment"
        "Location"
        "KernelName"
        "KernelRelease"
        "KernelVersion",
        "OperatingSystemPrettyName"
        "OperatingSystemCPEName"
        "HomeURL"
}
```

Supported Property (Methods) for setting hostname. For example: ```'{"property":"SetStaticHostname","value":"Zeus"}'```

```sh
        "SetHostname"
        "SetStaticHostname"
        "SetPrettyHostname"
        "SetIconName"
        "SetChassis"
        "SetDeployment"
        "SetLocation"
```

###### systemd-logind

Supported methods

```sh
        list-sessions
        list-users
        lock-session
        lock-sessions
        terminate-session
        terminate-user
```

example:

```sh
$ curl --header "Content-Type: application/json" --request GET --header "X-Session-Token: secret" http://localhost:8080/api/system/login/get/list-sessions
$ curl --header "Content-Type: application/json" --request GET --header "X-Session-Token: secret" http://localhost:8080/api/system/login/get/list-users
$ curl --header "Content-Type: application/json" --request POST --data '{"value":"1002"}' --header "X-Session-Token: secret" http://localhost:8080/api/system/login/post/terminate-user
```

###### ethtool

```sh
$ curl --header "X-Session-Token: secret" --request GET  http://localhost:8080/api/network/ethtool/vmnet8/get-link-features

{"esp-hw-offload":false,"esp-tx-csum-hw-offload":false,"fcoe-mtu":false,"highdma":false,"hw-tc-offload":false,"l2-fwd-offload":false,"loopback":false,"netns-local":false,"rx-all":false,"rx-checksum":false,"rx-fcs":false,"rx-gro":true,"rx-gro-hw":false,"rx-hashing":false,"rx-lro":false,"rx-ntuple-filter":false,"rx-udp_tunnel-port-offload":false,"rx-vlan-filter":false,"rx-vlan-hw-parse":false,"rx-vlan-stag-filter":false,"rx-vlan-stag-hw-parse":false,"tls-hw-record":false,"tls-hw-rx-offload":false,"tls-hw-tx-offload":false,"tx-checksum-fcoe-crc":false,"tx-checksum-ip-generic":false,"tx-checksum-ipv4":false,"tx-checksum-ipv6":false,"tx-checksum-sctp":false,"tx-esp-segmentation":false,"tx-fcoe-segmentation":false,"tx-generic-segmentation":false,"tx-gre-csum-segmentation":false,"tx-gre-segmentation":false,"tx-gso-partial":false,"tx-gso-robust":false,"tx-ipxip4-segmentation":false,"tx-ipxip6-segmentation":false,"tx-lockless":false,"tx-nocache-copy":false,"tx-scatter-gather":false,"tx-scatter-gather-fraglist":false,"tx-sctp-segmentation":false,"tx-tcp-ecn-segmentation":false,"tx-tcp-mangleid-segmentation":false,"tx-tcp-segmentation":false,"tx-tcp6-segmentation":false,"tx-udp-segmentation":false,"tx-udp_tnl-csum-segmentation":false,"tx-udp_tnl-segmentation":false,"tx-vlan-hw-insert":false,"tx-vlan-stag-hw-insert":false,"vlan-challenged":false}
```

```sh
$ curl --header "X-Session-Token: secret" --request GET http://localhost:8080/api/network/ethtool/wlp4s0/get-link-stat
{"ch_time":18446744073709551615,"ch_time_busy":18446744073709551615,"ch_time_ext_busy":18446744073709551615,"ch_time_rx":18446744073709551615,"ch_time_tx":18446744073709551615,"channel":0,"noise":18446744073709551615,"rx_bytes":1387313,"rx_dropped":45,"rx_duplicates":0,"rx_fragments":3255,"rx_packets":3275,"rxrate":117000000,"signal":227,"sta_state":4,"tx_bytes":584843,"tx_filtered":0,"tx_packets":2949,"tx_retries":321,"tx_retry_failed":0,"txrate":144400000}
```

```sh
$  curl --header "X-Session-Token: secret" --request GET http://localhost:8080/api/network/ethtool/wlp4s0/get-link-stat
{"ch_time":18446744073709551615,"ch_time_busy":18446744073709551615,"ch_time_ext_busy":18446744073709551615,"ch_time_rx":18446744073709551615,"ch_time_tx":18446744073709551615,"channel":0,"noise":18446744073709551615,"rx_bytes":1387313,"rx_dropped":45,"rx_duplicates":0,"rx_fragments":3255,"rx_packets":3275,"rxrate":117000000,"signal":227,"sta_state":4,"tx_bytes":584843,"tx_filtered":0,"tx_packets":2949,"tx_retries":321,"tx_retry_failed":0,"txrate":144400000}
```

```sh
$  curl --header "X-Session-Token: secret" --request GET http://localhost:8080/api/network/ethtool/wlp4s0/get-link-driver-name
{"action":"get-link-driver-name","link":"wlp4s0","reply":"iwlwifi"}

$  curl --header "X-Session-Token: secret" --request GET  http://localhost:8080/api/network/ethtool/vmnet8/get-link-driver-info
$  curl --header "X-Session-Token: secret" --request GET  http://localhost:8080/api/network/ethtool/vmnet8/get-link-permaddr
$  curl --header "X-Session-Token: secret" --request GET  http://localhost:8080/api/network/ethtool/vmnet8/get-link-eeprom
$  curl --header "X-Session-Token: secret" --request GET  http://localhost:8080/api/network/ethtool/vmnet8/get-link-msglvl
$  curl --header "X-Session-Token: secret" --request GET  http://localhost:8080/api/network/ethtool/vmnet8/get-link-mapped
```

###### Netlink

Get link netlink

```sh
$ curl --header "X-Session-Token: secret" --request GET http://localhost:8080/api/network/link/get/wlp4s0
{"index":5,"MTU":1500,"TxQLen":0,"Name":"wlp4s0","HardwareAdd":"7c:76:35:ea:89:90","LinkOperState":""}

$  curl --header "X-Session-Token: secret" --request GET http://localhost:8080/api/network/address/get/wlp4s0
[{"action":"","link":"wlp4s0","address":"192.168.43.105/24","label":""},{"action":"","link":"wlp4s0","address":"2409:4042:239c:7f9d:e45f:27a9:c6de:c39e/64","label":""},{"action":"","link":"wlp4s0","address":"fe80::c912:39ce:e9a3:aaca/64","label":""}]
```

get all links

```sh
$ curl --header "X-Session-Token: secret" --request GET http://localhost:8080/api/network/link/get
```

get all routes

```sh
$ curl --header "X-Session-Token: secret" --request GET http://localhost:8080/api/network/route/get
```

Replace route

```sh
$ curl --header "Content-Type: application/json" --request PUT --data '{"action":"replace-default-gw", "link":"dummy", "gateway":"192.168.1.3/24", "onlink":"true"}' --header "X-Session-Token: secret" http://localhost:8080/api/network/route/add
$ ip route
default via 192.168.1.3 dev dummy onlink
default via 192.168.225.1 dev wlp4s0 proto dhcp metric 600
172.16.130.0/24 dev vmnet8 proto kernel scope link src 172.16.130.1
192.168.1.0/24 dev dummy proto kernel scope link src 192.168.1.131
192.168.137.0/24 dev vmnet1 proto kernel scope link src 192.168.137.1
192.168.225.0/24 dev wlp4s0 proto kernel scope link src 192.168.225.101 metric 600
```

Get all addresses

```sh
http://localhost:8080/api/network/address/get
```

###### Add/Read/Delete confs in ```resolv.conf```

Get

```sh
$ http://localhost:8080/api/system/resolv
```

Add

``` sh

$ curl --header "Content-Type: application/json" --request POST --data '{"servers":["192.168.1.131","192.168.1.132"], "search":["hello","hello2"]}' --header "X-Session-Token: secret" http://localhost:8080/api/system/resolv/add
```

Delete

```sh
$ curl --header "Content-Type: application/json" --request DELETE --data '{"servers":["192.168.225.3","192.168.225.2"]}' --header "X-Session-Token: secret" http://localhost:8080/api/network/system/delete
```

###### Configure systemd-resolved

```sh
To get
http://localhost:8080/api/system/systemdresolved

Add
$  curl --header "Content-Type: application/json" --request POST --data '{"dns":["192.168.1.131","192.168.1.132"], "fallback_dns":["hello","hello2"]}' --header "X-Session-Token: secret" http://localhost:8080/api/system/systemdresolved/add
{"dns":["10.68.5.26 10.64.63.6 192.168.225.1","192.168.1.131","192.168.1.132"],"fallback_dns":["8.8.8.8 8.8.4.4 2001:4860:4860::8888 2001:4860:4860::8844","hello","hello2"]}

Delete
$  curl --header "Content-Type: application/json" --request DELETE --data '{"dns":["192.168.1.131"]}' --header "X-Session-Token: secret" http://localhost:8080/api/system/systemdresolved/delete
{"dns":["10.68.5.26","10.64.63.6","192.168.225.1"],"fallback_dns":["8.8.8.8","8.8.4.4","2001:4860:4860::8888","2001:4860:4860::8844"]}
```

###### configure coredump

```sh
$  curl --header "Content-Type: application/json" --request GET http://localhost:8080/api/system/coredump --header "X-Session-Token: secret"
$  curl --header "Content-Type: application/json" --request POST --data '{"Storage":"internal"}' --header "X-Session-Token: secret" http://localhost:8080/api/system/coredump/add
$  curl --header "Content-Type: application/json" --request DELETE --data '{"Storage":"#"}' --header "X-Session-Token: secret" http://localhost:8080/api/system/coredump/delete
```
###### kmod

```sh
curl --header "Content-Type: application/json" --request POST --data '{"name":"ipip"}' --header "X-Session-Token: secret" http://localhost:8080/api/system/kmod/modprobe
```

###### group add/delete/modify

```sh
$  curl --header "X-Session-Token: secret" --header "Content-Type: application/json" --request POST --data '{"name":"test", "gid":"1111"}' http://localhost:8080/api/system/group/add
$  curl --header "X-Session-Token: secret" --header "Content-Type: application/json" --request DELETE --data '{"name":"test"}' http://localhost:8080/api/system/group/delete
$  curl --header "X-Session-Token: secret" --header "Content-Type: application/json" --request PUT --data '{"name":"test1", "new_name":"test"}' http://localhost:8080/api/system/group/modify

```

###### user add/delete/modify

``` sh
$  curl --header "X-Session-Token: secret" --header "Content-Type: application/json" --request POST --data '{"username":"testuser3", "password":"password@321", "home_die":"/home/testuser3"}' http://localhost:8080/api/system/user/add
$  curl --header "X-Session-Token: secret" --header "Content-Type: application/json" --request DELETE --data '{"username":"testuser3"}' http://localhost:8080/api/system/user/delete
$  curl --header "X-Session-Token: secret" --header "Content-Type: application/json" --request PUT --data '{"username":"testuser3", "groups": ["1004"]}' http://localhost:8080/api/system/user/modify
```

###### sysctl

```sh
$ curl --header "Content-Type: application/json" --request POST --data '{"key":"net.ipv4.conf.all.rp_filter", "value":"0", "apply":"yes"}' --header "X-Session-Token: secret" http://localhost:8080/api/system/sysctl/add
$ curl --header "Content-Type: application/json" --request POST --data '{"key":"net.ipv4.conf.all.rp_filter", "value":"1"}' --header "X-Session-Token: secret" http://localhost:8080/api/system/sysctl/modify
$ curl --header "Content-Type: application/json" --request DELETE --data '{"key":"net.ipv4.conf.all.rp_filter", "value":"0"}' --header "X-Session-Token: secret" http://localhost:8080/api/system/sysctl/delete
```


### Todos

 - Write Tests
 - Networkd
 - iptables

License
----

Apchehe 2.0


**Free Software, Hell Yeah!**

[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)

   [git-repo-url]: <git@github.com:RestGW/api-routerd.git>

