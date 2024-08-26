package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/naol86/learn-go/CleanArchitecture/api/controller"
	"github.com/naol86/learn-go/CleanArchitecture/config"
	"github.com/naol86/learn-go/CleanArchitecture/domain"
	"github.com/naol86/learn-go/CleanArchitecture/pkg/mongo"
	"github.com/naol86/learn-go/CleanArchitecture/repository"
	"github.com/naol86/learn-go/CleanArchitecture/usecase"
)


func NewLoginRoute(env *config.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	lc := controller.LoginController{
		LoginUsecase: usecase.NewLoginUseCase(ur, timeout),
		Env: env,
	}

	group.POST("/login", lc.Login);
}