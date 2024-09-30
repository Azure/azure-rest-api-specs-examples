package armtrustedsigning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trustedsigning/armtrustedsigning"
)

// Generated from example definition: 2024-02-05-preview/CodeSigningAccounts_ListByResourceGroup.json
func ExampleCodeSigningAccountsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtrustedsigning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCodeSigningAccountsClient().NewListByResourceGroupPager("MyResourceGroup", nil)
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
		// page = armtrustedsigning.CodeSigningAccountsClientListByResourceGroupResponse{
		// 	CodeSigningAccountListResult: armtrustedsigning.CodeSigningAccountListResult{
		// 		Value: []*armtrustedsigning.CodeSigningAccount{
		// 			{
		// 				Name: to.Ptr("alpha"),
		// 				Type: to.Ptr("Microsoft.CodeSigning/codeSigningAccounts"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/MyResourceGroup/providers/Microsoft.CodeSigning/codeSigningAccounts/MyAccount"),
		// 				Location: to.Ptr("westcentralus"),
		// 				Properties: &armtrustedsigning.CodeSigningAccountProperties{
		// 					ProvisioningState: to.Ptr(armtrustedsigning.ProvisioningStateSucceeded),
		// 					SKU: &armtrustedsigning.AccountSKU{
		// 						Name: to.Ptr(armtrustedsigning.SKUNameBasic),
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 					"key1": to.Ptr("value1"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
