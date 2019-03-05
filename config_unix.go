// +build !windows

package main

const (
	rcFname     = "rc.conf"
	directFname = "direct.conf"
	proxyFname  = "proxy.conf"
	rejectFname = "reject.conf"
	CNIPFname   = "china_ip_list.conf"

	newLine = "\n"
)

func getDefaultRcFile() string {
	var homeConfigDir = path.Join(getUserHomeDir(), ".meow")
	if _, err := os.Stat(homeConfigDir); err == nil {
		return path.Join(path.Join(homeConfigDir, rcFname))
	}
	return rcFname
}
