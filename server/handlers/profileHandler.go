package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "ghs/proto"
	"ghs/server/resources"
)

func fromServiceResponse(r pb.OwnProfileResponse) resources.OwnProfileInfo {
	types := map[pb.EmploymentType]string{
		pb.EmploymentType_WORK:       "Work",
		pb.EmploymentType_HOBBY:      "Hobby",
		pb.EmploymentType_OPENSOURCE: "Opensource",
	}
	return resources.OwnProfileInfo{
		ContributionsDispersion: r.ContributionsDispersion,
		Type:                    types[r.Type],
		Languages:               r.Languages,
	}
}

func GetOwnProfileInfo(c echo.Context) error {
	client := c.Get("client").(*pb.GHAnalysisServiceClient)

	s, _ := session.Get("session", c)

	r, err := (*client).OwnProfileInfo(context.Background(), &pb.OwnProfileRequest{Token: s.Values["token"].(string)})

	if status.Code(err) == codes.InvalidArgument {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return c.JSON(http.StatusOK, fromServiceResponse(*r))
}
