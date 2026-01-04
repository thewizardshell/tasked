package handler

import (
	"net/http"
	"strconv"
	"tasked/internal/services"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	service *services.TaskService
}

func NewTaskHandler(service *services.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

// GetTask godoc
// @Summary Obtener tarea por ID
// @Description Retorna una tarea específica por su ID
// @Tags tasks
// @Security Bearer
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} domain.Task
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /tasks/{id} [get]
func (h *TaskHandler) GetTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task id"})
		return
	}

	task, err := h.service.GetTaskById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// ListTasksByUser godoc
// @Summary Listar tareas por usuario
// @Description Retorna todas las tareas de un usuario específico
// @Tags tasks
// @Security Bearer
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {array} domain.Task
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{user_id}/tasks [get]
func (h *TaskHandler) ListTasksByUser(c *gin.Context) {
	userIdParam := c.Param("id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	tasks, err := h.service.ListTaskByUser(c.Request.Context(), userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// UpdateTask godoc
// @Summary Actualizar tarea
// @Description Actualiza los datos de una tarea existente
// @Tags tasks
// @Security Bearer
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param task body UpdateTaskRequest true "Datos a actualizar"
// @Success 200 {object} domain.Task
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks/{id} [put]
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task id"})
		return
	}

	var req UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.service.UpdateTask(c.Request.Context(), id, req.Title, req.Description, req.Status, req.Priority, req.DueDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// DeleteTask godoc
// @Summary Eliminar tarea
// @Description Elimina una tarea del sistema
// @Tags tasks
// @Security Bearer
// @Param id path int true "Task ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks/{id} [delete]
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task id"})
		return
	}

	if err := h.service.DeleteTask(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}

// UpdateStatus godoc
// @Summary Actualizar estado de tarea
// @Description Actualiza únicamente el estado de una tarea
// @Tags tasks
// @Security Bearer
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param status body UpdateStatusRequest true "Nuevo estado"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks/{id}/status [patch]
func (h *TaskHandler) UpdateStatus(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task id"})
		return
	}

	var req UpdateStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.service.UpdateStatus(c.Request.Context(), id, req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update status"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// CreateTask godoc
// @Summary Crear una tarea
// @Description Crea una nueva tarea en el sistema
// @Tags tasks
// @Security Bearer
// @Accept json
// @Produce json
// @Param task body CreateTaskRequest true "Datos de la tarea"
// @Success 201 {object} domain.Task
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks [post]
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.service.CreateTask(c.Request.Context(), req.Title, req.Description, req.Status, req.Priority, req.UserID, req.DueDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}

type UpdateTaskRequest struct {
	Title       string `json:"title" binding:"required" example:"Completar informe"`
	Description string `json:"description" example:"Terminar el informe mensual"`
	Status      string `json:"status" example:"pending"`
	Priority    string `json:"priority" example:"high"`
	DueDate     string `json:"due_date" example:"2024-12-31"`
}

type UpdateStatusRequest struct {
	Status string `json:"status" binding:"required" example:"completed"`
}
type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required" example:"Completar informe"`
	Description string `json:"description" example:"Terminar el informe mensual"`
	Status      string `json:"status" example:"pending"`
	Priority    string `json:"priority" example:"high"`
	UserID      int64  `json:"user_id" binding:"required" example:"1"`
	DueDate     string `json:"due_date" example:"2024-12-31"`
}
