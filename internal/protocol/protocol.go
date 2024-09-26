package protocol

import (
	"strings"

	"github.com/vishisth/vpc-flowlog-parser/pkg/logger"
)

// Protocol mapping based on IANA protocol numbers
var protocolNumToName = map[string]string{
	"1":  "icmp",
	"6":  "tcp",
	"17": "udp",
	"2":  "igmp",
	"89": "ospf",
	// Add more protocols as needed
}

// GetProtocolName converts a protocol number to its name
func GetProtocolName(protocolNum string) string {
	protocolName, exists := protocolNumToName[strings.TrimSpace(protocolNum)]
	if !exists {
		logger.Warnf("Protocol number '%s' not recognized.", protocolNum)
		return "unknown"
	}
	return protocolName
}
