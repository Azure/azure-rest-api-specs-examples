package armcloudhealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cloudhealth/armcloudhealth"
)

// Generated from example definition: 2026-01-01-preview/Entities_IngestHealthReport.json
func ExampleEntitiesClient_IngestHealthReport() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcloudhealth.NewClientFactory("4980D7D5-4E07-47AD-AD34-E76C6BC9F061", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewEntitiesClient().IngestHealthReport(ctx, "rgopenapi", "myHealthModel", "entity1", armcloudhealth.HealthReportRequest{
		SignalName:  to.Ptr("uniqueSignalName1"),
		Value:       to.Ptr[float64](85.5),
		HealthState: to.Ptr(armcloudhealth.HealthStateDegraded),
		EvaluationRules: &armcloudhealth.HealthReportEvaluationRule{
			DegradedRule: &armcloudhealth.ThresholdRuleV2{
				Operator:  to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
				Threshold: to.Ptr[float64](70),
			},
			UnhealthyRule: &armcloudhealth.ThresholdRuleV2{
				Operator:  to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
				Threshold: to.Ptr[float64](90),
			},
		},
		ExpiresInMinutes:  to.Ptr[int32](60),
		AdditionalContext: to.Ptr("CPU usage elevated due to batch processing job"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
