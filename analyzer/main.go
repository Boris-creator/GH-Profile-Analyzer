package analyzer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/joho/godotenv"

	"github.com/Khan/genqlient/graphql"
)

type DeveloperType int32

const (
	WORK DeveloperType = iota + 1
	OPENSOURCE
	HOBBY
)

var ErrNotEnoughData = errors.New("too few activities for a reliable analysis")

type OwnProfileInfo struct {
	ContributionsDispersion float64
	Type                    DeveloperType
}

type Weeks getViewerViewerUserContributionsCollectionContributionCalendarWeeksContributionCalendarWeek

type authedTransport struct {
	key     string
	wrapped http.RoundTripper
}

func (t *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "bearer "+t.key)
	return t.wrapped.RoundTrip(req)
}

func init() {
	godotenv.Load(".env")
}

func getProfile(token string) (*getViewerResponse, error) {
	httpClient := http.Client{
		Transport: &authedTransport{
			key:     token,
			wrapped: http.DefaultTransport,
		},
	}
	graphqlClient := graphql.NewClient("https://api.github.com/graphql", &httpClient)

	to := time.Now()
	from := to.AddDate(-1, 0, 0)
	return getViewer(context.Background(), graphqlClient, from, to)
}

func ProfileInfo(token string) (OwnProfileInfo, error) {
	var err error
	defer func() {
		if err != nil {
			log.Println(err)
		}
	}()

	viewerResp, err := getProfile(token)
	if err != nil {
		return OwnProfileInfo{}, err
	}

	weeks := viewerResp.Viewer.GetContributionsCollection().ContributionCalendar.Weeks
	totalContributions := viewerResp.Viewer.ContributionsCollection.ContributionCalendar.TotalContributions

	var threshold = 1.5
	var minimumAmount = 10

	if totalContributions < minimumAmount {
		return OwnProfileInfo{}, ErrNotEnoughData
	}

	var sunday, monday, tuesday, wednesday, thursday, friday, saturday int
	weeksContributions := make([]int, 0, len(weeks))
	for _, week := range weeks {
		totalWeekCount := 0
		for _, day := range week.ContributionDays {
			switch day.Weekday {
			case 0:
				sunday += day.ContributionCount
			case 1:
				monday += day.ContributionCount
			case 2:
				tuesday += day.ContributionCount
			case 3:
				wednesday += day.ContributionCount
			case 4:
				thursday += day.ContributionCount
			case 5:
				friday += day.ContributionCount
			case 6:
				saturday += day.ContributionCount
			}
			totalWeekCount += day.ContributionCount
		}

		weeksContributions = append(weeksContributions, totalWeekCount)
	}
	averageWorkDayContributionsCount := float64(monday+tuesday+wednesday+thursday+friday) / float64(5*len(weeks))
	averageRestDayContributionsCount := float64(sunday+saturday) / float64(2*len(weeks))

	_, _ = json.MarshalIndent(viewerResp.Viewer.Repositories.Nodes, "", "  ")

	if len(viewerResp.Viewer.Repositories.Nodes) > 0 {
		repo := viewerResp.Viewer.Repositories.Nodes[0]
		if repo.DefaultBranchRef.Target.GetTypename() == "Commit" {
			b := getViewerViewerUserRepositoriesRepositoryConnectionNodesRepositoryDefaultBranchRefTargetGitObject(repo.DefaultBranchRef.Target).(*getViewerViewerUserRepositoriesRepositoryConnectionNodesRepositoryDefaultBranchRefTargetCommit)
			fmt.Println(b.History.Nodes)
		}
	}

	//fmt.Println(viewerResp.Viewer.MyName)

	var developerType DeveloperType

	if averageRestDayContributionsCount == 0 {
		developerType = WORK
	} else {
		switch difference := averageWorkDayContributionsCount / averageRestDayContributionsCount; {
		case difference > threshold:
			developerType = WORK
		case difference < 1/threshold:
			developerType = HOBBY
		default:
			developerType = OPENSOURCE
		}
	}

	return OwnProfileInfo{
		ContributionsDispersion: getDeviation(weeksContributions),
		Type:                    developerType,
	}, nil
}

func getAverage[N int | float64](values []N, mapper *(func(N) N)) float64 {
	m := func(v N) N { return v }
	if mapper != nil {
		m = *mapper
	}
	var sum N = 0
	for _, v := range values {
		sum += m(v)
	}
	return float64(sum) / float64(len(values))
}
func getDeviation[N int | float64](values []N) float64 {
	pow := func(v N) N { return N(math.Pow(float64(v), 2)) }
	dispersion := getAverage(values, &pow) - math.Pow(getAverage(values, nil), 2)
	return math.Sqrt(dispersion)
}

//go:generate go run github.com/Khan/genqlient genqlient.yaml
