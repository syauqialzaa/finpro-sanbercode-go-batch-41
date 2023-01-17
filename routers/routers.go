package routers

import (
	"database/sql"
	"final-project-sanbercode-go-batch-41/controllers"
	"final-project-sanbercode-go-batch-41/middleware"

	"github.com/gin-gonic/gin"
)

func Routers(db *sql.DB) *gin.Engine {
	router := gin.Default()

	// AUTHENTICATION
	router.POST("/register", middleware.Register)
	router.POST("/login", middleware.Login)

	// CRUD PROJECTS
	router.GET("/projects", controllers.GetAllProject)
	router.POST("/projects", middleware.AuthMiddlware(), controllers.CreateProject)
	router.PUT("/projects/:id", middleware.AuthMiddlware(), controllers.UpdateProject)
	router.DELETE("/projects/:id", middleware.AuthMiddlware(), controllers.DeleteProject)
	router.GET("/projects/:id/tasks", controllers.GetTaskByProjectID)    // get task by projectID

	// CRUD TASKS
	router.GET("/tasks", controllers.GetAllTask)
	router.POST("/tasks", middleware.AuthMiddlware(), controllers.CreateTask)
	router.PUT("/tasks/:id", middleware.AuthMiddlware(), controllers.UpdateTask)
	router.DELETE("/tasks/:id", middleware.AuthMiddlware(), controllers.DeleteTask)

	// CRUD TEAM
	router.GET("/teams", controllers.GetAllTeam)
	router.POST("/teams", middleware.AuthMiddlware(), controllers.CreateTeam)
	router.PUT("/teams/:id", middleware.AuthMiddlware(), controllers.UpdateTeam)
	router.DELETE("/teams/:id", middleware.AuthMiddlware(), controllers.DeleteTeam)
	router.GET("/teams/:id/members", controllers.GetMemberByTeamID)    // get member by teamID

	// CRUD MEMBER
	router.GET("/members", controllers.GetAllMember)
	router.POST("/members", middleware.AuthMiddlware(), controllers.AddMember)
	router.PUT("/members/:id", middleware.AuthMiddlware(), controllers.UpdateMember)
	router.DELETE("/members/:id", middleware.AuthMiddlware(), controllers.DeleteMember)

	return router
}