package armcloudhealth_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cloudhealth/armcloudhealth"
)

// Generated from example definition: 2026-05-01-preview/Entities_GetSignalHistory.json
func ExampleEntitiesClient_GetSignalHistory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcloudhealth.NewClientFactory("abcdef12-3456-7890-abcd-ef1234567890", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEntitiesClient().GetSignalHistory(ctx, "online-store-rg", "online-store", "web-frontend", armcloudhealth.SignalHistoryRequest{
		SignalName: to.Ptr("http-5xx"),
		StartAt:    to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-03T09:30:00Z"); return t }()),
		EndAt:      to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T09:30:00Z"); return t }()),
		Top:        to.Ptr[int32](7),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcloudhealth.EntitiesClientGetSignalHistoryResponse{
	// 	SignalHistoryResponse: armcloudhealth.SignalHistoryResponse{
	// 		EntityName: to.Ptr("web-frontend"),
	// 		SignalName: to.Ptr("http-5xx"),
	// 		History: []*armcloudhealth.SignalHistoryDataPoint{
	// 			{
	// 				OccurredAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-03T10:00:00Z"); return t}()),
	// 				Value: to.Ptr[float64](0),
	// 				HealthState: to.Ptr(armcloudhealth.HealthStateHealthy),
	// 			},
	// 			{
	// 				OccurredAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-03T14:00:00Z"); return t}()),
	// 				Value: to.Ptr[float64](3),
	// 				HealthState: to.Ptr(armcloudhealth.HealthStateHealthy),
	// 			},
	// 			{
	// 				OccurredAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-03T18:00:00Z"); return t}()),
	// 				Value: to.Ptr[float64](12),
	// 				HealthState: to.Ptr(armcloudhealth.HealthStateDegraded),
	// 			},
	// 			{
	// 				OccurredAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-03T22:00:00Z"); return t}()),
	// 				Value: to.Ptr[float64](27),
	// 				HealthState: to.Ptr(armcloudhealth.HealthStateUnhealthy),
	// 			},
	// 			{
	// 				OccurredAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T02:00:00Z"); return t}()),
	// 				Value: to.Ptr[float64](8),
	// 				HealthState: to.Ptr(armcloudhealth.HealthStateDegraded),
	// 			},
	// 			{
	// 				OccurredAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T06:00:00Z"); return t}()),
	// 				Value: to.Ptr[float64](1),
	// 				HealthState: to.Ptr(armcloudhealth.HealthStateHealthy),
	// 			},
	// 			{
	// 				OccurredAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T09:00:00Z"); return t}()),
	// 				Value: to.Ptr[float64](0),
	// 				HealthState: to.Ptr(armcloudhealth.HealthStateHealthy),
	// 			},
	// 		},
	// 		NextMarker: to.Ptr("eyJsYXN0VGltZXN0YW1wIjoiMjAyNi0wNS0wNFQwOTowMDowMFoifQ=="),
	// 	},
	// }
}
