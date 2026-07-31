package utils

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

// Cores ANSI para o terminal
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

// Top100Ports retorna as portas mais comuns (resumido)
func Top100Ports() []int {
	return []int{
		21, 22, 23, 25, 53, 80, 110, 111, 135, 139, 143, 443, 445, 993, 995, 1723, 3306, 3389, 5900, 8080,
	}
}

// ParsePorts parseia a string de portas e retorna um slice de ints
func ParsePorts(ports string) ([]int, error) {
	var result []int
	if ports == "top100" {
		return Top100Ports(), nil
	}

	parts := strings.Split(ports, ",")
	for _, part := range parts {
		if strings.Contains(part, "-") {
			ranges := strings.Split(part, "-")
			if len(ranges) != 2 {
				return nil, fmt.Errorf("formato de range inválido: %s", part)
			}
			start, err1 := strconv.Atoi(ranges[0])
			end, err2 := strconv.Atoi(ranges[1])
			if err1 != nil || err2 != nil || start > end {
				return nil, fmt.Errorf("range inválido: %s", part)
			}
			for i := start; i <= end; i++ {
				result = append(result, i)
			}
		} else {
			port, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("porta inválida: %s", part)
			}
			result = append(result, port)
		}
	}
	return result, nil
}

// ParseHosts processa a string do host (pode ser um IP, domínio ou CIDR)
func ParseHosts(hostStr string) ([]string, error) {
	if strings.Contains(hostStr, "/") {
		// CIDR
		ip, ipNet, err := net.ParseCIDR(hostStr)
		if err != nil {
			return nil, err
		}
		var ips []string
		for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); inc(ip) {
			ips = append(ips, ip.String())
		}
		// Remove network address and broadcast address
		if len(ips) > 2 {
			return ips[1 : len(ips)-1], nil
		}
		return ips, nil
	}
	return []string{hostStr}, nil
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// PrintBanner exibe o banner do projeto
func PrintBanner() {
	banner := `
    ____             __  _____                                 
   / __ \____  _____/ /_/ ___/_________ _____  ____  ___  _____
  / /_/ / __ \/ ___/ __/\__ \/ ___/ __ ` + "`" + `/ __ \/ __ \/ _ \/ ___/
 / ____/ /_/ / /  / /_ ___/ / /__/ /_/ / / / / / / /  __/ /    
/_/    \____/_/   \__//____/\___/\__,_/_/ /_/_/ /_/\___/_/     
`
	fmt.Println(Cyan + banner + Reset)
	fmt.Println(Yellow + "v1.0.0 - Fast Concurrent Port Scanner" + Reset)
	fmt.Println(Yellow + "--------------------------------------" + Reset)
}
