package main

func main() {

	////server:gin

	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.String(200, "pong")
	//})
	//err := r.RunTLS(":443", "certs/server.crt", "certs/server.key")

	////server:grpc

	//creds, err := credentials.NewServerTLSFromFile("certs/server.crt", "certs/server.key")
	//if err != nil {
	//	log.Fatalf("failed to load TLS cert: %v", err)
	//}
	//grpcServer := grpc.NewServer(grpc.Creds(creds))
	//pb.RegisterMyServiceServer(grpcServer, &MyService{})
	//grpcServer.Serve(lis)

	//client:http

	//certPool := x509.NewCertPool()
	//caCert, _ := os.ReadFile("certs/ca.crt")
	//certPool.AppendCertsFromPEM(caCert)
	//
	//client := &http.Client{
	//	Transport: &http.Transport{
	//		TLSClientConfig: &tls.Config{
	//			RootCAs: certPool, // 指定信任的 CA
	//		},
	//	},
	//}
	//
	//resp, err := client.Get("https://server.mydomain.com:443/ping")

	//	//client:grpc
	//
	//	caCert, _ := os.ReadFile("certs/ca.crt")
	//	certPool := x509.NewCertPool()
	//	certPool.AppendCertsFromPEM(caCert)
	//
	//	// mTLS
	//	cert, _ := tls.LoadX509KeyPair("certs/client.crt", "certs/client.key")
	//
	//	creds := credentials.NewTLS(&tls.Config{
	//		Certificates:       []tls.Certificate{cert}, // 若开启客户端验证
	//		RootCAs:            certPool,
	//		InsecureSkipVerify: false, // 不建议跳过
	//		ServerName:         "server.mydomain.com",
	//	})
	//
	//	conn, err := grpc.Dial("server.mydomain.com:443", grpc.WithTransportCredentials(creds))

}
