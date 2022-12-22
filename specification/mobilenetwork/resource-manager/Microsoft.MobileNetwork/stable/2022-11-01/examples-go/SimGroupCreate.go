package armmobilenetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2022-11-01/examples/SimGroupCreate.json
func ExampleSimGroupsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewSimGroupsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "testSimGroup", armmobilenetwork.SimGroup{
		Location: to.Ptr("eastus"),
		Identity: &armmobilenetwork.ManagedServiceIdentity{
			Type: to.Ptr(armmobilenetwork.ManagedServiceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armmobilenetwork.UserAssignedIdentity{
				"/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testUserAssignedManagedIdentity": {},
			},
		},
		Properties: &armmobilenetwork.SimGroupPropertiesFormat{
			EncryptionKey: &armmobilenetwork.KeyVaultKey{
				KeyURL: to.Ptr("https://contosovault.vault.azure.net/keys/azureKey"),
			},
			MobileNetwork: &armmobilenetwork.ResourceID{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork"),
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
	// TODO: use response item
	_ = res
}
