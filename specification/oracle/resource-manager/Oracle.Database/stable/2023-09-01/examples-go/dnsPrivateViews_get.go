package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1c63635d66ae38cff18045ab416a6572d3e15f6e/specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/dnsPrivateViews_get.json
func ExampleDNSPrivateViewsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDNSPrivateViewsClient().Get(ctx, "eastus", "ocid1....aaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DNSPrivateView = armoracledatabase.DNSPrivateView{
	// 	Type: to.Ptr("Oracle.Database/locations/dnsPrivateViews"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Oracle.Database/locations/eastus/dnsPrivateViews/ocid1....aaaaaa"),
	// 	Properties: &armoracledatabase.DNSPrivateViewProperties{
	// 		DisplayName: to.Ptr("example-dns-private-view"),
	// 		IsProtected: to.Ptr(false),
	// 		LifecycleState: to.Ptr(armoracledatabase.DNSPrivateViewsLifecycleStateActive),
	// 		Ocid: to.Ptr("ocid1....aaaaaa"),
	// 		Self: to.Ptr("https://api.example.com/your-dns-private-view"),
	// 		TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-26T12:34:56.000Z"); return t}()),
	// 		TimeUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-26T14:00:00.000Z"); return t}()),
	// 	},
	// }
}
