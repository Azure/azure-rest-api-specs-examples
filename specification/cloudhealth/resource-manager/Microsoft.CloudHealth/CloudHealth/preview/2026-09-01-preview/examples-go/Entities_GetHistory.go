package armcloudhealth_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cloudhealth/armcloudhealth"
)

// Generated from example definition: 2026-09-01-preview/Entities_GetHistory.json
func ExampleEntitiesClient_GetHistory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcloudhealth.NewClientFactory("abcdef12-3456-7890-abcd-ef1234567890", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEntitiesClient().GetHistory(ctx, "online-store-rg", "online-store", "web-frontend", armcloudhealth.EntityHistoryRequest{
		StartAt: to.Ptr(time.Date(2026, time.May, 3, 9, 30, 0, 0, time.UTC)),
		EndAt:   to.Ptr(time.Date(2026, time.May, 4, 9, 30, 0, 0, time.UTC)),
		Top:     to.Ptr[int32](100),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcloudhealth.EntitiesClientGetHistoryResponse{
	// 	EntityHistoryResponse: armcloudhealth.EntityHistoryResponse{
	// 		EntityName: to.Ptr("web-frontend"),
	// 		History: []*armcloudhealth.HealthStateTransition{
	// 			{
	// 				PreviousState: to.Ptr(armcloudhealth.HealthStateHealthy),
	// 				NewState: to.Ptr(armcloudhealth.HealthStateDegraded),
	// 				OccurredAt: to.Ptr(time.Date(2026, time.May, 3, 14, 30, 0, 0, time.UTC)),
	// 				Reason: to.Ptr("SignalTransition"),
	// 			},
	// 			{
	// 				PreviousState: to.Ptr(armcloudhealth.HealthStateDegraded),
	// 				NewState: to.Ptr(armcloudhealth.HealthStateUnhealthy),
	// 				OccurredAt: to.Ptr(time.Date(2026, time.May, 3, 18, 45, 0, 0, time.UTC)),
	// 				Reason: to.Ptr("ChildEntityTransition"),
	// 			},
	// 			{
	// 				PreviousState: to.Ptr(armcloudhealth.HealthStateUnhealthy),
	// 				NewState: to.Ptr(armcloudhealth.HealthStateDegraded),
	// 				OccurredAt: to.Ptr(time.Date(2026, time.May, 3, 22, 15, 0, 0, time.UTC)),
	// 				Reason: to.Ptr("ChildEntityTransition"),
	// 			},
	// 			{
	// 				PreviousState: to.Ptr(armcloudhealth.HealthStateDegraded),
	// 				NewState: to.Ptr(armcloudhealth.HealthStateHealthy),
	// 				OccurredAt: to.Ptr(time.Date(2026, time.May, 4, 2, 30, 0, 0, time.UTC)),
	// 				Reason: to.Ptr("SignalTransition"),
	// 			},
	// 		},
	// 		NextMarker: to.Ptr("eyJsYXN0VGltZXN0YW1wIjoiMjAyNi0wNS0wNFQwMjozMDowMFoifQ=="),
	// 	},
	// }
}
