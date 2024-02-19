package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	sqlDB := model.ConnectionDB() // DB接続
	defer sqlDB.Close()	// DB接続を閉じる
	e := echo.New()	// echoのインスタンスを作成
	router.SetRouter(e)	// ルーティングを設定
}
