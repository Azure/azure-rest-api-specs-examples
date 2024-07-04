package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1c63635d66ae38cff18045ab416a6572d3e15f6e/specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/dnsPrivateZones_get.json
func ExampleDNSPrivateZonesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDNSPrivateZonesClient().Get(ctx, "eastus", "example-dns-private-zone", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DNSPrivateZone = armoracledatabase.DNSPrivateZone{
	// 	Type: to.Ptr("Oracle.Database/locations/dnsPrivateZones"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Oracle.Database/locations/eastus/dnsPrivateZones/example-dns-private-zone"),
	// 	Properties: &armoracledatabase.DNSPrivateZoneProperties{
	// 		IsProtected: to.Ptr(false),
	// 		LifecycleState: to.Ptr(armoracledatabase.DNSPrivateZonesLifecycleStateActive),
	// 		Ocid: to.Ptr("your-zone-id"),
	// 		Self: to.Ptr("https://api.example.com/your-dns-private-zone"),
	// 		Serial: to.Ptr[int32](12345),
	// 		TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-26T12:34:56.000Z"); return t}()),
	// 		Version: to.Ptr("1.0.0.0"),
	// 		ViewID: to.Ptr("your-view-id"),
	// 		ZoneType: to.Ptr(armoracledatabase.ZoneTypePrimary),
	// 	},
	// }
}
