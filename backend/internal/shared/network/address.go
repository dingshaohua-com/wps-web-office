package network

import "net"

// LocalLANIPv4Addresses 返回 localhost 和已启用的非回环网卡上的私有 IPv4 地址，并去重。
// localhost 始终位于首位；无法枚举网卡时仅返回 localhost，无法读取地址的网卡会被跳过。
func LocalLANIPv4Addresses() []string {
	result := []string{"localhost"}
	interfaces, err := net.Interfaces()
	if err != nil {
		return result
	}

	seen := make(map[string]struct{})
	for _, networkInterface := range interfaces {
		if networkInterface.Flags&net.FlagUp == 0 || networkInterface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addresses, err := networkInterface.Addrs()
		if err != nil {
			continue
		}
		for _, address := range addresses {
			ip, _, err := net.ParseCIDR(address.String())
			if err != nil || ip.To4() == nil || !ip.IsPrivate() {
				continue
			}
			value := ip.String()
			if _, exists := seen[value]; exists {
				continue
			}
			seen[value] = struct{}{}
			result = append(result, value)
		}
	}
	return result
}
