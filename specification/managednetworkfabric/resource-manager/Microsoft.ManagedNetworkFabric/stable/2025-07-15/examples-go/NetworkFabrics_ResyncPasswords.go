package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/NetworkFabrics_ResyncPasswords.json
func ExampleNetworkFabricsClient_BeginResyncPasswords_successfulPasswordResyncForTheTerminalServerAndAllNetworkDevices() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("1234ABCD-0A1B-1234-5678-123456ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkFabricsClient().BeginResyncPasswords(ctx, "example-rg", "example-fabric", nil)
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
	// res = armmanagednetworkfabric.NetworkFabricsClientResyncPasswordsResponse{
	// 	NetworkFabricResyncPasswordsResponse: &armmanagednetworkfabric.NetworkFabricResyncPasswordsResponse{
	// 		ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/providers/Microsoft.ManagedNetworkFabric/locations/UKSOUTH/operationStatuses/b7c3d8f2-56e4-4a9b-8c7d-123456789ABC*1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF"),
	// 		ResourceID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric"),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-09-25T08:06:00.912Z"); return t}()),
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-09-25T09:10:00.912Z"); return t}()),
	// 		Status: to.Ptr("Succeeded"),
	// 	},
	// }
}
