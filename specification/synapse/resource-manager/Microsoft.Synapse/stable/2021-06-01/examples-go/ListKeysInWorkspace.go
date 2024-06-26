package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListKeysInWorkspace.json
func ExampleKeysClient_NewListByWorkspacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewKeysClient().NewListByWorkspacePager("ExampleResourceGroup", "ExampleWorkspace", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.KeyInfoListResult = armsynapse.KeyInfoListResult{
		// 	Value: []*armsynapse.Key{
		// 		{
		// 			Name: to.Ptr("key1"),
		// 			Type: to.Ptr("Microsoft.Synapse/workspaces/keys"),
		// 			ID: to.Ptr("/subscriptions/01234567-89ab-4def-0123-456789abcdef/resourceGroups/ExampleResourceGroup/providers/Microsoft.Synapse/workspaces/ExampleWorkspace/keys/key1"),
		// 			Properties: &armsynapse.KeyProperties{
		// 				IsActiveCMK: to.Ptr(false),
		// 				KeyVaultURL: to.Ptr("https://vault.azure.net/keys/somesecret1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("key2"),
		// 			Type: to.Ptr("Microsoft.Synapse/workspaces/keys"),
		// 			ID: to.Ptr("/subscriptions/01234567-89ab-4def-0123-456789abcdef/resourceGroups/ExampleResourceGroup/providers/Microsoft.Synapse/workspaces/ExampleWorkspace/keys/key2"),
		// 			Properties: &armsynapse.KeyProperties{
		// 				IsActiveCMK: to.Ptr(true),
		// 				KeyVaultURL: to.Ptr("https://vault.azure.net/keys/somesecret2"),
		// 			},
		// 	}},
		// }
	}
}
