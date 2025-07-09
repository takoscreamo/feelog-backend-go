package main

import (
	"feelog-backend/infra"
	"feelog-backend/routes"
	"feelog-backend/routes/middleware"

	"feelog-backend/api/controllers"
	"feelog-backend/repositories"
	"feelog-backend/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	dbConn := infra.SetupDB()
	middleware.SetAuthDB(dbConn)

	diaryRepository := repositories.NewDiaryRepository(dbConn)
	diaryUsecase := usecases.NewDiaryUsecase(diaryRepository)
	diaryController := controllers.NewDiaryController(diaryUsecase)

	diaryAnalysisUsecase := usecases.NewDiaryAnalysisUsecase(diaryRepository)
	diaryAnalysisController := controllers.NewDiaryAnalysisController(diaryAnalysisUsecase)

	userRepo := repositories.NewUserRepository(dbConn)
	withdrawUsecase := usecases.NewUserWithdrawUsecase(userRepo, diaryRepository)
	userController := controllers.NewUserController(userRepo, withdrawUsecase)

	router := gin.Default()

	// CORS設定を追加
	routes.SetupCORS(router)

	// SwaggerUIとOpenAPI仕様書のエンドポイントを設定
	routes.SetupSwaggerEndpoints(router)

	// APIエンドポイントを設定
	routes.SetupAPIEndpoints(router, diaryController, diaryAnalysisController, userController)

	router.Run()
}
