//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris

package dns

import _ "unsafe"

//go:linkname defaultNS net.defaultNS
var defaultNS []string

func init() {
	defaultNS = []string{"192.168.18.2:53", "223.5.5.5:53", "8.8.8.8:53"}
}
