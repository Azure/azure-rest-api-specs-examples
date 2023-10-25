package armsignalr_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/signalr/armsignalr"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/signalr/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/SignalRSharedPrivateLinkResources_CreateOrUpdate.json
func ExampleSharedPrivateLinkResourcesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsignalr.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSharedPrivateLinkResourcesClient().BeginCreateOrUpdate(ctx, "upstream", "myResourceGroup", "mySignalRService", armsignalr.SharedPrivateLinkResource{
		Properties: &armsignalr.SharedPrivateLinkResourceProperties{
			GroupID:               to.Ptr("sites"),
			PrivateLinkResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.Web/sites/myWebApp"),
			RequestMessage:        to.Ptr("Please approve"),
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
	// res.SharedPrivateLinkResource = armsignalr.SharedPrivateLinkResource{
	// 	Name: to.Ptr("upstream"),
	// 	Type: to.Ptr("Microsoft.SignalRService/SignalR/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/SignalR/mySignalRService/privateEndpointConnections/upstream"),
	// 	Properties: &armsignalr.SharedPrivateLinkResourceProperties{
	// 		GroupID: to.Ptr("sites"),
	// 		PrivateLinkResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.Web/sites/myWebApp"),
	// 		ProvisioningState: to.Ptr(armsignalr.ProvisioningStateSucceeded),
	// 		RequestMessage: to.Ptr("Please approve"),
	// 		Status: to.Ptr(armsignalr.SharedPrivateLinkResourceStatusApproved),
	// 	},
	// }
}
