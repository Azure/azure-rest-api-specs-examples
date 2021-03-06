package armquantum_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quantum/armquantum"
)

// x-ms-original-file: specification/quantum/resource-manager/Microsoft.Quantum/preview/2019-11-04-preview/examples/quantumWorkspacesPut.json
func ExampleWorkspacesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armquantum.NewWorkspacesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		armquantum.Workspace{
			Location: to.StringPtr("<location>"),
			Properties: &armquantum.WorkspaceResourceProperties{
				Providers: []*armquantum.Provider{
					{
						ProviderID:  to.StringPtr("<provider-id>"),
						ProviderSKU: to.StringPtr("<provider-sku>"),
					},
					{
						ProviderID:  to.StringPtr("<provider-id>"),
						ProviderSKU: to.StringPtr("<provider-sku>"),
					},
					{
						ProviderID:  to.StringPtr("<provider-id>"),
						ProviderSKU: to.StringPtr("<provider-sku>"),
					}},
				StorageAccount: to.StringPtr("<storage-account>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WorkspacesClientCreateOrUpdateResult)
}
