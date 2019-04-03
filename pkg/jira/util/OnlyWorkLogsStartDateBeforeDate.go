package util

import (
	"github.com/araddon/dateparse"
	"github.com/bmaximilian/gutils/pkg/jira/issues/models"
)

func OnlyWorkLogsStartDateBeforeDate(workLogReportItems []models.WorkLogReportItem, dateString string) ([]models.WorkLogReportItem, error) {
	parsedDate, parseErr := dateparse.ParseAny(dateString)
	if parseErr != nil {
		return workLogReportItems, parseErr
	}

	return FilterWorklogs(workLogReportItems, func(item models.WorkLogReportItem) bool {
		return item.StartedDate.Before(parsedDate) || item.StartedDate.Equal(parsedDate)
	}), nil
}
