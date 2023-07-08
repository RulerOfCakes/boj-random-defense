package solvedac

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	API_URL = "https://solved.ac/api/v3"
)

type SolvedAC_SortType string

const (
	SolvedAC_SortType_ID         SolvedAC_SortType = "id"
	SolvedAC_SortType_Level      SolvedAC_SortType = "level"
	SolvedAC_SortType_Title      SolvedAC_SortType = "title"
	SolvedAC_SortType_Solved     SolvedAC_SortType = "solved"
	SolvedAC_SortType_AverageTry SolvedAC_SortType = "average_try"
	SolvedAC_SortType_Random     SolvedAC_SortType = "random"
)

var (
	SolvedACProblemLevelToTitle = map[int32]string{
		0:  "Unrated",
		1:  "Bronze V",
		2:  "Bronze IV",
		3:  "Bronze III",
		4:  "Bronze II",
		5:  "Bronze I",
		6:  "Silver V",
		7:  "Silver IV",
		8:  "Silver III",
		9:  "Silver II",
		10: "Silver I",
		11: "Gold V",
		12: "Gold IV",
		13: "Gold III",
		14: "Gold II",
		15: "Gold I",
		16: "Platinum V",
		17: "Platinum IV",
		18: "Platinum III",
		19: "Platinum II",
		20: "Platinum I",
		21: "Diamond V",
		22: "Diamond IV",
		23: "Diamond III",
		24: "Diamond II",
		25: "Diamond I",
		26: "Ruby V",
		27: "Ruby IV",
		28: "Ruby III",
		29: "Ruby II",
		30: "Ruby I",
	}
)

type SolvedACTitle struct {
	Language            string `json:"language"`
	LanguageDisplayName string `json:"languageDisplayName"`
	Title               string `json:"title"`
	IsOriginal          bool   `json:"isOriginal"`
}

type SolvedACProblem struct {
	ProblemId         int32           `json:"problemId"`
	TitleKo           string          `json:"titleKo"`
	Titles            []SolvedACTitle `json:"titles"`
	IsSolvable        bool            `json:"isSolvable"`
	IsPartial         bool            `json:"isPartial"`
	AccepteduserCount int32           `json:"acceptedUserCount"`
	Level             int32           `json:"level"`
	VotedUserCount    int32           `json:"votedUserCount"`
	Sprout            bool            `json:"sprout"`
	GivesNoRating     bool            `json:"givesNoRating"`
	IsLevelLocked     bool            `json:"isLevelLocked"`
	AverageTries      float64         `json:"averageTries"`
	Official          bool            `json:"official"`
}

func GetProblems(query string, sort SolvedAC_SortType) ([]SolvedACProblem, error) {
	resource := "/search/problem"
	params := url.Values{}
	params.Add("query", query)
	params.Add("sort", string(sort))

	u, err := url.ParseRequestURI(API_URL)
	if err != nil {
		return nil, err
	}
	u.Path = resource
	u.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", u)

	resp, err := http.Get(urlStr)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ret struct {
		Count int32             `json:"count"`
		Items []SolvedACProblem `json:"items"`
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret.Items, nil
}
