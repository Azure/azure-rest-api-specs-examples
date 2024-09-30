package armtrustedsigning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trustedsigning/armtrustedsigning"
)

// Generated from example definition: 2024-02-05-preview/CodeSigningAccounts_Update.json
func ExampleCodeSigningAccountsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtrustedsigning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCodeSigningAccountsClient().BeginUpdate(ctx, "MyResourceGroup", "MyAccount", armtrustedsigning.CodeSigningAccountPatch{
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
	// res = armtrustedsigning.CodeSigningAccountsClientUpdateResponse{
	// 	CodeSigningAccount: &armtrustedsigning.CodeSigningAccount{
	// 		Name: to.Ptr("MyAccount"),
	// 		Type: to.Ptr("Microsoft.CodeSigning/codeSigningAccounts"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/MyResourceGroup/providers/Microsoft.CodeSigning/codeSigningAccounts/MyAccount"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armtrustedsigning.CodeSigningAccountProperties{
	// 			ProvisioningState: to.Ptr(armtrustedsigning.ProvisioningStateSucceeded),
	// 			SKU: &armtrustedsigning.AccountSKU{
	// 				Name: to.Ptr(armtrustedsigning.SKUNameBasic),
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("value1"),
	// 		},
	// 	},
	// }
}
