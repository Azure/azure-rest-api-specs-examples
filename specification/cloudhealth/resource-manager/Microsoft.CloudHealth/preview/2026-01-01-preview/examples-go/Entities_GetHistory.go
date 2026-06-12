package armcloudhealth_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cloudhealth/armcloudhealth"
)

// Generated from example definition: 2026-01-01-preview/Entities_GetHistory.json
func ExampleEntitiesClient_GetHistory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcloudhealth.NewClientFactory("4980D7D5-4E07-47AD-AD34-E76C6BC9F061", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEntitiesClient().GetHistory(ctx, "rgopenapi", "myHealthModel", "entity1", armcloudhealth.EntityHistoryRequest{
		StartAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-12-11T10:00:00Z"); return t }()),
		EndAt:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-12-12T10:00:00Z"); return t }()),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcloudhealth.EntitiesClientGetHistoryResponse{
	// 	EntityHistoryResponse: armcloudhealth.EntityHistoryResponse{
	// 		EntityName: to.Ptr("entity1"),
	// 		History: []*armcloudhealth.HealthStateTransition{
	// 			{
	// 				PreviousState: to.Ptr(armcloudhealth.HealthStateHealthy),
	// 				NewState: to.Ptr(armcloudhealth.HealthStateDegraded),
	// 				OccurredAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-12-11T14:30:00Z"); return t}()),
	// 				Reason: to.Ptr("SignalTransition"),
	// 			},
	// 			{
	// 				PreviousState: to.Ptr(armcloudhealth.HealthStateDegraded),
	// 				NewState: to.Ptr(armcloudhealth.HealthStateUnhealthy),
	// 				OccurredAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-12-11T18:45:00Z"); return t}()),
	// 			},
	// 			{
	// 				PreviousState: to.Ptr(armcloudhealth.HealthStateUnhealthy),
	// 				NewState: to.Ptr(armcloudhealth.HealthStateDegraded),
	// 				OccurredAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-12-11T22:15:00Z"); return t}()),
	// 			},
	// 			{
	// 				PreviousState: to.Ptr(armcloudhealth.HealthStateDegraded),
	// 				NewState: to.Ptr(armcloudhealth.HealthStateHealthy),
	// 				OccurredAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-12-12T02:30:00Z"); return t}()),
	// 			},
	// 		},
	// 	},
	// }
}
