package controller

import (
	"github.com/BevisDev/BevisBot/internal/service"
)

type TaskController struct {
	svc service.TaskService
}

func NewTaskController(
	svc service.TaskService,
) *TaskController {
	return &TaskController{
		svc: svc,
	}
}

//func (t *TaskController) CreateTask(c *gin.Context) {
//	var r request.Task
//	if err := c.ShouldBindJSON(&r); err != nil {
//		response.BadRequest(c, "", err.Error())
//		return
//	}
//
//	id, err := t.svc.CreateTask(c.Request.Context(), &r)
//	if err != nil {
//		response.BadRequest(c, "", err.Error())
//		return
//	}
//
//	response.Created(c, map[string]interface{}{
//		"id": id,
//	})
//}

//func (t *TaskController) Search(c *gin.Context) {
//	var r request.TaskSearch
//	if err := c.ShouldBindQuery(&r); err != nil {
//		response.BadRequest(c, "", err.Error())
//		return
//	}
//
//	tasks, total, err := t.svc.SearchTasks(
//		c.Request.Context(),
//		service.TaskFilter{
//			Q:        q.Q,
//			Page:     q.Page,
//			PageSize: q.Size,
//		},
//	)
//	if err != nil {
//		response.BadRequest(c, "", err.Error())
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"data": tasks,
//		"paging": gin.H{
//			"page":  q.Page,
//			"size":  q.Size,
//			"total": total,
//		},
//	})
//}
