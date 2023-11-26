package armdatashare_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/Shares_ListByAccount.json
func ExampleSharesClient_NewListByAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatashare.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSharesClient().NewListByAccountPager("SampleResourceGroup", "Account1", &armdatashare.SharesClientListByAccountOptions{SkipToken: nil,
		Filter:  nil,
		Orderby: nil,
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
		// page.ShareList = armdatashare.ShareList{
		// 	Value: []*armdatashare.Share{
		// 		{
		// 			Name: to.Ptr("Share1"),
		// 			Type: to.Ptr("Microsoft.DataShare/accounts/shares"),
		// 			ID: to.Ptr("/subscriptions/433a8dfd-e5d5-4e77-ad86-90acdc75eb1a/resourceGroups/SampleResourceGroup/providers/Microsoft.DataShare/accounts/Account1/shares/Share1"),
		// 			Properties: &armdatashare.ShareProperties{
		// 				Description: to.Ptr("share description"),
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-14T06:15:15.681Z"); return t}()),
		// 				ShareKind: to.Ptr(armdatashare.ShareKindCopyBased),
		// 				Terms: to.Ptr("Confidential"),
		// 				UserEmail: to.Ptr("johnsmith@microsoft.com"),
		// 				UserName: to.Ptr("John Smith"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Share1"),
		// 			Type: to.Ptr("Microsoft.DataShare/accounts/shares"),
		// 			ID: to.Ptr("/subscriptions/433a8dfd-e5d5-4e77-ad86-90acdc75eb1a/resourceGroups/SampleResourceGroup/providers/Microsoft.DataShare/accounts/Account1/shares/Share1"),
		// 			Properties: &armdatashare.ShareProperties{
		// 				Description: to.Ptr("share description"),
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-14T06:15:15.681Z"); return t}()),
		// 				ShareKind: to.Ptr(armdatashare.ShareKindCopyBased),
		// 				Terms: to.Ptr("Confidential"),
		// 				UserEmail: to.Ptr("johnsmith@microsoft.com"),
		// 				UserName: to.Ptr("John Smith"),
		// 			},
		// 	}},
		// }
	}
}
