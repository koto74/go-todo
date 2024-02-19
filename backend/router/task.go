package router

import (
	"os"
	"github.com/labsyspharm/echo"	
)

func GetTasksHandler(c echo.Context) error {
	// tasksにDBのタスクすべてを代入する。
	tasks, err := model.GetTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	// tasksをJSONで返す
	return c.JSON(http.StatusOK, tasks)
}
