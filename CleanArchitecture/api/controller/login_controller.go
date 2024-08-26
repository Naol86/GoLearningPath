package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/naol86/learn-go/CleanArchitecture/config"
	"github.com/naol86/learn-go/CleanArchitecture/domain"
	"golang.org/x/crypto/bcrypt"
)


type LoginController struct {
	LoginUsecase domain.LoginUsecase;
	Env *config.Env;
}

func (lc *LoginController) Login(c *gin.Context) {
	var request domain.LoginRequest;
	err := c.ShouldBind(&request);
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user, err := lc.LoginUsecase.GetUserByEmail(c, request.Email);
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid email"})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password));
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid password"})
		return
	}
	accessToken, _ := lc.LoginUsecase.CreateAccessToken(&user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour);
	refreshToken, _ := lc.LoginUsecase.CreateRefreshToken(&user, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour);
	var response domain.LoginResponse;
	var userResponse domain.UserResponse;
	userResponse.ID = user.ID.Hex();
	userResponse.Name = user.Name;
	userResponse.Email = user.Email;
	response.AccessToken = accessToken;
	response.RefreshToken = refreshToken;
	response.User = userResponse;
	c.JSON(200, gin.H{"data": response});
}