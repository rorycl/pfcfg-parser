// This part of the package reads pfsense configuration files from XML
// into a Go struct. The main struct was generated using zek :
// https://github.com/miku/zek. The BoolIfElementPresent is to deal with
// some "empty" tags; ditto for CustomTime. A few fields were renamed to
// pass linting.

package main

import (
	"encoding/xml"
	"strconv"
	"time"
)

// BoolIfElementPresent idea from
// https://stackoverflow.com/questions/23724591/golang-unmarshal-self-closing-tags
type BoolIfElementPresent bool

// UnmarshalXML unmarshalls an empty XML bool tag
func (c *BoolIfElementPresent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	*c = true
	return nil
}

// CustomTime unmarshals from a unix epoch time
type CustomTime struct {
	time.Time
}

// UnmarshalXML for CustomTime marshals
func (c *CustomTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var err error
	var v string
	d.DecodeElement(&v, &start)
	if v == "" {
		return nil
	}
	vi, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return err
	}
	ut := time.Unix(vi, 0)
	if err != nil {
		return err
	}
	*c = CustomTime{ut}
	return nil
}

// PFSenseConfig was generated using zek
type PFSenseConfig struct {
	XMLName    xml.Name `xml:"pfsense"`
	Text       string   `xml:",chardata"`
	Version    string   `xml:"version"` // 21.7
	Lastchange string   `xml:"lastchange"`
	Theme      string   `xml:"theme"` // pfsense_ng
	System     struct {
		Text         string `xml:",chardata"`
		Optimization string `xml:"optimization"`
		Hostname     string `xml:"hostname"`
		Domain       string `xml:"domain"`
		Group        []struct {
			Text        string   `xml:",chardata"`
			Name        string   `xml:"name"`
			Description string   `xml:"description"`
			Scope       string   `xml:"scope"`
			Gid         string   `xml:"gid"`
			Member      []string `xml:"member"`
			Priv        []string `xml:"priv"`
		} `xml:"group"`
		User []struct {
			Text             string `xml:",chardata"`
			Name             string `xml:"name"`
			Descr            string `xml:"descr"`
			Scope            string `xml:"scope"`
			Groupname        string `xml:"groupname"`
			UID              string `xml:"uid"`
			Priv             string `xml:"priv"`
			Expires          string `xml:"expires"`
			Authorizedkeys   string `xml:"authorizedkeys"`
			Ipsecpsk         string `xml:"ipsecpsk"`
			BcryptHash       string `xml:"bcrypt-hash"`
			Dashboardcolumns string `xml:"dashboardcolumns"`
			Webguicss        string `xml:"webguicss"`
		} `xml:"user"`
		Nextuid            string `xml:"nextuid"`
		Nextgid            string `xml:"nextgid"`
		Timezone           string `xml:"timezone"`
		TimeUpdateInterval string `xml:"time-update-interval"`
		Timeservers        string `xml:"timeservers"`
		Webgui             struct {
			Text              string `xml:",chardata"`
			Protocol          string `xml:"protocol"`
			Loginautocomplete string `xml:"loginautocomplete"`
			SslCertref        string `xml:"ssl-certref"`
			Port              string `xml:"port"`
			MaxProcs          string `xml:"max_procs"`
			Dashboardcolumns  string `xml:"dashboardcolumns"`
			Webguicss         string `xml:"webguicss"`
			Logincss          string `xml:"logincss"`
		} `xml:"webgui"`
		Disablenatreflection          string `xml:"disablenatreflection"`
		Disablesegmentationoffloading string `xml:"disablesegmentationoffloading"`
		Disablelargereceiveoffloading string `xml:"disablelargereceiveoffloading"`
		Ipv6allow                     string `xml:"ipv6allow"`
		PowerdAcMode                  string `xml:"powerd_ac_mode"`
		PowerdBatteryMode             string `xml:"powerd_battery_mode"`
		PowerdNormalMode              string `xml:"powerd_normal_mode"`
		Bogons                        struct {
			Text     string `xml:",chardata"`
			Interval string `xml:"interval"`
		} `xml:"bogons"`
		Language                string   `xml:"language"`
		DNS1gw                  string   `xml:"dns1gw"`
		DNS2gw                  string   `xml:"dns2gw"`
		DNS3gw                  string   `xml:"dns3gw"`
		DNS4gw                  string   `xml:"dns4gw"`
		Serialspeed             string   `xml:"serialspeed"`
		Primaryconsole          string   `xml:"primaryconsole"`
		AlreadyRunConfigUpgrade string   `xml:"already_run_config_upgrade"`
		CryptoHardware          string   `xml:"crypto_hardware"`
		Maximumtableentries     string   `xml:"maximumtableentries"`
		Dnsserver               []string `xml:"dnsserver"`
		SSH                     struct {
			Text   string `xml:",chardata"`
			Enable string `xml:"enable"`
		} `xml:"ssh"`
		HnAltqEnable string `xml:"hn_altq_enable"`
	} `xml:"system"`
	Interfaces struct {
		Text string `xml:",chardata"`
		Wan  struct {
			Text     string `xml:",chardata"`
			Enable   string `xml:"enable"`
			If       string `xml:"if"`
			Descr    string `xml:"descr"`
			Spoofmac string `xml:"spoofmac"`
			Ipaddr   string `xml:"ipaddr"`
			Subnet   string `xml:"subnet"`
			Gateway  string `xml:"gateway"`
		} `xml:"wan"`
		Lan struct {
			Text      string `xml:",chardata"`
			If        string `xml:"if"`
			Media     string `xml:"media"`
			Mediaopt  string `xml:"mediaopt"`
			Descr     string `xml:"descr"`
			Spoofmac  string `xml:"spoofmac"`
			Enable    string `xml:"enable"`
			Ipaddr    string `xml:"ipaddr"`
			Subnet    string `xml:"subnet"`
			Gateway   string `xml:"gateway"`
			Ipaddrv6  string `xml:"ipaddrv6"`
			Subnetv6  string `xml:"subnetv6"`
			Gatewayv6 string `xml:"gatewayv6"`
		} `xml:"lan"`
		Opt1 struct {
			Text     string `xml:",chardata"`
			If       string `xml:"if"`
			Descr    string `xml:"descr"`
			Spoofmac string `xml:"spoofmac"`
			Media    string `xml:"media"`
			Mediaopt string `xml:"mediaopt"`
			Enable   string `xml:"enable"`
			Ipaddr   string `xml:"ipaddr"`
			Subnet   string `xml:"subnet"`
		} `xml:"opt1"`
	} `xml:"interfaces"`
	Staticroutes string `xml:"staticroutes"`
	Dhcpd        struct {
		Text string `xml:",chardata"`
		Lan  struct {
			Text  string `xml:",chardata"`
			Range struct {
				Text string `xml:",chardata"`
				From string `xml:"from"`
				To   string `xml:"to"`
			} `xml:"range"`
			FailoverPeerip       string `xml:"failover_peerip"`
			Dhcpleaseinlocaltime string `xml:"dhcpleaseinlocaltime"`
			Defaultleasetime     string `xml:"defaultleasetime"`
			Maxleasetime         string `xml:"maxleasetime"`
			Netmask              string `xml:"netmask"`
			Gateway              string `xml:"gateway"`
			Domain               string `xml:"domain"`
			Domainsearchlist     string `xml:"domainsearchlist"`
			Ddnsdomain           string `xml:"ddnsdomain"`
			Ddnsdomainprimary    string `xml:"ddnsdomainprimary"`
			Ddnsdomainkeyname    string `xml:"ddnsdomainkeyname"`
			Ddnsdomainkey        string `xml:"ddnsdomainkey"`
			MacAllow             string `xml:"mac_allow"`
			MacDeny              string `xml:"mac_deny"`
			Tftp                 string `xml:"tftp"`
			Ldap                 string `xml:"ldap"`
			Nextserver           string `xml:"nextserver"`
			Filename             string `xml:"filename"`
			Filename32           string `xml:"filename32"`
			Filename64           string `xml:"filename64"`
			Rootpath             string `xml:"rootpath"`
			Numberoptions        string `xml:"numberoptions"`
		} `xml:"lan"`
	} `xml:"dhcpd"`
	Snmpd struct {
		Text        string `xml:",chardata"`
		Syslocation string `xml:"syslocation"`
		Syscontact  string `xml:"syscontact"`
		Rocommunity string `xml:"rocommunity"`
		Modules     struct {
			Text     string `xml:",chardata"`
			Mibii    string `xml:"mibii"`
			Netgraph string `xml:"netgraph"`
			Pf       string `xml:"pf"`
			Hostres  string `xml:"hostres"`
		} `xml:"modules"`
		Enable         string `xml:"enable"`
		Pollport       string `xml:"pollport"`
		Trapserver     string `xml:"trapserver"`
		Trapserverport string `xml:"trapserverport"`
		Trapstring     string `xml:"trapstring"`
		Bindip         string `xml:"bindip"`
	} `xml:"snmpd"`
	Diag struct {
		Text    string `xml:",chardata"`
		Ipv6nat struct {
			Text   string `xml:",chardata"`
			Ipaddr string `xml:"ipaddr"`
		} `xml:"ipv6nat"`
	} `xml:"diag"`
	Bridge string `xml:"bridge"`
	Syslog struct {
		Text     string `xml:",chardata"`
		Nentries string `xml:"nentries"`
		Sourceip string `xml:"sourceip"`
		Ipproto  string `xml:"ipproto"`
		Filter   string `xml:"filter"`
		System   string `xml:"system"`
	} `xml:"syslog"`
	Nat struct {
		Text     string `xml:",chardata"`
		Outbound struct {
			Text string `xml:",chardata"`
			Mode string `xml:"mode"`
		} `xml:"outbound"`
		Onetoone []struct {
			Text      string `xml:",chardata"`
			External  string `xml:"external"`
			Descr     string `xml:"descr"`
			Interface string `xml:"interface"`
			Source    struct {
				Text    string `xml:",chardata"`
				Address string `xml:"address"`
			} `xml:"source"`
			Destination struct {
				Text string `xml:",chardata"`
				Any  string `xml:"any"`
			} `xml:"destination"`
		} `xml:"onetoone"`
	} `xml:"nat"`
	Filter struct {
		Text string `xml:",chardata"`
		Rule []struct {
			Text         string `xml:",chardata"`
			ID           string `xml:"id"`
			Tracker      string `xml:"tracker"`
			Type         string `xml:"type"`
			Interface    string `xml:"interface"`
			Ipprotocol   string `xml:"ipprotocol"`
			Tag          string `xml:"tag"`
			Tagged       string `xml:"tagged"`
			Max          string `xml:"max"`
			MaxSrcNodes  string `xml:"max-src-nodes"`
			MaxSrcConn   string `xml:"max-src-conn"`
			MaxSrcStates string `xml:"max-src-states"`
			Statetimeout string `xml:"statetimeout"`
			Statetype    string `xml:"statetype"`
			Os           string `xml:"os"`
			Source       struct {
				Text    string               `xml:",chardata"`
				Address string               `xml:"address"`
				Any     BoolIfElementPresent `xml:"any"`
			} `xml:"source"`
			Destination struct {
				Text    string               `xml:",chardata"`
				Any     BoolIfElementPresent `xml:"any"`
				Port    string               `xml:"port"`
				Address string               `xml:"address"`
			} `xml:"destination"`
			Descr   string `xml:"descr"`
			Created struct {
				Text     string `xml:",chardata"`
				Time     string `xml:"time"`
				Username string `xml:"username"`
			} `xml:"created"`
			Updated struct {
				Text     string     `xml:",chardata"`
				Time     CustomTime `xml:"time"`
				Username string     `xml:"username"`
			} `xml:"updated"`
			Protocol string `xml:"protocol"`
			Disabled string `xml:"disabled"`
		} `xml:"rule"`
		Separator struct {
			Text string `xml:",chardata"`
			Wan  string `xml:"wan"`
		} `xml:"separator"`
	} `xml:"filter"`
	Shaper string `xml:"shaper"`
	Ipsec  struct {
		Text    string `xml:",chardata"`
		Vtimaps string `xml:"vtimaps"`
	} `xml:"ipsec"`
	Aliases struct {
		Text  string `xml:",chardata"`
		Alias []struct {
			Text    string `xml:",chardata"`
			Name    string `xml:"name"`
			Address string `xml:"address"`
			Descr   string `xml:"descr"`
			Type    string `xml:"type"`
			Detail  string `xml:"detail"`
		} `xml:"alias"`
	} `xml:"aliases"`
	Proxyarp string `xml:"proxyarp"`
	Cron     struct {
		Text string `xml:",chardata"`
		Item []struct {
			Text    string `xml:",chardata"`
			Minute  string `xml:"minute"`
			Hour    string `xml:"hour"`
			Mday    string `xml:"mday"`
			Month   string `xml:"month"`
			Wday    string `xml:"wday"`
			Who     string `xml:"who"`
			Command string `xml:"command"`
		} `xml:"item"`
	} `xml:"cron"`
	Wol string `xml:"wol"`
	Rrd struct {
		Text     string `xml:",chardata"`
		Enable   string `xml:"enable"`
		Category string `xml:"category"`
	} `xml:"rrd"`
	Widgets struct {
		Text     string `xml:",chardata"`
		Sequence string `xml:"sequence"`
	} `xml:"widgets"`
	Openvpn  string `xml:"openvpn"`
	Dnshaper string `xml:"dnshaper"`
	Unbound  struct {
		Text              string `xml:",chardata"`
		Enable            string `xml:"enable"`
		Dnssec            string `xml:"dnssec"`
		ActiveInterface   string `xml:"active_interface"`
		OutgoingInterface string `xml:"outgoing_interface"`
		CustomOptions     string `xml:"custom_options"`
	} `xml:"unbound"`
	Revision struct {
		Text        string `xml:",chardata"`
		Time        string `xml:"time"`
		Description string `xml:"description"`
		Username    string `xml:"username"`
	} `xml:"revision"`
	Dhcpdv6 string `xml:"dhcpdv6"`
	Cert    struct {
		Text  string `xml:",chardata"`
		Refid string `xml:"refid"`
		Descr string `xml:"descr"`
		Type  string `xml:"type"`
		Crt   string `xml:"crt"`
		Prv   string `xml:"prv"`
	} `xml:"cert"`
	Gateways struct {
		Text        string `xml:",chardata"`
		GatewayItem struct {
			Text       string `xml:",chardata"`
			Interface  string `xml:"interface"`
			Gateway    string `xml:"gateway"`
			Name       string `xml:"name"`
			Weight     string `xml:"weight"`
			Ipprotocol string `xml:"ipprotocol"`
			Interval   string `xml:"interval"`
			Descr      string `xml:"descr"`
		} `xml:"gateway_item"`
		Defaultgw4 string `xml:"defaultgw4"`
	} `xml:"gateways"`
	Ppps      string `xml:"ppps"`
	Virtualip struct {
		Text string `xml:",chardata"`
		Vip  []struct {
			Text       string `xml:",chardata"`
			Mode       string `xml:"mode"`
			Interface  string `xml:"interface"`
			Vhid       string `xml:"vhid"`
			Advskew    string `xml:"advskew"`
			Advbase    string `xml:"advbase"`
			Password   string `xml:"password"`
			Descr      string `xml:"descr"`
			Type       string `xml:"type"`
			SubnetBits string `xml:"subnet_bits"`
			Subnet     string `xml:"subnet"`
			Uniqid     string `xml:"uniqid"`
		} `xml:"vip"`
	} `xml:"virtualip"`
	Hasync struct {
		Text                            string `xml:",chardata"`
		Synchronizerules                string `xml:"synchronizerules"`
		Synchronizeschedules            string `xml:"synchronizeschedules"`
		Synchronizealiases              string `xml:"synchronizealiases"`
		Synchronizenat                  string `xml:"synchronizenat"`
		Synchronizeipsec                string `xml:"synchronizeipsec"`
		Synchronizevirtualip            string `xml:"synchronizevirtualip"`
		Synchronizednsforwarder         string `xml:"synchronizednsforwarder"`
		Pfsyncpeerip                    string `xml:"pfsyncpeerip"`
		Pfsyncinterface                 string `xml:"pfsyncinterface"`
		Synchronizetoip                 string `xml:"synchronizetoip"`
		Username                        string `xml:"username"`
		Password                        string `xml:"password"`
		Synchronizeusers                string `xml:"synchronizeusers"`
		Synchronizeauthservers          string `xml:"synchronizeauthservers"`
		Synchronizecerts                string `xml:"synchronizecerts"`
		Synchronizeopenvpn              string `xml:"synchronizeopenvpn"`
		Synchronizedhcpd                string `xml:"synchronizedhcpd"`
		Synchronizewol                  string `xml:"synchronizewol"`
		Synchronizestaticroutes         string `xml:"synchronizestaticroutes"`
		Synchronizetrafficshaper        string `xml:"synchronizetrafficshaper"`
		Synchronizetrafficshaperlimiter string `xml:"synchronizetrafficshaperlimiter"`
		Synchronizetrafficshaperlayer7  string `xml:"synchronizetrafficshaperlayer7"`
		Synchronizecaptiveportal        string `xml:"synchronizecaptiveportal"`
		Pfsyncenabled                   string `xml:"pfsyncenabled"`
	} `xml:"hasync"`
	Dnsmasq struct {
		Text          string `xml:",chardata"`
		Enable        string `xml:"enable"`
		CustomOptions string `xml:"custom_options"`
		Interface     string `xml:"interface"`
	} `xml:"dnsmasq"`
	Installedpackages struct {
		Text        string `xml:",chardata"`
		Zabbixagent struct {
			Text   string `xml:",chardata"`
			Config struct {
				Text             string `xml:",chardata"`
				Agentenabled     string `xml:"agentenabled"`
				Server           string `xml:"server"`
				Serveractive     string `xml:"serveractive"`
				Hostname         string `xml:"hostname"`
				Listenip         string `xml:"listenip"`
				Listenport       string `xml:"listenport"`
				Refreshactchecks string `xml:"refreshactchecks"`
				Timeout          string `xml:"timeout"`
				Buffersend       string `xml:"buffersend"`
				Buffersize       string `xml:"buffersize"`
				Startagents      string `xml:"startagents"`
				Userparams       string `xml:"userparams"`
			} `xml:"config"`
		} `xml:"zabbixagent"`
	} `xml:"installedpackages"`
	Dyndnses string `xml:"dyndnses"`
}
