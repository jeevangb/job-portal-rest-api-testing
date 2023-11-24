package handler

import (
	"jeevan/jobportal/internal/auth"
	"jeevan/jobportal/internal/middleware"
	"jeevan/jobportal/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service service.UserService
}

func SetApi(a auth.Authentication, src service.UserService) *gin.Engine {
	r := gin.New()
	// h := NewHandler()
	h := Handler{
		Service: src,
	}
	m, err := middleware.NewMiddleware(a)
	if err != nil {
		log.Panic("middlewares not setup")
	}
	r.Use(m.Log(), gin.Recovery())
	r.GET("/check", m.Authenticate(Check))

	r.POST("/signup", h.SignUp)
	r.POST("/Login", h.Login)

	r.POST("/addCompany", m.Authenticate((h.AddCompany)))
	r.GET("/viewAllCompanies/all", m.Authenticate(h.ViewAllCompanies))
	r.GET("/viewCompany/:id", m.Authenticate((h.ViewCompany)))

	r.POST("/addJob/:cid", m.Authenticate(h.AddJob))
	r.GET("/job/view/:id", m.Authenticate(h.ViewJob))
	r.GET("/viewAllJobs/all", m.Authenticate(h.ViewAllJobs))
	r.GET("/viewJobById/:id", m.Authenticate(h.ViewJobByID))
	r.POST("/ProcessJobById", m.Authenticate(h.ProcessJobDetails))

	r.POST("/ForgotPassword", h.ForgotPasswod)
	r.POST("/ResetPassword", h.ResetPassword)

	return r

}

func Check(c *gin.Context) {
	c.JSON(200, gin.H{
		"Message": "ok",
	})
}
