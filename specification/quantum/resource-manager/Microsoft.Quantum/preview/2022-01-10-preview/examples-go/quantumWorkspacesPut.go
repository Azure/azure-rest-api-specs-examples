package armquantum_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quantum/armquantum"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/quantum/resource-manager/Microsoft.Quantum/preview/2022-01-10-preview/examples/quantumWorkspacesPut.json
func ExampleWorkspacesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armquantum.NewWorkspacesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"quantumResourcegroup",
		"quantumworkspace1",
		armquantum.Workspace{
			Location: to.Ptr("West US"),
			Properties: &armquantum.WorkspaceResourceProperties{
				Providers: []*armquantum.Provider{
					{
						ProviderID:  to.Ptr("Honeywell"),
						ProviderSKU: to.Ptr("Basic"),
					},
					{
						ProviderID:  to.Ptr("IonQ"),
						ProviderSKU: to.Ptr("Basic"),
					},
					{
						ProviderID:  to.Ptr("OneQBit"),
						ProviderSKU: to.Ptr("Basic"),
					}},
				StorageAccount: to.Ptr("/subscriptions/1C4B2828-7D49-494F-933D-061373BE28C2/resourceGroups/quantumResourcegroup/providers/Microsoft.Storage/storageAccounts/testStorageAccount"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
