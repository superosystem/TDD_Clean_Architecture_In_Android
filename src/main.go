package main

import (
	"context"
	"github.com/labstack/echo/v4"
	_config "github.com/superosystem/bantumanten-backend/src/app/config"
	"github.com/superosystem/bantumanten-backend/src/app/routes"
	_dbDriver "github.com/superosystem/bantumanten-backend/src/drivers/mysql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type operation func(context.Context) error

func main() {
	// DATABASE SETUP
	cfgDB := _dbDriver.ConfigMySQL{
		MYSQL_USERNAME: _config.GetEnvValue("DB_USERNAME"),
		MYSQL_PASSWORD: _config.GetEnvValue("DB_PASSWORD"),
		MYSQL_HOST:     _config.GetEnvValue("DB_HOST"),
		MYSQL_PORT:     _config.GetEnvValue("DB_PORT"),
		MYSQL_NAME:     _config.GetEnvValue("DB_NAME"),
	}
	mysqlDB := cfgDB.InitMySQLDatabase()
	_dbDriver.MySQLAutoMigrate(mysqlDB)

	ctx := context.Background()

	// ECHO
	e := echo.New()

	route := routes.Config{
		Echo:      e,
		MySQLCONN: mysqlDB,
	}
	route.Start()

	go func() {
		if err := e.Start(_config.GetEnvValue("SERVER_PORT")); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()
	wait := gracefulShutdown(ctx, 5*time.Second, map[string]operation{
		"mysql": func(ctx context.Context) error {
			return _dbDriver.CloseDB(mysqlDB)
		},
		"http-server": func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})
	<-wait
}

// graceful shutdown perform application shutdown gracefully
func gracefulShutdown(ctx context.Context, timeout time.Duration, ops map[string]operation) <-chan struct{} {
	wait := make(chan struct{})

	go func() {
		s := make(chan os.Signal, 1)

		// add any other syscall that you want to be notified with
		signal.Notify(s, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)
		<-s

		log.Println("shutting down")

		// set timeout for the ops to be done to prevent system hang
		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Printf("timeout %d ms has been elapsed, force exit", timeout.Milliseconds())
			os.Exit(0)
		})

		defer timeoutFunc.Stop()

		var wg sync.WaitGroup

		// do the operation asynchronously to save time
		for key, op := range ops {
			wg.Add(1)
			innerOp := op
			innerKey := key
			go func() {
				defer wg.Done()

				log.Printf("cleaning up: %v", innerKey)

				if err := innerOp(ctx); err != nil {
					log.Printf("%s: clean up failed: %s", innerKey, err.Error())
					return
				}

				log.Printf("%s was shutdown gracefully", innerKey)
			}()
		}

		wg.Wait()

		close(wait)
	}()

	return wait
}
