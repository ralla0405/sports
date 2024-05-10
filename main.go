package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type LeagueTable struct {
	Table []struct {
		TeamName string `json:"strTeam"`
		Rank     string `json:"intRank"`
	} `json:"table"`
}

func getLeagueTable(leagueID string) (*LeagueTable, error) {
	url := fmt.Sprintf("https://www.thesportsdb.com/api/v1/json/3/lookuptable.php?l=%s&s=2023-2024", leagueID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var table LeagueTable
	err = json.Unmarshal(body, &table)
	if err != nil {
		return nil, err
	}

	return &table, nil
}

func main() {
	// 예시로 사용한 리그 ID는 시간에 따라 변경될 수 있습니다.
	// 정확한 리그 ID는 TheSportsDB의 API 문서를 참조하세요.
	//kLeagueID := "5984"       // K리그의 리그 ID 예시
	premierLeagueID := "4328" // 프리미어 리그의 리그 ID 예시

	//fmt.Println("K리그 순위:")
	//kLeagueTable, err := getLeagueTable(kLeagueID)
	//if err != nil {
	//	fmt.Println("Error fetching K-League rankings:", err)
	//	return
	//}
	//
	//for _, team := range kLeagueTable.Table {
	//	fmt.Printf("%s. %s\n", team.Rank, team.TeamName)
	//}

	fmt.Println("\n프리미어 리그 순위:")
	premierLeagueTable, err := getLeagueTable(premierLeagueID)
	if err != nil {
		fmt.Println("Error fetching Premier League rankings:", err)
		return
	}

	for _, team := range premierLeagueTable.Table {
		fmt.Printf("%s. %s\n", team.Rank, team.TeamName)
	}
}
