package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/44319b51c6f952fdc9543d3dc4fdd9959350d102/specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-03-01/examples/RegionInfos_Get.json
func ExampleResourceRegionInfosClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewResourceRegionInfosClient().Get(ctx, "eastus", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RegionInfoResource = armnetapp.RegionInfoResource{
	// 	Name: to.Ptr("eastus/default"),
	// 	Type: to.Ptr("Microsoft.NetApp/locations/regionInfos"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.NetApp/locations/eastus/regionInfos/default"),
	// 	Properties: &armnetapp.RegionInfo{
	// 		AvailabilityZoneMappings: []*armnetapp.RegionInfoAvailabilityZoneMappingsItem{
	// 			{
	// 				AvailabilityZone: to.Ptr("1"),
	// 				IsAvailable: to.Ptr(true),
	// 			},
	// 			{
	// 				AvailabilityZone: to.Ptr("2"),
	// 				IsAvailable: to.Ptr(true),
	// 			},
	// 			{
	// 				AvailabilityZone: to.Ptr("3"),
	// 				IsAvailable: to.Ptr(true),
	// 		}},
	// 		StorageToNetworkProximity: to.Ptr(armnetapp.RegionStorageToNetworkProximityT2),
	// 	},
	// }
}
