package handlers

import (
	"github.com/av-belyakov/placeholder-doc-basedb-bi-zone/internal/datamodels"
)

// NewListBiZoneHandlerAlerts начальный обработчик событий топика 'alerts'
func NewListBiZoneHandlerAlerts(alert *datamodels.VerifiedBiZoneAlert) map[string][]func(any) {
	return map[string][]func(any){
		"title":             {alert.SetAnyTitle},
		"severity":          {alert.SetAnySeverity},
		"updated_time":      {alert.SetAnyUpdatedTime},
		"event_end_time":    {alert.SetAnyEventEndTime},
		"recommendations":   {alert.SetAnyRecommendations},
		"platform_hostname": {alert.SetAnyPlatformHostname},
	}
}
