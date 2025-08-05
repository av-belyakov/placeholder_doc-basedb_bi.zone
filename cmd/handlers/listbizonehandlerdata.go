package handlers

import "github.com/av-belyakov/placeholder-doc-basedb-bi-zone/internal/datamodels"

// NewListBiZoneHandlerData начальный обработчик событий топика 'data.*'
func NewListBiZoneHandlerData(data *datamodels.BiZoneData) map[string][]func(any) {
	return map[string][]func(any){
		"data.id":                   {data.SetAnyIDNum},
		"data._id":                  {data.SetAnyID},
		"data.uuid":                 {data.SetAnyUUID},
		"data.title":                {data.SetAnyTitle},
		"data.sensor":               {data.SetAnySensor},
		"data.ip_home":              {data.SetAnyIPHome},
		"data.ip_exter":             {data.SetAnyIPExter},
		"data.url____ftp":           {data.SetAnyURLFTP},
		"data.url___http":           {data.SetAnyURLHTTP},
		"data.event_type":           {data.SetAnyEventType},
		"data.url_arkime":           {data.SetAnyURLArkime},
		"data.confidence":           {data.SetAnyConfidence},
		"data.description":          {data.SetAnyDescription},
		"data.external_id":          {data.SetAnyExternalID},
		"data.created_time":         {data.SetAnyCreatedTime},
		"data.platform_type":        {data.SetAnyPlatformType},
		"data.all____ip_ext":        {data.SetAnyAllIPExt},
		"data.response_team":        {data.SetAnyResponseTeam},
		"data.detection_rule":       {data.SetAnyDetectionRule},
		"data.customer_system":      {data.SetAnyCustomerSystem},
		"data.event_start_time":     {data.SetAnyEventStartTime},
		"data.last_detection_time":  {data.SetAnyLastDetectionTime},
		"data.affected_log_sources": {data.SetAnyAffectedLogSources},
		"data.first_detection_time": {data.SetAnyFirstDetectionTime},
		//ниже работа со срезам содержащими простые типы
		"data.snort_sid":    {data.SetAnySnortSid},
		"data.all_sensors":  {data.SetAnyAllSensor},
		"data.all__ip_home": {data.SetAnyAllIPHome},
	}
}
