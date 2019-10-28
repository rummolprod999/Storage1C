package main

func init() {
	CreateEnv()
}

func main() {
	defer SaveStack()
	Logging("start")
	Logging("end")
}
