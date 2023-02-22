package di

import (
	"database/sql"

	"github.com/atuy1213/myscrap/myscrap-api/adapter/controller"
	"github.com/atuy1213/myscrap/myscrap-api/adapter/repository"
	"github.com/atuy1213/myscrap/myscrap-api/adapter/util"
	"github.com/atuy1213/myscrap/myscrap-api/driver/db"
	"github.com/atuy1213/myscrap/myscrap-api/entity/domain"
	"github.com/atuy1213/myscrap/myscrap-api/usecase"
)

// inject db

func InjectDB() *sql.DB {
	return db.NewPostgresDB()
}

// inject repository

func InjectUserRepository() usecase.UserRepositoryInterface {
	return repository.NewUserRepository()
}

func InjectArticleRepository() usecase.ArticleRepositoryInterface {
	return repository.NewArticleRepository()
}

// inject util

func InjectUUIDUtil() usecase.UUIDUtilInterface {
	return util.NewGoogleUUID()
}

func InjectTimeUtil() usecase.TimeUtilInterface {
	return util.NewTimeUtil()
}

// inject usecase

func InjectAuthenticator() domain.AuthenticatorInterface {
	return usecase.NewMyscrapAuthenticator(
		InjectDB(),
		InjectUserRepository(),
		InjectUUIDUtil(),
		InjectTimeUtil(),
	)
}

func InjectClipper() domain.ClipperInterface {
	return usecase.NewClipper(
		InjectDB(),
		InjectArticleRepository(),
		InjectTimeUtil(),
	)
}

// inject controller

func InjectAuthenticationController() usecase.AuthenticationControllerInterface {
	return controller.NewAuthenticationController(
		InjectAuthenticator(),
	)
}

func InjectClipController() usecase.ClipControllerInteface {
	return controller.NewClipController(
		InjectClipper(),
	)
}
