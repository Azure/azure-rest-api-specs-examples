package armartifactsigning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/artifactsigning/armartifactsigning"
)

// Generated from example definition: 2025-10-13/CodeSigningAccounts_Update.json
func ExampleCodeSigningAccountsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armartifactsigning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCodeSigningAccountsClient().BeginUpdate(ctx, "MyResourceGroup", "MyAccount", armartifactsigning.CodeSigningAccountPatch{
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
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
	// res = armartifactsigning.CodeSigningAccountsClientUpdateResponse{
	// 	CodeSigningAccount: &armartifactsigning.CodeSigningAccount{
	// 		Name: to.Ptr("MyAccount"),
	// 		Type: to.Ptr("Microsoft.CodeSigning/codeSigningAccounts"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/MyResourceGroup/providers/Microsoft.CodeSigning/codeSigningAccounts/MyAccount"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armartifactsigning.CodeSigningAccountProperties{
	// 			ProvisioningState: to.Ptr(armartifactsigning.ProvisioningStateSucceeded),
	// 			SKU: &armartifactsigning.AccountSKU{
	// 				Name: to.Ptr(armartifactsigning.SKUNameBasic),
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("value1"),
	// 		},
	// 	},
	// }
}
