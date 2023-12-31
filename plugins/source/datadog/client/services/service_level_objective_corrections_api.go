// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"net/http"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

//go:generate mockgen -package=mocks -destination=../mocks/service_level_objective_corrections_api.go -source=service_level_objective_corrections_api.go ServiceLevelObjectiveCorrectionsAPIClient
type ServiceLevelObjectiveCorrectionsAPIClient interface {
	ListSLOCorrection(context.Context, ...datadogV1.ListSLOCorrectionOptionalParameters) (datadogV1.SLOCorrectionListResponse, *http.Response, error)
}
