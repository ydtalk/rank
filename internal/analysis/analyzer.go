package analysis

import (
	"sort"
)

type MatchStats struct {
	TotalGames     int     `json:"total_games"`
	WinRate        float64 `json:"win_rate"`
	AverageKDA     float64 `json:"average_kda"`
	AverageCS      float64 `json:"average_cs"`
	MostPlayedRole string  `json:"most_played_role"`
	BestChampion   string  `json:"best_champion"`
	WorstChampion  string  `json:"worst_champion"`
}

type ChampionStats struct {
	Name       string  `json:"name"`
	Games      int     `json:"games"`
	WinRate    float64 `json:"win_rate"`
	AverageKDA float64 `json:"average_kda"`
}

// AnalyzeMatches 分析比赛数据
func AnalyzeMatches(matches []map[string]interface{}) MatchStats {
	if len(matches) == 0 {
		return MatchStats{}
	}

	stats := MatchStats{
		TotalGames: len(matches),
	}

	// 计算胜率
	wins := 0
	totalKDA := 0.0
	totalCS := 0.0
	roleCount := make(map[string]int)
	championStats := make(map[string]ChampionStats)

	for _, match := range matches {
		// 统计胜负
		if match["win"] == true {
			wins++
		}

		// 统计 KDA
		kills := match["kills"].(float64)
		deaths := match["deaths"].(float64)
		assists := match["assists"].(float64)
		kda := (kills + assists) / deaths
		totalKDA += kda

		// 统计补刀
		cs := match["cs"].(float64)
		totalCS += cs

		// 统计位置
		role := match["role"].(string)
		roleCount[role]++

		// 统计英雄数据
		champion := match["champion"].(string)
		champStats := championStats[champion]
		champStats.Name = champion
		champStats.Games++
		if match["win"] == true {
			champStats.WinRate += 1
		}
		champStats.AverageKDA += kda
		championStats[champion] = champStats
	}

	// 计算平均值
	stats.WinRate = float64(wins) / float64(len(matches)) * 100
	stats.AverageKDA = totalKDA / float64(len(matches))
	stats.AverageCS = totalCS / float64(len(matches))

	// 找出最常玩的位置
	maxRoleCount := 0
	for role, count := range roleCount {
		if count > maxRoleCount {
			maxRoleCount = count
			stats.MostPlayedRole = role
		}
	}

	// 计算每个英雄的平均数据
	var championList []ChampionStats
	for _, stats := range championStats {
		stats.WinRate = stats.WinRate / float64(stats.Games) * 100
		stats.AverageKDA = stats.AverageKDA / float64(stats.Games)
		championList = append(championList, stats)
	}

	// 按胜率排序找出最佳和最差英雄
	sort.Slice(championList, func(i, j int) bool {
		return championList[i].WinRate > championList[j].WinRate
	})

	if len(championList) > 0 {
		stats.BestChampion = championList[0].Name
		stats.WorstChampion = championList[len(championList)-1].Name
	}

	return stats
}
