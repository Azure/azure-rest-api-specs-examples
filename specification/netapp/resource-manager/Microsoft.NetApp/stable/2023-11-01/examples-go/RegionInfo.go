package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-11-01/examples/RegionInfo.json
func ExampleResourceClient_QueryRegionInfo() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewResourceClient().QueryRegionInfo(ctx, "eastus", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RegionInfo = armnetapp.RegionInfo{
	// 	AvailabilityZoneMappings: []*armnetapp.RegionInfoAvailabilityZoneMappingsItem{
	// 		{
	// 			AvailabilityZone: to.Ptr("1"),
	// 			IsAvailable: to.Ptr(true),
	// 	}},
	// 	StorageToNetworkProximity: to.Ptr(armnetapp.RegionStorageToNetworkProximityT2),
	// }
}
