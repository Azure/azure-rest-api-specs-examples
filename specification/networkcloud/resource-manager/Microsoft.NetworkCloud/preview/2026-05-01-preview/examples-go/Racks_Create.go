package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: 2026-05-01-preview/Racks_Create.json
func ExampleRacksClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("123e4567-e89b-12d3-a456-426655440000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRacksClient().BeginCreateOrUpdate(ctx, "resourceGroupName", "rackName", armnetworkcloud.Rack{
		ExtendedLocation: &armnetworkcloud.ExtendedLocation{
			Name: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
			Type: to.Ptr("CustomLocation"),
		},
		Location: to.Ptr("location"),
		Properties: &armnetworkcloud.RackProperties{
			AvailabilityZone: to.Ptr("1"),
			RackLocation:     to.Ptr("Rack 28"),
			RackSerialNumber: to.Ptr("RACK_SERIAL_NUMBER"),
			RackSKUID:        to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/rackSkus/rackSkuName"),
		},
		Tags: map[string]*string{
			"key1": to.Ptr("myvalue1"),
			"key2": to.Ptr("myvalue2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetworkcloud.RacksClientCreateOrUpdateResponse{
	// 	Rack: armnetworkcloud.Rack{
	// 		ExtendedLocation: &armnetworkcloud.ExtendedLocation{
	// 			Name: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
	// 			Type: to.Ptr("CustomLocation"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/racks/rackName"),
	// 		Location: to.Ptr("location"),
	// 		Name: to.Ptr("rackName"),
	// 		Properties: &armnetworkcloud.RackProperties{
	// 			AvailabilityZone: to.Ptr("1"),
	// 			ClusterID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/clusters/clusterName"),
	// 			DetailedStatus: to.Ptr(armnetworkcloud.RackDetailedStatusAvailable),
	// 			DetailedStatusMessage: to.Ptr("Rack is operational"),
	// 			ProvisioningState: to.Ptr(armnetworkcloud.RackProvisioningStateSucceeded),
	// 			RackLocation: to.Ptr("Rack 28"),
	// 			RackSerialNumber: to.Ptr("RACK_SERIAL_NUMBER"),
	// 			RackSKUID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/rackSkus/rackSkuName"),
	// 		},
	// 		SystemData: &armnetworkcloud.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
	// 			CreatedBy: to.Ptr("identityA"),
	// 			CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("identityB"),
	// 			LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("myvalue1"),
	// 			"key2": to.Ptr("myvalue2"),
	// 		},
	// 		Type: to.Ptr("Microsoft.NetworkCloud/racks"),
	// 	},
	// }
}
