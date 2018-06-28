package tools

import "github.com/joho/godotenv"

func init() {
	// 必须要使用godotenv.Load()在初始化的时候，否则无法使用os.Getenv("PORT")读取到.env文件
	godotenv.Load()
}
