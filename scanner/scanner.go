package scanner

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// Result representa o resultado do scan de uma porta
type Result struct {
	IP      string
	Port    int
	State   string // "open" ou "closed"
	Service string // banner ou nome do serviço
}

// Scanner estrutura principal do scanner
type Scanner struct {
	TargetIPs []string
	Ports     []int
	Threads   int
	Timeout   time.Duration
}

// NewScanner cria uma nova instância de Scanner
func NewScanner(ips []string, ports []int, threads int, timeout time.Duration) *Scanner {
	return &Scanner{
		TargetIPs: ips,
		Ports:     ports,
		Threads:   threads,
		Timeout:   timeout,
	}
}

// Run executa o scan usando um pool de goroutines (worker pool)
func (s *Scanner) Run() []Result {
	var results []Result
	var mu sync.Mutex

	// Canais para tarefas (apenas portas por simplicidade)
	portsChan := make(chan int, s.Threads)
	var wg sync.WaitGroup

	// SYN scan concept: No TCP Connect Scan (usado aqui com net.Dial), o SO completa o 3-way handshake (SYN, SYN-ACK, ACK).
	// Em um SYN scan real (stealth), enviamos apenas SYN e esperamos SYN-ACK, encerrando com RST.
	// Como estamos usando apenas a biblioteca padrão net sem raw sockets, fazemos o TCP Connect scan.

	// Inicia os workers
	for i := 0; i < s.Threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for port := range portsChan {
				for _, ip := range s.TargetIPs {
					target := fmt.Sprintf("%s:%d", ip, port)
					// TCP Connect
					conn, err := net.DialTimeout("tcp", target, s.Timeout)
					
					res := Result{
						IP:   ip,
						Port: port,
					}

					if err != nil {
						res.State = "closed"
					} else {
						res.State = "open"
						conn.Close()
						// Tenta pegar o banner (serviço)
						res.Service = GrabBanner(ip, port, s.Timeout)
						
						mu.Lock()
						results = append(results, res)
						mu.Unlock()
					}
				}
			}
		}()
	}

	// Envia as tarefas
	for _, port := range s.Ports {
		portsChan <- port
	}
	close(portsChan)

	// Espera todos terminarem
	wg.Wait()

	return results
}
