package routers

import (
	"github.com/gin-gonic/gin"
	"online-learning-platform/internal/rest/handlers"
	"online-learning-platform/pkg/middleware"
)

type Routers struct {
	authHandlers  handlers.AuthHandlers
	courseHandler handlers.CourseHandler
	lessonHandler handlers.LessonHandler
}

func NewRouters(authHandlers handlers.AuthHandlers, courseHandler handlers.CourseHandler, lessonHandler handlers.LessonHandler) *Routers {
	return &Routers{
		authHandlers:  authHandlers,
		courseHandler: courseHandler,
		lessonHandler: lessonHandler,
	}
}

func (r *Routers) SetupRoutes(app *gin.Engine) {
	authRouter := app.Group("/auth")
	{
		authRouter.GET("/who-am-i", r.authHandlers.WhoAmI)
		authRouter.GET("/register", r.authHandlers.RegisterForm)
		authRouter.POST("/register", r.authHandlers.Register)
		authRouter.GET("/login", r.authHandlers.LoginForm)
		authRouter.POST("/login", r.authHandlers.Login)
		authRouter.POST("/logout", middleware.RequireAuthMiddleware, r.authHandlers.Logout)
		authRouter.PUT("/update", middleware.RequireAuthMiddleware, r.authHandlers.Update)
		authRouter.DELETE("/delete", middleware.RequireAuthMiddleware, r.authHandlers.Delete)
	}

	courseRouter := app.Group("/courses")
	{
		courseRouter.GET("/", middleware.RequireAuthMiddleware, r.courseHandler.GetAllCourses)
		courseRouter.GET("/:id", middleware.RequireAuthMiddleware, r.courseHandler.GetCourseByID)
		courseRouter.GET("/add", middleware.RequireAuthMiddleware, r.courseHandler.CreateCourseForm)
		courseRouter.POST("/add", middleware.RequireAuthMiddleware, r.courseHandler.CreateCourse)
		courseRouter.PUT("/:id/update", middleware.RequireAuthMiddleware, r.courseHandler.UpdateCourse)
		courseRouter.DELETE("/:id/delete", middleware.RequireAuthMiddleware, r.courseHandler.DeleteCourse)
	}

	lessonRouter := app.Group("/lessons")
	{
		lessonRouter.GET("/", middleware.RequireAuthMiddleware, r.lessonHandler.GetAllLessons)
		lessonRouter.GET("/:id", middleware.RequireAuthMiddleware, r.lessonHandler.GetLessonByID)
		lessonRouter.GET("/add", middleware.RequireAuthMiddleware, r.lessonHandler.CreateLessonForm)
		lessonRouter.POST("/add", middleware.RequireAuthMiddleware, r.lessonHandler.CreateLesson)
		lessonRouter.PUT("/:id/update", middleware.RequireAuthMiddleware, r.lessonHandler.UpdateLesson)
		lessonRouter.DELETE("/:id/delete", middleware.RequireAuthMiddleware, r.lessonHandler.DeleteLessonByID)
	}
}
