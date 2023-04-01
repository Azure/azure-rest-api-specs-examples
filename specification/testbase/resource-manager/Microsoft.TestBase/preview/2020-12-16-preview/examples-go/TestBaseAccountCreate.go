package armtestbase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/testbase/resource-manager/Microsoft.TestBase/preview/2020-12-16-preview/examples/TestBaseAccountCreate.json
func ExampleAccountsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtestbase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccountsClient().BeginCreate(ctx, "contoso-rg1", "contoso-testBaseAccount1", armtestbase.AccountResource{
		Location: to.Ptr("westus"),
		Properties: &armtestbase.AccountResourceProperties{
			SKU: &armtestbase.AccountSKU{
				Name: to.Ptr("S0"),
				Tier: to.Ptr(armtestbase.TierStandard),
			},
		},
	}, &armtestbase.AccountsClientBeginCreateOptions{Restore: nil})
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
	// res.AccountResource = armtestbase.AccountResource{
	// 	Name: to.Ptr("contoso-testBaseAccount1"),
	// 	Type: to.Ptr("Microsoft.TestBase/testBaseAccounts"),
	// 	ID: to.Ptr("/subscriptions/476f61a4-952c-422a-b4db-568a828f35df/resourceGroups/contoso-rg1/providers/Microsoft.TestBase/testBaseAccounts/contoso-testBaseAccount1"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armtestbase.AccountResourceProperties{
	// 		AccessLevel: to.Ptr("StandardAccess"),
	// 		ProvisioningState: to.Ptr(armtestbase.ProvisioningStateSucceeded),
	// 		SKU: &armtestbase.AccountSKU{
	// 			Name: to.Ptr("S0"),
	// 			Capabilities: []*armtestbase.AccountSKUCapability{
	// 				{
	// 					Name: to.Ptr("SLA"),
	// 					Value: to.Ptr("99.9%"),
	// 				},
	// 				{
	// 					Name: to.Ptr("maximumPackageCount"),
	// 					Value: to.Ptr("5"),
	// 				},
	// 				{
	// 					Name: to.Ptr("packageMaximumRunSeconds"),
	// 					Value: to.Ptr("3600"),
	// 			}},
	// 			Tier: to.Ptr(armtestbase.TierStandard),
	// 		},
	// 	},
	// }
}
