package model

import (
	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

// Task型はuuid.UUID型のIDとstring型のTitle、bool型のCompletedを持つ
type Task struct {
	id		uuid.UUID
	Title	string
	Completed	bool
}

// 引数はなく、戻り値は[]Task型のtasksとerror型のerr
func GetTasks() ([]Task, error) {
	// からのタスクのスライスである、tasksを定義
	var tasks []Task

	// tasksにDBのタスクすべてを代入する。
	err := db.Find(&tasks).Error

	return tasks, err
}
