// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/savingsplans"
)

//go:generate mockgen -package=mocks -destination=../mocks/savingsplans.go -source=savingsplans.go SavingsplansClient
type SavingsplansClient interface {
	DescribeSavingsPlanRates(context.Context, *savingsplans.DescribeSavingsPlanRatesInput, ...func(*savingsplans.Options)) (*savingsplans.DescribeSavingsPlanRatesOutput, error)
	DescribeSavingsPlans(context.Context, *savingsplans.DescribeSavingsPlansInput, ...func(*savingsplans.Options)) (*savingsplans.DescribeSavingsPlansOutput, error)
	DescribeSavingsPlansOfferingRates(context.Context, *savingsplans.DescribeSavingsPlansOfferingRatesInput, ...func(*savingsplans.Options)) (*savingsplans.DescribeSavingsPlansOfferingRatesOutput, error)
	DescribeSavingsPlansOfferings(context.Context, *savingsplans.DescribeSavingsPlansOfferingsInput, ...func(*savingsplans.Options)) (*savingsplans.DescribeSavingsPlansOfferingsOutput, error)
	ListTagsForResource(context.Context, *savingsplans.ListTagsForResourceInput, ...func(*savingsplans.Options)) (*savingsplans.ListTagsForResourceOutput, error)
}
