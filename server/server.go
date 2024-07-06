package main

import (
	"fmt"
	"log"
	"os"

	pb "ghs/proto"
	"ghs/server/handlers"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println(err.Error())
	}

	client, err := StartClient()
	if err != nil {
		log.Println(err.Error())
	}

	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))))

	e.Use(func(pb *pb.GHAnalysisServiceClient) echo.MiddlewareFunc {
		return func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				c.Set("client", pb)
				return next(c)
			}
		}
	}(&client))

	e.File("/", "../public/index.html")
	e.POST("/analyze", handlers.GetOwnProfileInfo)
	e.GET("/login", handlers.OauthGHLogin)
	e.GET("/login1", handlers.OauthGHLoginExtended)
	e.GET("/auth/gh/callback", handlers.OauthGHCallback)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))))
}

func StartClient() (pb.GHAnalysisServiceClient, error) {
	conn, err := grpc.NewClient(
		fmt.Sprintf("localhost:%s", os.Getenv("SERVICE_PORT")),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	return pb.NewGHAnalysisServiceClient(conn), nil
}
