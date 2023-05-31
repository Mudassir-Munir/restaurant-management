package routes

import (
	"github.com/Mudassir-Munir/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)


func NoteRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/notes", controllers.GetNotes())
	incommingRoutes.GET("/notes/:note_id", controllers.GetNote())
	incommingRoutes.GET("/notes", controllers.CreateNote())
	incommingRoutes.PATCH("/notes/:note_id", controllers.UpdateNote())
}
