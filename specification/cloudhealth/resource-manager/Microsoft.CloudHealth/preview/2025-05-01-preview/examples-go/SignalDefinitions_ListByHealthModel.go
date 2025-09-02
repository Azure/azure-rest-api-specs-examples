package armcloudhealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cloudhealth/armcloudhealth"
)

// Generated from example definition: 2025-05-01-preview/SignalDefinitions_ListByHealthModel.json
func ExampleSignalDefinitionsClient_NewListByHealthModelPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcloudhealth.NewClientFactory("4980D7D5-4E07-47AD-AD34-E76C6BC9F061", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSignalDefinitionsClient().NewListByHealthModelPager("rgopenapi", "myHealthModel", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armcloudhealth.SignalDefinitionsClientListByHealthModelResponse{
		// 	SignalDefinitionListResult: armcloudhealth.SignalDefinitionListResult{
		// 		Value: []*armcloudhealth.SignalDefinition{
		// 			{
		// 				Properties: &armcloudhealth.ResourceMetricSignalDefinitionProperties{
		// 					MetricNamespace: to.Ptr("microsoft.compute/virtualMachines"),
		// 					MetricName: to.Ptr("cpuusage"),
		// 					AggregationType: to.Ptr(armcloudhealth.MetricAggregationTypeNone),
		// 					Dimension: to.Ptr("node"),
		// 					DimensionFilter: to.Ptr("node1"),
		// 					ProvisioningState: to.Ptr(armcloudhealth.HealthModelProvisioningStateSucceeded),
		// 					DisplayName: to.Ptr("cpu usage"),
		// 					SignalKind: to.Ptr(armcloudhealth.SignalKindAzureResourceMetric),
		// 					RefreshInterval: to.Ptr(armcloudhealth.RefreshIntervalPT1M),
		// 					Labels: map[string]*string{
		// 						"key4788": to.Ptr("bar"),
		// 					},
		// 					TimeGrain: to.Ptr("PT10M"),
		// 					DataUnit: to.Ptr("Count"),
		// 					EvaluationRules: &armcloudhealth.EvaluationRule{
		// 						DegradedRule: &armcloudhealth.ThresholdRule{
		// 							Operator: to.Ptr(armcloudhealth.SignalOperatorLowerThan),
		// 							Threshold: to.Ptr("10"),
		// 						},
		// 						UnhealthyRule: &armcloudhealth.ThresholdRule{
		// 							Operator: to.Ptr(armcloudhealth.SignalOperatorLowerThan),
		// 							Threshold: to.Ptr("1"),
		// 						},
		// 					},
		// 				},
		// 				ID: to.Ptr("/subscriptions/4980D7D5-4E07-47AD-AD34-E76C6BC9F061/resourceGroups/rgopenapi/providers/Microsoft.CloudHealth/healthmodels/myHealthModel/signalDefinitions/sig1"),
		// 				Name: to.Ptr("sig1"),
		// 				Type: to.Ptr("Microsoft.CloudHealth/healthmodels/signalDefinitions"),
		// 				SystemData: &armcloudhealth.SystemData{
		// 					CreatedBy: to.Ptr("cbhzxxlvkmufetjjjwtk"),
		// 					CreatedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-18T14:04:09.327Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("arz"),
		// 					LastModifiedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-18T14:04:09.328Z"); return t}()),
		// 				},
		// 			},
		// 			{
		// 				Properties: &armcloudhealth.ResourceMetricSignalDefinitionProperties{
		// 					MetricNamespace: to.Ptr("microsoft.compute/virtualMachines"),
		// 					MetricName: to.Ptr("memoryusage"),
		// 					AggregationType: to.Ptr(armcloudhealth.MetricAggregationTypeAverage),
		// 					ProvisioningState: to.Ptr(armcloudhealth.HealthModelProvisioningStateSucceeded),
		// 					DisplayName: to.Ptr("memorey usage"),
		// 					SignalKind: to.Ptr(armcloudhealth.SignalKindAzureResourceMetric),
		// 					RefreshInterval: to.Ptr(armcloudhealth.RefreshIntervalPT1M),
		// 					Labels: map[string]*string{
		// 						"key4788": to.Ptr("bar"),
		// 					},
		// 					TimeGrain: to.Ptr("PT10M"),
		// 					DataUnit: to.Ptr("Count"),
		// 					EvaluationRules: &armcloudhealth.EvaluationRule{
		// 						DynamicDetectionRule: &armcloudhealth.DynamicDetectionRule{
		// 							DynamicThresholdModel: to.Ptr(armcloudhealth.DynamicThresholdModelAnomalyDetection),
		// 							ModelSensitivity: to.Ptr[float32](4.5),
		// 							DynamicThresholdDirection: to.Ptr(armcloudhealth.DynamicThresholdDirectionGreaterThan),
		// 						},
		// 					},
		// 				},
		// 				ID: to.Ptr("/subscriptions/4980D7D5-4E07-47AD-AD34-E76C6BC9F061/resourceGroups/rgopenapi/providers/Microsoft.CloudHealth/healthmodels/myHealthModel/signalDefinitions/sig2"),
		// 				Name: to.Ptr("sig2"),
		// 				Type: to.Ptr("Microsoft.CloudHealth/healthmodels/signalDefinitions"),
		// 				SystemData: &armcloudhealth.SystemData{
		// 					CreatedBy: to.Ptr("cbhzxxlvkmufetjjjwtk"),
		// 					CreatedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-18T14:04:09.327Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("arz"),
		// 					LastModifiedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-18T14:04:09.328Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/aqvdhgvi"),
		// 	},
		// }
	}
}
