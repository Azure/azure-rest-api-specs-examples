package armloadtesting_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/loadtesting/armloadtesting"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/LoadTests_Update.json
func ExampleLoadTestsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armloadtesting.NewLoadTestsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx, "dummyrg", "myLoadTest", armloadtesting.LoadTestResourcePatchRequestBody{
		Identity: &armloadtesting.ManagedServiceIdentity{
			Type: to.Ptr(armloadtesting.ManagedServiceIdentityTypeSystemAssignedUserAssigned),
			UserAssignedIdentities: map[string]*armloadtesting.UserAssignedIdentity{
				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/dummyrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": {},
			},
		},
		Properties: &armloadtesting.LoadTestResourcePatchRequestBodyProperties{
			Description: to.Ptr("This is new load test resource"),
			Encryption: &armloadtesting.EncryptionProperties{
				Identity: &armloadtesting.EncryptionPropertiesIdentity{
					Type: to.Ptr(armloadtesting.TypeSystemAssigned),
				},
				KeyURL: to.Ptr("https://dummy.vault.azure.net/keys/dummykey1"),
			},
		},
		Tags: map[string]*string{
			"Division": to.Ptr("LT"),
			"Team":     to.Ptr("Dev Exp"),
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
