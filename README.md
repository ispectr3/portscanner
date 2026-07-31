# 🚀 PortScanner-Go

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/License-MIT-green.svg?style=for-the-badge)

**PortScanner-Go** é um scanner de portas de rede rápido e concorrente escrito em Go. Ele utiliza o modelo de concorrência nativo do Go (*goroutines* com o padrão *Worker Pool*) para varrer grandes quantidades de portas rapidamente.

## ✨ Funcionalidades

- **Alta Performance:** Escaneamento concorrente com número configurável de goroutines.
- **TCP Connect Scan:** Scanner baseado na API `net` padrão do Go. *(Nota: O conceito de stealth SYN scan requer raw sockets e privilégios elevados; no entanto, este projeto foca em TCP Connect)*.
- **Detecção de Serviço (Banner Grabbing):** Tenta identificar o serviço rodando na porta enviando requisições genéricas.
- **Alvos Flexíveis:** Escaneie um IP único ou um range usando formato CIDR (ex: `192.168.0.0/24`).
- **Formatos de Porta:** Suporta portas separadas por vírgula (`80,443`), ranges (`1-1024`) ou as top 100 portas (`top100`).
- **Exportação de Dados:** Salve os resultados em JSON ou CSV para análise posterior.
- **Interface Bonita:** Banner em ASCII e saída colorida no terminal usando ANSI escape codes.

## 🛠️ Instalação

Certifique-se de ter o Go instalado (versão 1.21 ou superior).

```bash
git clone https://github.com/kaique/portscanner-go.git
cd portscanner-go
go mod tidy
go build -o portscanner main.go
```

## 💻 Uso

```bash
./portscanner -host <alvo> [opções]
```

### Exemplos:

Escanear as top 100 portas de um único host:
```bash
./portscanner -host 192.168.1.10 -ports top100
```

Escanear range de portas em uma rede inteira, usando 500 threads e exportando para CSV:
```bash
./portscanner -host 192.168.1.0/24 -ports 1-1000 -threads 500 -output results.csv -format csv
```

### Opções Disponíveis:

- `-host`: Host único ou rede CIDR (Obrigatório).
- `-ports`: Portas a escanear (padrão: `top100`).
- `-threads`: Número de goroutines concorrentes (padrão: `100`).
- `-timeout`: Timeout em milissegundos por porta (padrão: `1000`).
- `-output`: Arquivo de saída opcional.
- `-format`: Formato de saída para o arquivo (`json` ou `csv`).

## ⚠️ Disclaimer

Este projeto foi desenvolvido com propósitos estritamente educacionais e de pesquisa em cibersegurança. O uso desta ferramenta para testar infraestruturas sem autorização prévia mútua pode ser ilegal em seu país ou região. O desenvolvedor não se responsabiliza por qualquer uso indevido.

## 📜 Licença

Distribuído sob a licença MIT. Veja `LICENSE` para mais informações.
