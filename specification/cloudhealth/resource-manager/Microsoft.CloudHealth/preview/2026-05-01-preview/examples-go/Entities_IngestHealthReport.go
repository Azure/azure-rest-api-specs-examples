package armcloudhealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cloudhealth/armcloudhealth"
)

// Generated from example definition: 2026-05-01-preview/Entities_IngestHealthReport.json
func ExampleEntitiesClient_IngestHealthReport() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcloudhealth.NewClientFactory("abcdef12-3456-7890-abcd-ef1234567890", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewEntitiesClient().IngestHealthReport(ctx, "online-store-rg", "online-store", "orders-api", armcloudhealth.HealthReportRequest{
		SignalName:  to.Ptr("error-rate"),
		HealthState: to.Ptr(armcloudhealth.HealthStateUnhealthy),
		Value:       to.Ptr[float64](6.5),
		EvaluationRules: &armcloudhealth.HealthReportEvaluationRule{
			DegradedRule: &armcloudhealth.ThresholdRuleV2{
				Operator:  to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
				Threshold: to.Ptr[float64](1),
			},
			UnhealthyRule: &armcloudhealth.ThresholdRuleV2{
				Operator:  to.Ptr(armcloudhealth.SignalOperatorGreaterThan),
				Threshold: to.Ptr[float64](5),
			},
		},
		ExpiresInMinutes:  to.Ptr[int32](60),
		AdditionalContext: to.Ptr("Elevated 5xx error rate during the checkout traffic spike."),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
