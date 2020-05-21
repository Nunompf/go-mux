package main

func main() {
	a := App{}
	a.Initialize(DbUser, DbPw, DbName)
	a.Run(":8010")
}

