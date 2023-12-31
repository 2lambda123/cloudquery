// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

//go:generate mockgen -package=mocks -destination=../mocks/cloudwatch.go -source=cloudwatch.go CloudwatchClient
type CloudwatchClient interface {
	DescribeAlarmHistory(context.Context, *cloudwatch.DescribeAlarmHistoryInput, ...func(*cloudwatch.Options)) (*cloudwatch.DescribeAlarmHistoryOutput, error)
	DescribeAlarms(context.Context, *cloudwatch.DescribeAlarmsInput, ...func(*cloudwatch.Options)) (*cloudwatch.DescribeAlarmsOutput, error)
	DescribeAlarmsForMetric(context.Context, *cloudwatch.DescribeAlarmsForMetricInput, ...func(*cloudwatch.Options)) (*cloudwatch.DescribeAlarmsForMetricOutput, error)
	DescribeAnomalyDetectors(context.Context, *cloudwatch.DescribeAnomalyDetectorsInput, ...func(*cloudwatch.Options)) (*cloudwatch.DescribeAnomalyDetectorsOutput, error)
	DescribeInsightRules(context.Context, *cloudwatch.DescribeInsightRulesInput, ...func(*cloudwatch.Options)) (*cloudwatch.DescribeInsightRulesOutput, error)
	GetDashboard(context.Context, *cloudwatch.GetDashboardInput, ...func(*cloudwatch.Options)) (*cloudwatch.GetDashboardOutput, error)
	GetInsightRuleReport(context.Context, *cloudwatch.GetInsightRuleReportInput, ...func(*cloudwatch.Options)) (*cloudwatch.GetInsightRuleReportOutput, error)
	GetMetricData(context.Context, *cloudwatch.GetMetricDataInput, ...func(*cloudwatch.Options)) (*cloudwatch.GetMetricDataOutput, error)
	GetMetricStatistics(context.Context, *cloudwatch.GetMetricStatisticsInput, ...func(*cloudwatch.Options)) (*cloudwatch.GetMetricStatisticsOutput, error)
	GetMetricStream(context.Context, *cloudwatch.GetMetricStreamInput, ...func(*cloudwatch.Options)) (*cloudwatch.GetMetricStreamOutput, error)
	GetMetricWidgetImage(context.Context, *cloudwatch.GetMetricWidgetImageInput, ...func(*cloudwatch.Options)) (*cloudwatch.GetMetricWidgetImageOutput, error)
	ListDashboards(context.Context, *cloudwatch.ListDashboardsInput, ...func(*cloudwatch.Options)) (*cloudwatch.ListDashboardsOutput, error)
	ListManagedInsightRules(context.Context, *cloudwatch.ListManagedInsightRulesInput, ...func(*cloudwatch.Options)) (*cloudwatch.ListManagedInsightRulesOutput, error)
	ListMetricStreams(context.Context, *cloudwatch.ListMetricStreamsInput, ...func(*cloudwatch.Options)) (*cloudwatch.ListMetricStreamsOutput, error)
	ListMetrics(context.Context, *cloudwatch.ListMetricsInput, ...func(*cloudwatch.Options)) (*cloudwatch.ListMetricsOutput, error)
	ListTagsForResource(context.Context, *cloudwatch.ListTagsForResourceInput, ...func(*cloudwatch.Options)) (*cloudwatch.ListTagsForResourceOutput, error)
}
