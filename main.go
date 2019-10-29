package main

func init() {
	CreateEnv()
}

func main() {
	defer SaveStack()
	Logging("start")
	st := Storage1C{}
	st.Run()
	Logging("end")
}
