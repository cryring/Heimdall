package main

type Config struct {
	TcpAddress       string `arg:"--tcp_addr"`
	TlsAddress       string `arg:"--tls_addr"`
	WebsocketAddress string `arg:"--ws_addr"`
	ClusterAddress   string `arg:"--cluster_addr"`
	LogPath          string `arg:"--log_path"`
}
