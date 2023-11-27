package armdatalakestore_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-store/armdatalakestore"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/Accounts_ListByResourceGroup.json
func ExampleAccountsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatalakestore.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountsClient().NewListByResourceGroupPager("contosorg", &armdatalakestore.AccountsClientListByResourceGroupOptions{Filter: to.Ptr("test_filter"),
		Top:     to.Ptr[int32](1),
		Skip:    to.Ptr[int32](1),
		Select:  to.Ptr("test_select"),
		Orderby: to.Ptr("test_orderby"),
		Count:   to.Ptr(false),
	})
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
		// page.AccountListResult = armdatalakestore.AccountListResult{
		// 	Value: []*armdatalakestore.AccountBasic{
		// 		{
		// 			Name: to.Ptr("contosoadla"),
		// 			Type: to.Ptr("test_type"),
		// 			ID: to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345"),
		// 			Location: to.Ptr("eastus2"),
		// 			Tags: map[string]*string{
		// 				"test_key": to.Ptr("test_value"),
		// 			},
		// 			Properties: &armdatalakestore.AccountPropertiesBasic{
		// 				AccountID: to.Ptr("94f4bf5d-78a9-4c31-8aa7-b34d07bad898"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-14T20:21:56.681Z"); return t}()),
		// 				Endpoint: to.Ptr("testadlfs17607.azuredatalakestore.net"),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-14T20:21:56.681Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armdatalakestore.DataLakeStoreAccountStatusSucceeded),
		// 				State: to.Ptr(armdatalakestore.DataLakeStoreAccountStateActive),
		// 			},
		// 	}},
		// }
	}
}
