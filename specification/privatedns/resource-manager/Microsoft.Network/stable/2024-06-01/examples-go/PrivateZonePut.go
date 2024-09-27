package armprivatedns_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4991356eec55985c1af47096c9c2091126a7d0f/specification/privatedns/resource-manager/Microsoft.Network/stable/2024-06-01/examples/PrivateZonePut.json
func ExamplePrivateZonesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armprivatedns.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateZonesClient().BeginCreateOrUpdate(ctx, "resourceGroup1", "privatezone1.com", armprivatedns.PrivateZone{
		Location: to.Ptr("Global"),
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
		},
	}, &armprivatedns.PrivateZonesClientBeginCreateOrUpdateOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
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
	// res.PrivateZone = armprivatedns.PrivateZone{
	// 	Name: to.Ptr("privatezone1.com"),
	// 	Type: to.Ptr("Microsoft.Network/privateDnsZones"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroup1/providers/Microsoft.Network/privateDnsZones/privatezone1.com"),
	// 	Location: to.Ptr("global"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 	},
	// 	Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 	Properties: &armprivatedns.PrivateZoneProperties{
	// 		MaxNumberOfRecordSets: to.Ptr[int64](5000),
	// 		MaxNumberOfVirtualNetworkLinks: to.Ptr[int64](100),
	// 		MaxNumberOfVirtualNetworkLinksWithRegistration: to.Ptr[int64](50),
	// 		NumberOfRecordSets: to.Ptr[int64](1),
	// 		NumberOfVirtualNetworkLinks: to.Ptr[int64](0),
	// 		NumberOfVirtualNetworkLinksWithRegistration: to.Ptr[int64](0),
	// 		ProvisioningState: to.Ptr(armprivatedns.ProvisioningStateSucceeded),
	// 	},
	// }
}
