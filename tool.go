package tool

import "fmt"

const (
	ConnectionString string = "127.0.0.1:XXX..."
)

func Hello() string {
	return "hello from tool project"
}

func ConnectDB() {
	fmt.Println("資料庫連接中...")
}

func CloseDB() {
	fmt.Println("資料庫關閉")
}
