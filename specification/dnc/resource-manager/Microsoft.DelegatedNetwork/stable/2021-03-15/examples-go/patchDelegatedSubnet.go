package armdelegatednetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/delegatednetwork/armdelegatednetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dnc/resource-manager/Microsoft.DelegatedNetwork/stable/2021-03-15/examples/patchDelegatedSubnet.json
func ExampleDelegatedSubnetServiceClient_BeginPatchDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdelegatednetwork.NewDelegatedSubnetServiceClient("613192d7-503f-477a-9cfe-4efc3ee2bd60", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginPatchDetails(ctx,
		"TestRG",
		"delegated1",
		armdelegatednetwork.ResourceUpdateParameters{
			Tags: map[string]*string{
				"key": to.Ptr("value"),
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
