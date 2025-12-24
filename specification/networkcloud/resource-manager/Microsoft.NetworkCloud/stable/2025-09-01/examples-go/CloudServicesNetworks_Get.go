package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/46c51b03d99b113ecc3b38883e3cb2d395fe94a4/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/CloudServicesNetworks_Get.json
func ExampleCloudServicesNetworksClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCloudServicesNetworksClient().Get(ctx, "resourceGroupName", "cloudServicesNetworkName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CloudServicesNetwork = armnetworkcloud.CloudServicesNetwork{
	// 	Name: to.Ptr("cloudServicesNetworkName"),
	// 	Type: to.Ptr("Microsoft.NetworkCloud/cloudServicesNetworks"),
	// 	ID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/cloudServicesNetworks/cloudServicesNetworkName"),
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
	// 		Name: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
	// 		Type: to.Ptr("CustomLocation"),
	// 	},
	// 	Properties: &armnetworkcloud.CloudServicesNetworkProperties{
	// 		AdditionalEgressEndpoints: []*armnetworkcloud.EgressEndpoint{
	// 			{
	// 				Category: to.Ptr("azure-resource-management"),
	// 				Endpoints: []*armnetworkcloud.EndpointDependency{
	// 					{
	// 						DomainName: to.Ptr("storageaccountex.blob.core.windows.net"),
	// 						Port: to.Ptr[int64](443),
	// 				}},
	// 		}},
	// 		AssociatedResourceIDs: []*string{
	// 			to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/virtualMachines/virtualMachineName")},
	// 			ClusterID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/clusters/clusterName"),
	// 			DetailedStatus: to.Ptr(armnetworkcloud.CloudServicesNetworkDetailedStatusAvailable),
	// 			DetailedStatusMessage: to.Ptr("Cloud services network is up"),
	// 			EnableDefaultEgressEndpoints: to.Ptr(armnetworkcloud.CloudServicesNetworkEnableDefaultEgressEndpointsFalse),
	// 			EnabledEgressEndpoints: []*armnetworkcloud.EgressEndpoint{
	// 				{
	// 					Category: to.Ptr("azure-resource-management"),
	// 					Endpoints: []*armnetworkcloud.EndpointDependency{
	// 						{
	// 							DomainName: to.Ptr("storageaccountex.blob.core.windows.net"),
	// 							Port: to.Ptr[int64](443),
	// 					}},
	// 			}},
	// 			InterfaceName: to.Ptr("eth0"),
	// 			ProvisioningState: to.Ptr(armnetworkcloud.CloudServicesNetworkProvisioningStateSucceeded),
	// 			StorageOptions: &armnetworkcloud.CloudServicesNetworkStorageOptions{
	// 				Mode: to.Ptr(armnetworkcloud.CloudServicesNetworkStorageModeStandard),
	// 				SizeMiB: to.Ptr[int64](1048576),
	// 				StorageApplianceID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/StorageAppliances/storageApplianceName"),
	// 			},
	// 			StorageStatus: &armnetworkcloud.CloudServicesNetworkStorageStatus{
	// 				Mode: to.Ptr(armnetworkcloud.CloudServicesNetworkStorageModeStandard),
	// 				SizeMiB: to.Ptr[int64](1048576),
	// 				Status: to.Ptr(armnetworkcloud.CloudServicesNetworkStorageStatusStatusAvailable),
	// 				StatusMessage: to.Ptr("Storage is available"),
	// 				VolumeID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/Volumes/volumeName"),
	// 			},
	// 		},
	// 	}
}
