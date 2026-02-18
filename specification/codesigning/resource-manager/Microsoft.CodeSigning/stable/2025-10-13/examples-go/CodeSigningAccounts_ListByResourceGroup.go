package armartifactsigning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/artifactsigning/armartifactsigning"
)

// Generated from example definition: 2025-10-13/CodeSigningAccounts_ListByResourceGroup.json
func ExampleCodeSigningAccountsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armartifactsigning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
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
		// page = armartifactsigning.CodeSigningAccountsClientListByResourceGroupResponse{
		// 	CodeSigningAccountListResult: armartifactsigning.CodeSigningAccountListResult{
		// 		Value: []*armartifactsigning.CodeSigningAccount{
		// 			{
		// 				Name: to.Ptr("alpha"),
		// 				Type: to.Ptr("Microsoft.CodeSigning/codeSigningAccounts"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/MyResourceGroup/providers/Microsoft.CodeSigning/codeSigningAccounts/MyAccount"),
		// 				Location: to.Ptr("westcentralus"),
		// 				Properties: &armartifactsigning.CodeSigningAccountProperties{
		// 					ProvisioningState: to.Ptr(armartifactsigning.ProvisioningStateSucceeded),
		// 					SKU: &armartifactsigning.AccountSKU{
		// 						Name: to.Ptr(armartifactsigning.SKUNameBasic),
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
