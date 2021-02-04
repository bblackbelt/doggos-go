package main

import (
	"net/http"

	"doggos/domain"
	"doggos/domain/adapter"
	ihttp "doggos/executor"
	"doggos/middleware"
	"doggos/repo"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	executor := ihttp.NewHttpExecutor(ihttp.NewHttpClient(&http.Client{}))
	r := repo.NewDoggoRepository(executor, ihttp.NewHttpRequestFactory())

	usecase := domain.NewDoggoUseCase(r, adapter.NewDoggoMapper())
	middleware.RegisterDoggosHandler(e, usecase)

	e.Start(":9090")
}
