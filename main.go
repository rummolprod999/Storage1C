package main

func init() {
	CreateEnv()
}

func main() {
	defer SaveStack()
	Logging("start")
	st := Storage1C{}
	st.Init()
	st.Run()
	Logging("end")
}
