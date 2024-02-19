package router

import (
	"os"

	"github.com/labstackc/echo/v4/middleware"

	_ "net/http"
	"github.com/labstack/echo/v4"
)

// Routingを設定する関数 引数はecho.echo型であり、戻り値はerror型
func SetRouter(e *echo.Echo) error {

	// APIが叩かれた時にログを出す
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
    	Format: "${time_rfc3339_nano} ${host} ${method} ${uri} ${status} ${header}\n",
    	Output: os.Stdout,
	}))
	// 予想外のエラーが発生した際でも、サーバーを落とさないようにする
	e.Use(middleware.Recover())
	// CORSに対応する
	e.Use(middleware.CORS())

	// ルーティングを設定
	e.GET("/api/tasks", GetTasksHandler)
	e.POST("/tasks", CreateTask)
	e.PUT("/tasks/:id", UpdateTask)
	e.DELETE("/tasks/:id", DeleteTask)

	// サーバーを起動
	err := e.Start(":8000")

	return err
}
