package armengagementfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/engagementfabric/armengagementfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/engagementfabric/resource-manager/Microsoft.EngagementFabric/preview/2018-09-01/examples/AccountsListKeysExample.json
func ExampleAccountsClient_NewListKeysPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armengagementfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountsClient().NewListKeysPager("ExampleRg", "ExampleAccount", nil)
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
		// page.KeyDescriptionList = armengagementfabric.KeyDescriptionList{
		// 	Value: []*armengagementfabric.KeyDescription{
		// 		{
		// 			Name: to.Ptr("Full"),
		// 			Rank: to.Ptr(armengagementfabric.KeyRankPrimaryKey),
		// 			Value: to.Ptr("<ExampleFullPrimaryKeyValue>"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Full"),
		// 			Rank: to.Ptr(armengagementfabric.KeyRankSecondaryKey),
		// 			Value: to.Ptr("<ExampleFullSecondaryKeyValue>"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Device"),
		// 			Rank: to.Ptr(armengagementfabric.KeyRankPrimaryKey),
		// 			Value: to.Ptr("<ExampleDevicePrimaryKeyValue>"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Device"),
		// 			Rank: to.Ptr(armengagementfabric.KeyRankSecondaryKey),
		// 			Value: to.Ptr("<ExampleDeviceSecondaryKeyValue>"),
		// 	}},
		// }
	}
}
