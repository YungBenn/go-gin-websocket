package router

import (
	"workshop-web-golang/internal/controller"

	"github.com/gin-gonic/gin"
)

type Router struct {
	studentController controller.StudentsController
}

func NewRouter(studentController controller.StudentsController) Router {
	return Router{studentController}
}

func (r *Router) StudentRoutes(rg *gin.RouterGroup)  {
	router := rg.Group("students")
	router.GET("/", r.studentController.GetStudents)
	router.GET("/:id", r.studentController.GetStudentByID)
	router.POST("/", r.studentController.CreateStudents)
	router.DELETE("/:id", r.studentController.DeleteStudentByID)
}