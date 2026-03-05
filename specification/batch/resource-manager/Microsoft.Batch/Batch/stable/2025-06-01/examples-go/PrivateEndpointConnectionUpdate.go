package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v4"
)

// Generated from example definition: 2025-06-01/PrivateEndpointConnectionUpdate.json
func ExamplePrivateEndpointConnectionClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("12345678-1234-1234-1234-123456789012", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionClient().BeginUpdate(ctx, "default-azurebatch-japaneast", "sampleacct", "testprivateEndpointConnection5.24d6b4b5-e65c-4330-bbe9-3a290d62f8e0", armbatch.PrivateEndpointConnection{
		Properties: &armbatch.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armbatch.PrivateLinkServiceConnectionState{
				Description: to.Ptr("Approved by xyz.abc@company.com"),
				Status:      to.Ptr(armbatch.PrivateLinkServiceConnectionStatusApproved),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbatch.PrivateEndpointConnectionClientUpdateResponse{
	// 	PrivateEndpointConnection: &armbatch.PrivateEndpointConnection{
	// 		Name: to.Ptr("testprivateEndpointConnection5.24d6b4b5-e65c-4330-bbe9-3a290d62f8e0"),
	// 		Type: to.Ptr("Microsoft.Batch/batchAccounts/privateEndpointConnections"),
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/privateEndpointConnections/testprivateEndpointConnection5.24d6b4b5-e65c-4330-bbe9-3a290d62f8e0"),
	// 		Properties: &armbatch.PrivateEndpointConnectionProperties{
	// 			GroupIDs: []*string{
	// 				to.Ptr("batchAccount"),
	// 			},
	// 			PrivateEndpoint: &armbatch.PrivateEndpoint{
	// 				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Network/privateEndpoints/testprivateEndpoint"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armbatch.PrivateLinkServiceConnectionState{
	// 				Description: to.Ptr("Approved by xyz.abc@company.com"),
	// 				Status: to.Ptr(armbatch.PrivateLinkServiceConnectionStatusApproved),
	// 			},
	// 			ProvisioningState: to.Ptr(armbatch.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
