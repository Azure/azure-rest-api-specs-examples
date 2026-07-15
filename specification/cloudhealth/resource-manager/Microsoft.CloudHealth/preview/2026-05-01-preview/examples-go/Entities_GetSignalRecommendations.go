package armcloudhealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cloudhealth/armcloudhealth"
)

// Generated from example definition: 2026-05-01-preview/Entities_GetSignalRecommendations.json
func ExampleEntitiesClient_GetSignalRecommendations() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcloudhealth.NewClientFactory("abcdef12-3456-7890-abcd-ef1234567890", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEntitiesClient().GetSignalRecommendations(ctx, "online-store-rg", "online-store", "orders-db", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcloudhealth.EntitiesClientGetSignalRecommendationsResponse{
	// 	GetSignalRecommendationsResponse: armcloudhealth.GetSignalRecommendationsResponse{
	// 		RecommendedSignals: []*armcloudhealth.SignalConfiguration{
	// 			{
	// 				SignalID: to.Ptr("sql-cpu-percent"),
	// 				MetricNamespace: to.Ptr("Microsoft.Sql/servers/databases"),
	// 				MetricName: to.Ptr("cpu_percent"),
	// 				AggregationType: to.Ptr(armcloudhealth.MetricAggregationTypeAverage),
	// 				Unit: to.Ptr("Percent"),
	// 				TimeGrain: to.Ptr("PT5M"),
	// 				EvaluationRules: &armcloudhealth.EvaluationRule{
	// 					DegradedRule: &armcloudhealth.ThresholdRuleV2{
	// 						Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
	// 						Threshold: to.Ptr[float64](70),
	// 					},
	// 					UnhealthyRule: &armcloudhealth.ThresholdRuleV2{
	// 						Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
	// 						Threshold: to.Ptr[float64](90),
	// 					},
	// 				},
	// 			},
	// 			{
	// 				SignalID: to.Ptr("sql-dtu-consumption"),
	// 				MetricNamespace: to.Ptr("Microsoft.Sql/servers/databases"),
	// 				MetricName: to.Ptr("dtu_consumption_percent"),
	// 				AggregationType: to.Ptr(armcloudhealth.MetricAggregationTypeAverage),
	// 				Unit: to.Ptr("Percent"),
	// 				TimeGrain: to.Ptr("PT5M"),
	// 				EvaluationRules: &armcloudhealth.EvaluationRule{
	// 					DegradedRule: &armcloudhealth.ThresholdRuleV2{
	// 						Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
	// 						Threshold: to.Ptr[float64](75),
	// 					},
	// 					UnhealthyRule: &armcloudhealth.ThresholdRuleV2{
	// 						Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
	// 						Threshold: to.Ptr[float64](90),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		RecommendedConfigurations: []*armcloudhealth.SignalConfiguration{
	// 			{
	// 				SignalID: to.Ptr("sql-storage-percent"),
	// 				MetricNamespace: to.Ptr("Microsoft.Sql/servers/databases"),
	// 				MetricName: to.Ptr("storage_percent"),
	// 				AggregationType: to.Ptr(armcloudhealth.MetricAggregationTypeMaximum),
	// 				Unit: to.Ptr("Percent"),
	// 				TimeGrain: to.Ptr("PT15M"),
	// 				EvaluationRules: &armcloudhealth.EvaluationRule{
	// 					DegradedRule: &armcloudhealth.ThresholdRuleV2{
	// 						Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
	// 						Threshold: to.Ptr[float64](80),
	// 					},
	// 					UnhealthyRule: &armcloudhealth.ThresholdRuleV2{
	// 						Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
	// 						Threshold: to.Ptr[float64](95),
	// 					},
	// 				},
	// 			},
	// 			{
	// 				SignalID: to.Ptr("sql-deadlocks"),
	// 				MetricNamespace: to.Ptr("Microsoft.Sql/servers/databases"),
	// 				MetricName: to.Ptr("deadlock"),
	// 				AggregationType: to.Ptr(armcloudhealth.MetricAggregationTypeTotal),
	// 				Unit: to.Ptr("Count"),
	// 				TimeGrain: to.Ptr("PT5M"),
	// 				EvaluationRules: &armcloudhealth.EvaluationRule{
	// 					UnhealthyRule: &armcloudhealth.ThresholdRuleV2{
	// 						Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
	// 						Threshold: to.Ptr[float64](0),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
