package mono;


type Config struct {
	Addr string
}

func(c *Config)addr() string {
	addr := c.Addr
	if "" == addr {
		addr = ":80"
	}
	return addr
}

