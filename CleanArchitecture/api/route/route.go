package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/naol86/learn-go/CleanArchitecture/config"
	"github.com/naol86/learn-go/CleanArchitecture/pkg/mongo"
)



func Setup(env *config.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRoute := gin.Group("")
	// public routes
	NewSignupRoute(env, timeout, db, publicRoute);
	NewLoginRoute(env, timeout, db, publicRoute);
}