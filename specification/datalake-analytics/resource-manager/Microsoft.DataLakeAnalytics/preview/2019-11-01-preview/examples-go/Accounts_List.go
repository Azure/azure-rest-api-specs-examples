package armdatalakeanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/preview/2019-11-01-preview/examples/Accounts_List.json
func ExampleAccountsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatalakeanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountsClient().NewListPager(&armdatalakeanalytics.AccountsClientListOptions{Filter: to.Ptr("test_filter"),
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
		// page.AccountListResult = armdatalakeanalytics.AccountListResult{
		// 	Value: []*armdatalakeanalytics.AccountBasic{
		// 		{
		// 			Name: to.Ptr("test_account"),
		// 			Type: to.Ptr("test_type"),
		// 			ID: to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345"),
		// 			Properties: &armdatalakeanalytics.AccountPropertiesBasic{
		// 				AccountID: to.Ptr("94f4bf5d-78a9-4c31-8aa7-b34d07bad898"),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-14T20:21:56.681Z"); return t}()),
		// 				Endpoint: to.Ptr("test_endpoint"),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-14T20:21:56.681Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armdatalakeanalytics.DataLakeAnalyticsAccountStatusSucceeded),
		// 				State: to.Ptr(armdatalakeanalytics.DataLakeAnalyticsAccountStateActive),
		// 			},
		// 	}},
		// }
	}
}
