package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/Volumes_Create.json
func ExampleVolumesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumesClient().BeginCreateOrUpdate(ctx, "resourceGroupName", "volumeName", armnetworkcloud.Volume{
		Location: to.Ptr("location"),
		Tags: map[string]*string{
			"key1": to.Ptr("myvalue1"),
			"key2": to.Ptr("myvalue2"),
		},
		ExtendedLocation: &armnetworkcloud.ExtendedLocation{
			Name: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
			Type: to.Ptr("CustomLocation"),
		},
		Properties: &armnetworkcloud.VolumeProperties{
			SizeMiB: to.Ptr[int64](10000),
		},
	}, nil)
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
	// res.Volume = armnetworkcloud.Volume{
	// 	Name: to.Ptr("volumeName"),
	// 	Type: to.Ptr("Microsoft.NetworkCloud/volumes"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/volumes/volumeName"),
	// 	SystemData: &armnetworkcloud.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
	// 		CreatedBy: to.Ptr("identityA"),
	// 		CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("identityB"),
	// 		LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("location"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("myvalue1"),
	// 		"key2": to.Ptr("myvalue2"),
	// 	},
	// 	ExtendedLocation: &armnetworkcloud.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
	// 		Type: to.Ptr("CustomLocation"),
	// 	},
	// 	Properties: &armnetworkcloud.VolumeProperties{
	// 		AttachedTo: []*string{
	// 			to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/virtualMachines/virtualMachineName")},
	// 			DetailedStatus: to.Ptr(armnetworkcloud.VolumeDetailedStatusActive),
	// 			DetailedStatusMessage: to.Ptr("Volume is available"),
	// 			ProvisioningState: to.Ptr(armnetworkcloud.VolumeProvisioningStateSucceeded),
	// 			SerialNumber: to.Ptr("XXXX-XXXX-XXXX-XXXX"),
	// 			SizeMiB: to.Ptr[int64](10000),
	// 		},
	// 	}
}
