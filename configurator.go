package main

var Config = struct {
	ServerPort   uint   `default:"8081"`
	PathCTool1cd string `required:"true"`
	DB           struct {
		Name     string
		Host     string `default:"localhost"`
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
		Port     uint   `default:"3306"`
	}

	Storages []struct {
		Name string `required:"true"`
		Path string `required:"true"`
	}
}{}
