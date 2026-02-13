package armdisconnectedoperations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/disconnectedoperations/armdisconnectedoperations"
)

// Generated from example definition: 2025-06-01-preview/DisconnectedOperations_Update_MaximumSet_Gen.json
func ExampleClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdisconnectedoperations.NewClientFactory("1F6CACA0-5FFA-47AD-94FD-42368F71E49E", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().Update(ctx, "rgdisconnectedoperations", "demo-resource", armdisconnectedoperations.DisconnectedOperationUpdate{
		Tags: map[string]*string{
			"key2": to.Ptr("value2"),
		},
		Properties: &armdisconnectedoperations.DisconnectedOperationUpdateProperties{
			ConnectionIntent:   to.Ptr(armdisconnectedoperations.ConnectionIntentConnected),
			RegistrationStatus: to.Ptr(armdisconnectedoperations.RegistrationStatusRegistered),
			DeviceVersion:      to.Ptr("2.0.0"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdisconnectedoperations.ClientUpdateResponse{
	// 	DisconnectedOperation: &armdisconnectedoperations.DisconnectedOperation{
	// 		Properties: &armdisconnectedoperations.DisconnectedOperationProperties{
	// 			ProvisioningState: to.Ptr(armdisconnectedoperations.ResourceProvisioningStateSucceeded),
	// 			StampID: to.Ptr("663ee8a3-4ea8-48ec-8810-b1f8b86cb270"),
	// 			BillingModel: to.Ptr(armdisconnectedoperations.BillingModelCapacity),
	// 			ConnectionIntent: to.Ptr(armdisconnectedoperations.ConnectionIntentDisconnected),
	// 			ConnectionStatus: to.Ptr(armdisconnectedoperations.ConnectionStatusDisconnected),
	// 			RegistrationStatus: to.Ptr(armdisconnectedoperations.RegistrationStatusRegistered),
	// 			DeviceVersion: to.Ptr("2.0.0"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("value1"),
	// 			"key2": to.Ptr("value2"),
	// 		},
	// 		Location: to.Ptr("eastus"),
	// 		ID: to.Ptr("/subscriptions/1F6CACA0-5FFA-47AD-94FD-42368F71E49E/resourceGroups/rgdisconnectedOperations/providers/Microsoft.Edge/disconnectedOperations/demo-resource"),
	// 		Name: to.Ptr("demo-resource"),
	// 		Type: to.Ptr("Microsoft.Edge/disconnectedOperations"),
	// 		SystemData: &armdisconnectedoperations.SystemData{
	// 			CreatedBy: to.Ptr("kuqdoz"),
	// 			CreatedByType: to.Ptr(armdisconnectedoperations.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-05T20:03:12.507Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("oxaiifvhjvngkllxgultcvbghtyyx"),
	// 			LastModifiedByType: to.Ptr(armdisconnectedoperations.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-05T20:03:12.507Z"); return t}()),
	// 		},
	// 	},
	// }
}
