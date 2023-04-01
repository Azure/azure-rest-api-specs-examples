package armtestbase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/testbase/resource-manager/Microsoft.TestBase/preview/2020-12-16-preview/examples/TestBaseAccountsListBySubscription.json
func ExampleAccountsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtestbase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountsClient().NewListBySubscriptionPager(&armtestbase.AccountsClientListBySubscriptionOptions{GetDeleted: nil})
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
		// page.AccountListResult = armtestbase.AccountListResult{
		// 	Value: []*armtestbase.AccountResource{
		// 		{
		// 			Name: to.Ptr("contoso-testBaseAccount1"),
		// 			Type: to.Ptr("Microsoft.TestBase/testBaseAccounts"),
		// 			ID: to.Ptr("/subscriptions/476f61a4-952c-422a-b4db-568a828f35df/resourceGroups/contoso-rg1/providers/Microsoft.TestBase/testBaseAccounts/contoso-testBaseAccount1"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armtestbase.AccountResourceProperties{
		// 				AccessLevel: to.Ptr("EarlyAccess"),
		// 				ProvisioningState: to.Ptr(armtestbase.ProvisioningStateSucceeded),
		// 				SKU: &armtestbase.AccountSKU{
		// 					Name: to.Ptr("S0"),
		// 					Capabilities: []*armtestbase.AccountSKUCapability{
		// 						{
		// 							Name: to.Ptr("SLA"),
		// 							Value: to.Ptr("99.9%"),
		// 						},
		// 						{
		// 							Name: to.Ptr("maximumPackageCount"),
		// 							Value: to.Ptr("5"),
		// 						},
		// 						{
		// 							Name: to.Ptr("packageMaximumRunSeconds"),
		// 							Value: to.Ptr("3600"),
		// 					}},
		// 					Tier: to.Ptr(armtestbase.TierStandard),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("contoso-testBaseAccount2"),
		// 			Type: to.Ptr("Microsoft.TestBase/testBaseAccounts"),
		// 			ID: to.Ptr("/subscriptions/476f61a4-952c-422a-b4db-568a828f35df/resourceGroups/contoso-rg2/providers/Microsoft.TestBase/testBaseAccounts/contoso-testBaseAccount2"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armtestbase.AccountResourceProperties{
		// 				AccessLevel: to.Ptr("EarlyAccess"),
		// 				ProvisioningState: to.Ptr(armtestbase.ProvisioningStateSucceeded),
		// 				SKU: &armtestbase.AccountSKU{
		// 					Name: to.Ptr("S0"),
		// 					Capabilities: []*armtestbase.AccountSKUCapability{
		// 						{
		// 							Name: to.Ptr("SLA"),
		// 							Value: to.Ptr("99.999%"),
		// 						},
		// 						{
		// 							Name: to.Ptr("maximumPackageCount"),
		// 							Value: to.Ptr("100"),
		// 						},
		// 						{
		// 							Name: to.Ptr("packageMaximumRunSeconds"),
		// 							Value: to.Ptr("10800"),
		// 					}},
		// 					Tier: to.Ptr(armtestbase.TierStandard),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
