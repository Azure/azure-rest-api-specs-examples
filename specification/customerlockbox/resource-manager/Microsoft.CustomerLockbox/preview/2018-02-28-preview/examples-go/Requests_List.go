package armcustomerlockbox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerlockbox/armcustomerlockbox"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/customerlockbox/resource-manager/Microsoft.CustomerLockbox/preview/2018-02-28-preview/examples/Requests_List.json
func ExampleRequestsClient_NewListPager_listLockboxRequestsWithNoFilters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcustomerlockbox.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRequestsClient().NewListPager("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", &armcustomerlockbox.RequestsClientListOptions{Filter: nil})
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
		// page.RequestListResult = armcustomerlockbox.RequestListResult{
		// 	Value: []*armcustomerlockbox.LockboxRequestResponse{
		// 		{
		// 			Name: to.Ptr("LockboxRequest_xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
		// 			Type: to.Ptr("requests"),
		// 			ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/providers/Microsoft.CustomerLockbox/requests/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
		// 			Properties: &armcustomerlockbox.LockboxRequestResponseProperties{
		// 				CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-17T23:15:42.031Z"); return t}()),
		// 				Duration: to.Ptr("8"),
		// 				ExpirationDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-21T23:15:42.031Z"); return t}()),
		// 				Justification: to.Ptr("Microsoft Support Team is requesting access to your resource temporarily for troubleshooting."),
		// 				RequestID: to.Ptr("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
		// 				ResourceIDs: to.Ptr("SUBSCRIPTIONS/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/RESOURCEGROUPS/DEMORG/PROVIDERS/MICROSOFT.COMPUTE/VIRTUALMACHINES/DEMO"),
		// 				ResourceType: to.Ptr("VMSS"),
		// 				Status: to.Ptr(armcustomerlockbox.StatusExpired),
		// 				SubscriptionID: to.Ptr("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
		// 				SupportCaseURL: to.Ptr("https://portal.azure.com/#resource/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/providers/microsoft.support/supporttickets/12345"),
		// 				SupportRequest: to.Ptr("12345"),
		// 				Workitemsource: to.Ptr("Other"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("LockboxRequest_xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
		// 			Type: to.Ptr("requests"),
		// 			ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/providers/Microsoft.CustomerLockbox/requests/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
		// 			Properties: &armcustomerlockbox.LockboxRequestResponseProperties{
		// 				CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-19T03:12:51.267Z"); return t}()),
		// 				Duration: to.Ptr("8"),
		// 				ExpirationDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-23T03:12:51.267Z"); return t}()),
		// 				Justification: to.Ptr("Microsoft Support Team is requesting access to your resource temporarily for troubleshooting."),
		// 				RequestID: to.Ptr("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
		// 				ResourceIDs: to.Ptr("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
		// 				ResourceType: to.Ptr("Subscription"),
		// 				Status: to.Ptr(armcustomerlockbox.StatusExpired),
		// 				SubscriptionID: to.Ptr("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
		// 				SupportCaseURL: to.Ptr("https://portal.azure.com/#resource/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/providers/microsoft.support/supporttickets/120101324000234"),
		// 				SupportRequest: to.Ptr("120101324000234"),
		// 				Workitemsource: to.Ptr("SupportRequest"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("LockboxRequest_xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
		// 			Type: to.Ptr("requests"),
		// 			ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/providers/Microsoft.CustomerLockbox/requests/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
		// 			Properties: &armcustomerlockbox.LockboxRequestResponseProperties{
		// 				CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-24T15:08:17.590Z"); return t}()),
		// 				Duration: to.Ptr("Support request lifetime"),
		// 				ExpirationDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-28T15:08:17.590Z"); return t}()),
		// 				Justification: to.Ptr("Microsoft Support Team is requesting access to your resource temporarily for troubleshooting."),
		// 				RequestID: to.Ptr("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
		// 				ResourceIDs: to.Ptr("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
		// 				ResourceType: to.Ptr("Subscription"),
		// 				Status: to.Ptr(armcustomerlockbox.StatusDenied),
		// 				SubscriptionID: to.Ptr("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
		// 				SupportCaseURL: to.Ptr("https://portal.azure.com/#resource/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/providers/microsoft.support/supporttickets/120101324000234"),
		// 				SupportRequest: to.Ptr("120101324000234"),
		// 				Workitemsource: to.Ptr("SupportRequest"),
		// 			},
		// 	}},
		// }
	}
}
