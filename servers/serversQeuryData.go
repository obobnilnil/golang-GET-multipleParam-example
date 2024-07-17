package servers

import (
	"database/sql"
	"multipleParam_git/handlers"
	"multipleParam_git/repositories"
	"multipleParam_git/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutesQueryData(router *gin.Engine, db *sql.DB) {

	r := repositories.NewRepositoryAdapter(db)
	s := services.NewServiceAdapter(r)
	h := handlers.NewHanerhandlerAdapter(s)

	router.GET("/api/getUniversalInfoByCatalog", h.GetUniversalInfoHandlers2)
}
