package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1c63635d66ae38cff18045ab416a6572d3e15f6e/specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/dnsPrivateViews_listByLocation.json
func ExampleDNSPrivateViewsClient_NewListByLocationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDNSPrivateViewsClient().NewListByLocationPager("eastus", nil)
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
		// page.DNSPrivateViewListResult = armoracledatabase.DNSPrivateViewListResult{
		// 	Value: []*armoracledatabase.DNSPrivateView{
		// 		{
		// 			Type: to.Ptr("Oracle.Database/locations/dnsPrivateViews"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Oracle.Database/locations/eastus/dnsPrivateViews/ocid1....aaaaaa"),
		// 			Properties: &armoracledatabase.DNSPrivateViewProperties{
		// 				DisplayName: to.Ptr("example-dns-private-view1"),
		// 				IsProtected: to.Ptr(false),
		// 				LifecycleState: to.Ptr(armoracledatabase.DNSPrivateViewsLifecycleStateActive),
		// 				Ocid: to.Ptr("ocid1....aaaaaa"),
		// 				Self: to.Ptr("https://api.example.com/view1"),
		// 				TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-26T12:34:56.000Z"); return t}()),
		// 				TimeUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-26T14:00:00.000Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Type: to.Ptr("Oracle.Database/locations/dnsPrivateViews"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Oracle.Database/locations/eastus/dnsPrivateViews/ocid1....aaaaab"),
		// 			Properties: &armoracledatabase.DNSPrivateViewProperties{
		// 				DisplayName: to.Ptr("example-dns-private-view2"),
		// 				IsProtected: to.Ptr(true),
		// 				LifecycleState: to.Ptr(armoracledatabase.DNSPrivateViewsLifecycleState("Creating")),
		// 				Ocid: to.Ptr("ocid1....aaaaab"),
		// 				Self: to.Ptr("https://api.example.com/view2"),
		// 				TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-27T09:45:00.000Z"); return t}()),
		// 				TimeUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-27T10:30:00.000Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
