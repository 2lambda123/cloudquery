// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness"
)

//go:generate mockgen -package=mocks -destination=../mocks/route53recoveryreadiness.go -source=route53recoveryreadiness.go Route53recoveryreadinessClient
type Route53recoveryreadinessClient interface {
	GetArchitectureRecommendations(context.Context, *route53recoveryreadiness.GetArchitectureRecommendationsInput, ...func(*route53recoveryreadiness.Options)) (*route53recoveryreadiness.GetArchitectureRecommendationsOutput, error)
	GetCell(context.Context, *route53recoveryreadiness.GetCellInput, ...func(*route53recoveryreadiness.Options)) (*route53recoveryreadiness.GetCellOutput, error)
	GetCellReadinessSummary(context.Context, *route53recoveryreadiness.GetCellReadinessSummaryInput, ...func(*route53recoveryreadiness.Options)) (*route53recoveryreadiness.GetCellReadinessSummaryOutput, error)
	GetReadinessCheck(context.Context, *route53recoveryreadiness.GetReadinessCheckInput, ...func(*route53recoveryreadiness.Options)) (*route53recoveryreadiness.GetReadinessCheckOutput, error)
	GetReadinessCheckResourceStatus(context.Context, *route53recoveryreadiness.GetReadinessCheckResourceStatusInput, ...func(*route53recoveryreadiness.Options)) (*route53recoveryreadiness.GetReadinessCheckResourceStatusOutput, error)
	GetReadinessCheckStatus(context.Context, *route53recoveryreadiness.GetReadinessCheckStatusInput, ...func(*route53recoveryreadiness.Options)) (*route53recoveryreadiness.GetReadinessCheckStatusOutput, error)
	GetRecoveryGroup(context.Context, *route53recoveryreadiness.GetRecoveryGroupInput, ...func(*route53recoveryreadiness.Options)) (*route53recoveryreadiness.GetRecoveryGroupOutput, error)
	GetRecoveryGroupReadinessSummary(context.Context, *route53recoveryreadiness.GetRecoveryGroupReadinessSummaryInput, ...func(*route53recoveryreadiness.Options)) (*route53recoveryreadiness.GetRecoveryGroupReadinessSummaryOutput, error)
	GetResourceSet(context.Context, *route53recoveryreadiness.GetResourceSetInput, ...func(*route53recoveryreadiness.Options)) (*route53recoveryreadiness.GetResourceSetOutput, error)
	ListCells(context.Context, *route53recoveryreadiness.ListCellsInput, ...func(*route53recoveryreadiness.Options)) (*route53recoveryreadiness.ListCellsOutput, error)
	ListCrossAccountAuthorizations(context.Context, *route53recoveryreadiness.ListCrossAccountAuthorizationsInput, ...func(*route53recoveryreadiness.Options)) (*route53recoveryreadiness.ListCrossAccountAuthorizationsOutput, error)
	ListReadinessChecks(context.Context, *route53recoveryreadiness.ListReadinessChecksInput, ...func(*route53recoveryreadiness.Options)) (*route53recoveryreadiness.ListReadinessChecksOutput, error)
	ListRecoveryGroups(context.Context, *route53recoveryreadiness.ListRecoveryGroupsInput, ...func(*route53recoveryreadiness.Options)) (*route53recoveryreadiness.ListRecoveryGroupsOutput, error)
	ListResourceSets(context.Context, *route53recoveryreadiness.ListResourceSetsInput, ...func(*route53recoveryreadiness.Options)) (*route53recoveryreadiness.ListResourceSetsOutput, error)
	ListRules(context.Context, *route53recoveryreadiness.ListRulesInput, ...func(*route53recoveryreadiness.Options)) (*route53recoveryreadiness.ListRulesOutput, error)
	ListTagsForResources(context.Context, *route53recoveryreadiness.ListTagsForResourcesInput, ...func(*route53recoveryreadiness.Options)) (*route53recoveryreadiness.ListTagsForResourcesOutput, error)
}
