package controller

import (
	"log"
	authService "mit-api/pkg/auth/service"
	userModel "mit-api/pkg/user/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateTags login
// @Summary login
// @Description login
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body request.LoginRequest true "user"
// @Router /auth/login [post]
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		// otp := helpers.GenerateOpt()
		// mail.Send("Verification OTP", "Your OTP is: "+otp.Secret())
		var user userModel.User
		if err := c.BindJSON(&user); err != nil {
			log.Println(err)
		}

		loginResult := authService.Login(*user.Email, *user.Password)

		if !loginResult.Success {
			c.JSON(loginResult.HttpCode, &loginResult)

			return
		}

		c.JSON(http.StatusOK, &loginResult)
	}
}

// CreateTags signup
// @Summary signup
// @Description signup
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body request.SignUpRequest true "user"
// @Router /auth/signup [post]
func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user userModel.User

		if err := c.BindJSON(&user); err != nil {
			log.Println(err)
		}

		signUpResult := authService.SignUp(user)
		if !signUpResult.Success {
			c.JSON(http.StatusBadRequest, signUpResult)

			return
		}

		c.JSON(http.StatusOK, &signUpResult)
	}
}
