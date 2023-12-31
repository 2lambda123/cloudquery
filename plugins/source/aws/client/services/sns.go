// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

//go:generate mockgen -package=mocks -destination=../mocks/sns.go -source=sns.go SnsClient
type SnsClient interface {
	GetDataProtectionPolicy(context.Context, *sns.GetDataProtectionPolicyInput, ...func(*sns.Options)) (*sns.GetDataProtectionPolicyOutput, error)
	GetEndpointAttributes(context.Context, *sns.GetEndpointAttributesInput, ...func(*sns.Options)) (*sns.GetEndpointAttributesOutput, error)
	GetPlatformApplicationAttributes(context.Context, *sns.GetPlatformApplicationAttributesInput, ...func(*sns.Options)) (*sns.GetPlatformApplicationAttributesOutput, error)
	GetSMSAttributes(context.Context, *sns.GetSMSAttributesInput, ...func(*sns.Options)) (*sns.GetSMSAttributesOutput, error)
	GetSMSSandboxAccountStatus(context.Context, *sns.GetSMSSandboxAccountStatusInput, ...func(*sns.Options)) (*sns.GetSMSSandboxAccountStatusOutput, error)
	GetSubscriptionAttributes(context.Context, *sns.GetSubscriptionAttributesInput, ...func(*sns.Options)) (*sns.GetSubscriptionAttributesOutput, error)
	GetTopicAttributes(context.Context, *sns.GetTopicAttributesInput, ...func(*sns.Options)) (*sns.GetTopicAttributesOutput, error)
	ListEndpointsByPlatformApplication(context.Context, *sns.ListEndpointsByPlatformApplicationInput, ...func(*sns.Options)) (*sns.ListEndpointsByPlatformApplicationOutput, error)
	ListOriginationNumbers(context.Context, *sns.ListOriginationNumbersInput, ...func(*sns.Options)) (*sns.ListOriginationNumbersOutput, error)
	ListPhoneNumbersOptedOut(context.Context, *sns.ListPhoneNumbersOptedOutInput, ...func(*sns.Options)) (*sns.ListPhoneNumbersOptedOutOutput, error)
	ListPlatformApplications(context.Context, *sns.ListPlatformApplicationsInput, ...func(*sns.Options)) (*sns.ListPlatformApplicationsOutput, error)
	ListSMSSandboxPhoneNumbers(context.Context, *sns.ListSMSSandboxPhoneNumbersInput, ...func(*sns.Options)) (*sns.ListSMSSandboxPhoneNumbersOutput, error)
	ListSubscriptions(context.Context, *sns.ListSubscriptionsInput, ...func(*sns.Options)) (*sns.ListSubscriptionsOutput, error)
	ListSubscriptionsByTopic(context.Context, *sns.ListSubscriptionsByTopicInput, ...func(*sns.Options)) (*sns.ListSubscriptionsByTopicOutput, error)
	ListTagsForResource(context.Context, *sns.ListTagsForResourceInput, ...func(*sns.Options)) (*sns.ListTagsForResourceOutput, error)
	ListTopics(context.Context, *sns.ListTopicsInput, ...func(*sns.Options)) (*sns.ListTopicsOutput, error)
}
