package scanner

import (
	"fmt"
	"net"
	"time"
)

// GrabBanner tenta ler o banner inicial do serviço na porta aberta
func GrabBanner(ip string, port int, timeout time.Duration) string {
	target := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", target, timeout)
	if err != nil {
		return ""
	}
	defer conn.Close()

	// Timeout de leitura
	conn.SetReadDeadline(time.Now().Add(timeout))

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		// Tenta enviar uma requisição HTTP genérica se não receber nada
		if port == 80 || port == 443 || port == 8080 {
			conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
			n, err = conn.Read(buffer)
			if err == nil && n > 0 {
				return parseHTTPBanner(string(buffer[:n]))
			}
		}
		return ""
	}

	return cleanBanner(string(buffer[:n]))
}

func cleanBanner(banner string) string {
	// Remove quebras de linha para exibição formatada
	b := []byte(banner)
	for i, v := range b {
		if v == '\n' || v == '\r' || v == '\t' {
			b[i] = ' '
		}
	}
	return string(b)
}

func parseHTTPBanner(banner string) string {
	// Extrai a primeira linha do cabeçalho
	for i := 0; i < len(banner); i++ {
		if banner[i] == '\n' || banner[i] == '\r' {
			return banner[:i]
		}
	}
	return banner
}
