package analyzer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"slices"
	"sync"
	"time"

	"github.com/joho/godotenv"

	"github.com/Khan/genqlient/graphql"
	"github.com/go-enry/go-enry/v2"
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
	Languages               []string
}

type Weeks getViewerViewerUserContributionsCollectionContributionCalendarWeeksContributionCalendarWeek

type commit struct {
	Author string
	Repo   string
	Hash   string
}

type commitInfo struct {
	Files []struct {
		Filename string
	}
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

	totalContributions := viewerResp.Viewer.ContributionsCollection.ContributionCalendar.TotalContributions

	var minimumAmount = 10

	if totalContributions < minimumAmount {
		return OwnProfileInfo{}, ErrNotEnoughData
	}

	developerType, contributionsDispersion, err := getContributionsActivity(viewerResp.Viewer)

	commitsLanguages, err := getCommitsLanguages(viewerResp.Viewer, token)
	if err != nil {
		return OwnProfileInfo{}, err
	}

	return OwnProfileInfo{
		ContributionsDispersion: contributionsDispersion,
		Type:                    developerType,
		Languages:               commitsLanguages,
	}, nil
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

	t, _ := getMyId(context.Background(), graphqlClient)

	return getViewer(context.Background(), graphqlClient, t.Viewer.Id, from, to, to.AddDate(-3, 0, 0))
}

func getCommitByHash(userName string, repo, hash, token string) (commitInfo, error) {
	httpClient := http.Client{
		Transport: &authedTransport{
			key:     token,
			wrapped: http.DefaultTransport,
		},
	}
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/repos/%s/%s/commits/%s", userName, repo, hash), nil)

	var commit commitInfo
	response, err := httpClient.Do(req)
	if err != nil {
		return commit, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&commit)
	return commit, err
}

func getCommitsLanguages(viewer getViewerViewerUser, token string) ([]string, error) {
	wg := sync.WaitGroup{}
	var mu sync.Mutex
	var commits []commit

	var ignoredLangs = [9]string{
		"XML",
		"YAML",
		"JSON",
		"TOML",
		"Dotenv",
		"Ignore List",
		"Adblock Filter List",
		"Git Config",
		"EditorConfig",
	}

	for _, repo := range viewer.Repositories.Nodes {
		if repo.DefaultBranchRef.Target != nil && repo.DefaultBranchRef.Target.GetTypename() == "Commit" {
			b := getViewerViewerUserRepositoriesRepositoryConnectionNodesRepositoryDefaultBranchRefTargetGitObject(repo.DefaultBranchRef.Target).(*getViewerViewerUserRepositoriesRepositoryConnectionNodesRepositoryDefaultBranchRefTargetCommit)
			if b == nil {
				continue
			}
			for _, node := range b.History.Nodes {
				commits = append(commits, commit{
					viewer.Login,
					repo.Name,
					node.Oid,
				})
			}
		}
	}

	wg.Add(len(commits))

	var commitsLangs = make(map[string]int)
	for _, c := range commits {
		go func() {
			defer wg.Done()

			res, _ := getCommitByHash(c.Author, c.Repo, c.Hash, token)

			commitLangs := make(map[string]bool)
			for _, filename := range res.Files {
				lang, safe := enry.GetLanguageByExtension(filename.Filename)
				if !safe || slices.Contains(ignoredLangs[:], lang) {
					continue
				}
				// We do not take into account the number of files changed,
				// as it may depend on the project structure and does not reflect the actual amount of changes
				commitLangs[lang] = true
			}

			mu.Lock()
			for lang := range commitLangs {
				commitsLangs[lang]++
			}
			mu.Unlock()
		}()
	}
	wg.Wait()

	return topN(10, commitsLangs), nil
}

func getContributionsActivity(viewer getViewerViewerUser) (DeveloperType, float64, error) {
	weeks := viewer.GetContributionsCollection().ContributionCalendar.Weeks

	var threshold = 1.5

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

	return developerType, getDeviation(weeksContributions), nil
}
