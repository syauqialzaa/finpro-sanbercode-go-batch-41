package routers

import (
	"database/sql"
	"final-project-sanbercode-go-batch-41/controllers"

	"github.com/gin-gonic/gin"
)

func Routers(db *sql.DB) *gin.Engine {
	router := gin.Default()

	// CRUD PROJECTS
	router.GET("/projects", controllers.GetAllTeam)
	router.POST("/projects", controllers.CreateTeam)
	// PUT("/projects/:id")
	// DELETE("/projects/:id")
	// GET("/projects/:id/tasks")	// get task by projectID

	// CRUD TASKS
	// GET("/tasks")
	// POST("/tasks")
	// PUT("/tasks/:id")
	// DELETE("/tasks/:id")

	// CRUD TEAM
	// GET("/teams")
	// POST("/teams")
	// PUT("/teams/:id")
	// DELETE("/teams/:id")
	// GET("/teams/:id/member")	// get member by teamID

	// CRUD MEMBER
	// GET("/members")
	// POST("/members")
	// PUT("/members/:id")
	// DELETE("/members/:id")

	return router
}