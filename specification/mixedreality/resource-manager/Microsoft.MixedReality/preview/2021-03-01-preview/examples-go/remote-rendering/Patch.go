package armmixedreality_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mixedreality/armmixedreality"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/remote-rendering/Patch.json
func ExampleRemoteRenderingAccountsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmixedreality.NewRemoteRenderingAccountsClient("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"MyResourceGroup",
		"MyAccount",
		armmixedreality.RemoteRenderingAccount{
			Location: to.Ptr("eastus2euap"),
			Tags: map[string]*string{
				"hero":    to.Ptr("romeo"),
				"heroine": to.Ptr("juliet"),
			},
			Identity: &armmixedreality.Identity{
				Type: to.Ptr("SystemAssigned"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
