package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateKey.json
func ExampleKeysClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewKeysClient("01234567-89ab-4def-0123-456789abcdef", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "ExampleResourceGroup", "ExampleWorkspace", "somekey", armsynapse.Key{
		Properties: &armsynapse.KeyProperties{
			IsActiveCMK: to.Ptr(true),
			KeyVaultURL: to.Ptr("https://vault.azure.net/keys/somesecret"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Key = armsynapse.Key{
	// 	Name: to.Ptr("somekey"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces/keys"),
	// 	ID: to.Ptr("/subscriptions/01234567-89ab-4def-0123-456789abcdef/resourceGroups/ExampleResourceGroup/providers/Microsoft.Synapse/workspaces/ExampleWorkspace/keys/somekey"),
	// 	Properties: &armsynapse.KeyProperties{
	// 		IsActiveCMK: to.Ptr(true),
	// 		KeyVaultURL: to.Ptr("https://vault.azure.net/keys/somesecret"),
	// 	},
	// }
}
