package armprivatedns_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns/v2"
)

// Generated from example definition: 2024-06-01/PrivateZoneListInResourceGroup.json
func ExamplePrivateZonesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armprivatedns.NewClientFactory("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateZonesClient().NewListByResourceGroupPager("resourceGroup1", nil)
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
		// page = armprivatedns.PrivateZonesClientListByResourceGroupResponse{
		// 	PrivateZoneListResult: armprivatedns.PrivateZoneListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/subscriptionId/resourceGroups/resourceGroup1/providers/Microsoft.Network/privateDnsZones?api-version=2024-06-01&$skipToken=skipToken"),
		// 		Value: []*armprivatedns.PrivateZone{
		// 			{
		// 				Name: to.Ptr("privatezone1.com"),
		// 				Type: to.Ptr("Microsoft.Network/privateDnsZones"),
		// 				Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroup1/providers/Microsoft.Network/privateDnsZones/privatezone1.com"),
		// 				Location: to.Ptr("global"),
		// 				Properties: &armprivatedns.PrivateZoneProperties{
		// 					MaxNumberOfRecordSets: to.Ptr[int64](5000),
		// 					MaxNumberOfVirtualNetworkLinks: to.Ptr[int64](100),
		// 					MaxNumberOfVirtualNetworkLinksWithRegistration: to.Ptr[int64](50),
		// 					NumberOfRecordSets: to.Ptr[int64](1),
		// 					NumberOfVirtualNetworkLinks: to.Ptr[int64](0),
		// 					NumberOfVirtualNetworkLinksWithRegistration: to.Ptr[int64](0),
		// 					ProvisioningState: to.Ptr(armprivatedns.ProvisioningStateSucceeded),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key1": to.Ptr("value1"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("privatezone2.com"),
		// 				Type: to.Ptr("Microsoft.Network/privateDnsZones"),
		// 				Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroup1/providers/Microsoft.Network/privateDnsZones/privatezone2.com"),
		// 				Location: to.Ptr("global"),
		// 				Properties: &armprivatedns.PrivateZoneProperties{
		// 					MaxNumberOfRecordSets: to.Ptr[int64](5000),
		// 					MaxNumberOfVirtualNetworkLinks: to.Ptr[int64](100),
		// 					MaxNumberOfVirtualNetworkLinksWithRegistration: to.Ptr[int64](50),
		// 					NumberOfRecordSets: to.Ptr[int64](1),
		// 					NumberOfVirtualNetworkLinks: to.Ptr[int64](0),
		// 					NumberOfVirtualNetworkLinksWithRegistration: to.Ptr[int64](0),
		// 					ProvisioningState: to.Ptr(armprivatedns.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
