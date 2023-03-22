package utilities

import (
	"strconv"
	bd "v3/bd/task"
)

func IsValidActionUser(taskID string, userID string) bool {

	idTask, _ := strconv.ParseInt(taskID, 10, 0)
	idUser, _ := strconv.ParseInt(userID, 10, 0)

	task := bd.GetTask(int(idTask))

	if int(idUser) != task.UserID {
		return false
	}

	return true
}
