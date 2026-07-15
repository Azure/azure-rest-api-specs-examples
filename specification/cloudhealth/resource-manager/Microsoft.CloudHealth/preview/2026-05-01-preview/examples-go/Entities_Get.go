package armcloudhealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cloudhealth/armcloudhealth"
)

// Generated from example definition: 2026-05-01-preview/Entities_Get.json
func ExampleEntitiesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcloudhealth.NewClientFactory("abcdef12-3456-7890-abcd-ef1234567890", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEntitiesClient().Get(ctx, "online-store-rg", "online-store", "orders-db", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcloudhealth.EntitiesClientGetResponse{
	// 	Entity: armcloudhealth.Entity{
	// 		Properties: &armcloudhealth.EntityProperties{
	// 			DisplayName: to.Ptr("Orders Database"),
	// 			CanvasPosition: &armcloudhealth.EntityCoordinates{
	// 				X: to.Ptr[float32](480),
	// 				Y: to.Ptr[float32](360),
	// 			},
	// 			Icon: &armcloudhealth.IconDefinition{
	// 				IconName: to.Ptr("SQLDatabase"),
	// 			},
	// 			HealthObjective: to.Ptr[float32](99.5),
	// 			Impact: to.Ptr(armcloudhealth.EntityImpactStandard),
	// 			Tags: map[string]*string{
	// 				"environment": to.Ptr("production"),
	// 				"team": to.Ptr("online-store"),
	// 			},
	// 			SignalGroups: &armcloudhealth.SignalGroups{
	// 				AzureResource: &armcloudhealth.AzureResourceSignals{
	// 					AuthenticationSetting: to.Ptr("default-auth"),
	// 					AzureResourceID: to.Ptr("/subscriptions/abcdef12-3456-7890-abcd-ef1234567890/resourceGroups/online-store-rg/providers/Microsoft.Sql/servers/online-store-sql/databases/orders"),
	// 					Signals: []*armcloudhealth.AzureResourceSignal{
	// 						{
	// 							SignalKind: to.Ptr(armcloudhealth.SignalKindAzureResourceMetric),
	// 							Name: to.Ptr("sql-cpu"),
	// 							SignalDefinitionName: to.Ptr("sql-cpu-percent"),
	// 							Status: &armcloudhealth.SignalStatus{
	// 								HealthState: to.Ptr(armcloudhealth.HealthStateHealthy),
	// 								Value: to.Ptr[float64](38.5),
	// 								ReportedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T09:30:00.000Z"); return t}()),
	// 							},
	// 							DisplayName: to.Ptr("CPU utilization"),
	// 							RefreshInterval: to.Ptr(armcloudhealth.RefreshIntervalPT1M),
	// 							DataUnit: to.Ptr("Percent"),
	// 							MetricNamespace: to.Ptr("Microsoft.Sql/servers/databases"),
	// 							MetricName: to.Ptr("cpu_percent"),
	// 							TimeGrain: to.Ptr("PT5M"),
	// 							AggregationType: to.Ptr(armcloudhealth.MetricAggregationTypeAverage),
	// 							EvaluationRules: &armcloudhealth.EvaluationRule{
	// 								DegradedRule: &armcloudhealth.ThresholdRuleV2{
	// 									Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
	// 									Threshold: to.Ptr[float64](70),
	// 								},
	// 								UnhealthyRule: &armcloudhealth.ThresholdRuleV2{
	// 									Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
	// 									Threshold: to.Ptr[float64](90),
	// 								},
	// 							},
	// 						},
	// 						{
	// 							SignalKind: to.Ptr(armcloudhealth.SignalKindAzureResourceMetric),
	// 							Name: to.Ptr("sql-dtu"),
	// 							Status: &armcloudhealth.SignalStatus{
	// 								HealthState: to.Ptr(armcloudhealth.HealthStateHealthy),
	// 								Value: to.Ptr[float64](52.1),
	// 								ReportedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T09:30:00.000Z"); return t}()),
	// 							},
	// 							DisplayName: to.Ptr("DTU consumption"),
	// 							RefreshInterval: to.Ptr(armcloudhealth.RefreshIntervalPT1M),
	// 							DataUnit: to.Ptr("Percent"),
	// 							MetricNamespace: to.Ptr("Microsoft.Sql/servers/databases"),
	// 							MetricName: to.Ptr("dtu_consumption_percent"),
	// 							TimeGrain: to.Ptr("PT5M"),
	// 							AggregationType: to.Ptr(armcloudhealth.MetricAggregationTypeAverage),
	// 							EvaluationRules: &armcloudhealth.EvaluationRule{
	// 								DegradedRule: &armcloudhealth.ThresholdRuleV2{
	// 									Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
	// 									Threshold: to.Ptr[float64](75),
	// 								},
	// 								UnhealthyRule: &armcloudhealth.ThresholdRuleV2{
	// 									Operator: to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
	// 									Threshold: to.Ptr[float64](90),
	// 								},
	// 							},
	// 						},
	// 					},
	// 					ResourceHealth: &armcloudhealth.AzureResourceHealthSignal{
	// 						Enabled: to.Ptr(armcloudhealth.ResourceHealthAvailabilityStateSignalBehaviorEnabled),
	// 						SignalName: to.Ptr("resourcehealth-availabilitystate"),
	// 						Status: &armcloudhealth.AzureResourceHealthSignalStatus{
	// 							HealthState: to.Ptr(armcloudhealth.HealthStateHealthy),
	// 							ReportedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T09:30:00.000Z"); return t}()),
	// 							AvailabilityState: to.Ptr(armcloudhealth.ResourceHealthAvailabilityStateAvailable),
	// 							Summary: to.Ptr("The database is available."),
	// 						},
	// 					},
	// 				},
	// 				Dependencies: &armcloudhealth.DependenciesSignalGroupV2{
	// 					AggregationType: to.Ptr(armcloudhealth.DependenciesAggregationTypeWorstOf),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armcloudhealth.HealthModelProvisioningStateSucceeded),
	// 			HealthState: to.Ptr(armcloudhealth.HealthStateHealthy),
	// 		},
	// 		ID: to.Ptr("/subscriptions/abcdef12-3456-7890-abcd-ef1234567890/resourceGroups/online-store-rg/providers/Microsoft.CloudHealth/healthmodels/online-store/entities/orders-db"),
	// 		Name: to.Ptr("orders-db"),
	// 		Type: to.Ptr("Microsoft.CloudHealth/healthmodels/entities"),
	// 		SystemData: &armcloudhealth.SystemData{
	// 			CreatedBy: to.Ptr("admin@contoso.com"),
	// 			CreatedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T08:15:00.000Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("admin@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T09:30:00.000Z"); return t}()),
	// 		},
	// 	},
	// }
}
