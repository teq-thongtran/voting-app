package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"time"

	"github.com/soheilhy/cmux"
	"gorm.io/gorm"

	"myapp/config"
	serviceHttp "myapp/http"
	"myapp/migration"
	"myapp/mysql"
	"myapp/repository"
	"myapp/usecase"
)

const VERSION = "1.0.0"

func main() {
	var (
		//cfg     = config.GetConfig()
		taskPtr = flag.String("task", "server", "server")
	)

	flag.Parse()

	// setup locale
	{
		loc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
		if err != nil {
			fmt.Println("ERROR LOCALE: ", err)
		}

		time.Local = loc
	}

	client := mysql.GetClient
	repo := repository.New(client)
	useCase := usecase.New(repo)

	switch *taskPtr {
	case "server":
		executeServer(useCase, repo, client)
	default:
		executeServer(useCase, repo, client)
	}
}

func executeServer(useCase *usecase.UseCase, repo *repository.Repository, client func(ctx context.Context) *gorm.DB) {
	cfg := config.GetConfig()

	// migration
	migration.Up(client(context.Background()))

	l, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		fmt.Println("ERROR LISTEN: ", err)
	}

	m := cmux.New(l)
	httpL := m.Match(cmux.HTTP1Fast())
	errs := make(chan error)

	// http
	{
		h := serviceHttp.NewHTTPHandler(useCase, repo)
		go func() {
			h.Listener = httpL
			errs <- h.Start("")
		}()
	}
	go func() {
		errs <- m.Serve()
	}()

	err = <-errs
	if err != nil {
		fmt.Println("MAIN ERRORS: ", err)
	}
}
