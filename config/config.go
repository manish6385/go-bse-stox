package config

type Conf struct {
	*Gin
	GinServer GinServer
}

type GinServer struct {
	hostIPAddress string
	hostPort      int
}

func Init() *Conf {
	var conf Conf

	conf.GinServer.hostIPAddress = ""
	conf.GinServer.hostPort = 8080

	return &conf
}
