package armcloudhealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cloudhealth/armcloudhealth"
)

// Generated from example definition: 2025-05-01-preview/SignalDefinitions_Get.json
func ExampleSignalDefinitionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcloudhealth.NewClientFactory("4980D7D5-4E07-47AD-AD34-E76C6BC9F061", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSignalDefinitionsClient().Get(ctx, "rgopenapi", "myHealthModel", "sig1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcloudhealth.SignalDefinitionsClientGetResponse{
	// 	SignalDefinition: &armcloudhealth.SignalDefinition{
	// 		Properties: &armcloudhealth.ResourceMetricSignalDefinitionProperties{
	// 			MetricNamespace: to.Ptr("microsoft.compute/virtualMachines"),
	// 			MetricName: to.Ptr("cpuusage"),
	// 			AggregationType: to.Ptr(armcloudhealth.MetricAggregationTypeNone),
	// 			Dimension: to.Ptr("nodename"),
	// 			DimensionFilter: to.Ptr("node1"),
	// 			ProvisioningState: to.Ptr(armcloudhealth.HealthModelProvisioningStateSucceeded),
	// 			DisplayName: to.Ptr("cpu upsage"),
	// 			SignalKind: to.Ptr(armcloudhealth.SignalKindAzureResourceMetric),
	// 			RefreshInterval: to.Ptr(armcloudhealth.RefreshIntervalPT1M),
	// 			Labels: map[string]*string{
	// 				"key4788": to.Ptr("foo"),
	// 			},
	// 			TimeGrain: to.Ptr("PT1M"),
	// 			DataUnit: to.Ptr("Count"),
	// 			EvaluationRules: &armcloudhealth.EvaluationRule{
	// 				DynamicDetectionRule: &armcloudhealth.DynamicDetectionRule{
	// 					DynamicThresholdModel: to.Ptr(armcloudhealth.DynamicThresholdModelAnomalyDetection),
	// 					ModelSensitivity: to.Ptr[float32](4.5),
	// 					DynamicThresholdDirection: to.Ptr(armcloudhealth.DynamicThresholdDirectionGreaterThan),
	// 				},
	// 			},
	// 		},
	// 		ID: to.Ptr("/subscriptions/4980D7D5-4E07-47AD-AD34-E76C6BC9F061/resourceGroups/rgopenapi/providers/Microsoft.CloudHealth/healthmodels/myHealthModel/signalDefinitions/sig1"),
	// 		Name: to.Ptr("sig1"),
	// 		Type: to.Ptr("Microsoft.CloudHealth/healthmodels/signalDefinitions"),
	// 		SystemData: &armcloudhealth.SystemData{
	// 			CreatedBy: to.Ptr("cbhzxxlvkmufetjjjwtk"),
	// 			CreatedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-18T14:04:09.327Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("arz"),
	// 			LastModifiedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-18T14:04:09.328Z"); return t}()),
	// 		},
	// 	},
	// }
}
