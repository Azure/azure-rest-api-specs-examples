package armmixedreality_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mixedreality/armmixedreality"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/remote-rendering/GetByResourceGroup.json
func ExampleRemoteRenderingAccountsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmixedreality.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRemoteRenderingAccountsClient().NewListByResourceGroupPager("MyResourceGroup", nil)
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
		// page.RemoteRenderingAccountPage = armmixedreality.RemoteRenderingAccountPage{
		// 	Value: []*armmixedreality.RemoteRenderingAccount{
		// 		{
		// 			Name: to.Ptr("alpha"),
		// 			Type: to.Ptr("Microsoft.MixedReality/remoteRenderingAccounts"),
		// 			ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/MyResourceGroup/providers/Microsoft.MixedReality/remoteRenderingAccounts/alpha"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Identity: &armmixedreality.Identity{
		// 				Type: to.Ptr("SystemAssigned"),
		// 			},
		// 			Properties: &armmixedreality.AccountProperties{
		// 				AccountDomain: to.Ptr("mixedreality.azure.com"),
		// 				AccountID: to.Ptr("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("omega"),
		// 			Type: to.Ptr("Microsoft.MixedReality/remoteRenderingAccounts"),
		// 			ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/MyResourceGroup/providers/Microsoft.MixedReality/remoteRenderingAccounts/omega"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Identity: &armmixedreality.Identity{
		// 				Type: to.Ptr("SystemAssigned"),
		// 			},
		// 			Properties: &armmixedreality.AccountProperties{
		// 				AccountDomain: to.Ptr("mixedreality.azure.com"),
		// 				AccountID: to.Ptr("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
		// 			},
		// 	}},
		// }
	}
}
