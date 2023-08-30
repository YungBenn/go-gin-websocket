package controller

import (
	"net/http"
	"workshop-web-golang/internal/model"
	"workshop-web-golang/internal/util"
	ws "workshop-web-golang/internal/websocket"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StudentsController struct {
	DB *gorm.DB
}

func NewStudentsController(DB *gorm.DB) StudentsController {
	return StudentsController{DB}
}

func (sc *StudentsController) CreateStudents(c *gin.Context) {
	var payload *model.NewStudent

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	newStudent := model.Students{
		Nama:    payload.Nama,
		Jurusan: payload.Jurusan,
		Status:  payload.Status,
	}

	result := sc.DB.Create(&newStudent)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	ws.SendWebSocketUpdate("Create new student success")
	c.JSON(http.StatusCreated, util.DataResponse("Created", newStudent))
}

func (sc *StudentsController) GetStudents(c *gin.Context) {
	var students []model.Students

	results := sc.DB.Find(&students)
	if results.Error != nil {
		c.JSON(400, gin.H{
			"error": results.Error.Error(),
		})
		return
	}

	ws.SendWebSocketUpdate("Get students success")
	c.JSON(http.StatusOK, util.DataResponse("OK", students))
}

func (sc *StudentsController) GetStudentByID(c *gin.Context) {
	studentID := c.Param("id")

	var student model.Students
	if err := sc.DB.First(&student, studentID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ws.SendWebSocketUpdate("Get student by ID success")
	c.JSON(http.StatusOK, util.DataResponse("OK", student))
}

func (sc *StudentsController) DeleteStudentByID(c *gin.Context) {
	studentID := c.Param("id")

	if err := sc.DB.Delete(&model.Students{}, "id = ?", studentID).Error; err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	ws.SendWebSocketUpdate("Delete student by ID success")
	c.JSON(http.StatusOK, util.MessageResponse("Deleted"))
}
