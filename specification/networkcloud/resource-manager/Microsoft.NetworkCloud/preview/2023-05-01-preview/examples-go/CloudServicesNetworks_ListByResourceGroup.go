package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/CloudServicesNetworks_ListByResourceGroup.json
func ExampleCloudServicesNetworksClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCloudServicesNetworksClient().NewListByResourceGroupPager("resourceGroupName", nil)
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
		// page.CloudServicesNetworkList = armnetworkcloud.CloudServicesNetworkList{
		// 	Value: []*armnetworkcloud.CloudServicesNetwork{
		// 		{
		// 			Name: to.Ptr("cloudServicesNetworkName"),
		// 			Type: to.Ptr("Microsoft.NetworkCloud/cloudServicesNetworks"),
		// 			ID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/cloudServicesNetworks/cloudServicesNetworkName"),
		// 			SystemData: &armnetworkcloud.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
		// 				CreatedBy: to.Ptr("identityA"),
		// 				CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("identityB"),
		// 				LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("location"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("myvalue1"),
		// 				"key2": to.Ptr("myvalue2"),
		// 			},
		// 			ExtendedLocation: &armnetworkcloud.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
		// 				Type: to.Ptr("CustomLocation"),
		// 			},
		// 			Properties: &armnetworkcloud.CloudServicesNetworkProperties{
		// 				AdditionalEgressEndpoints: []*armnetworkcloud.EgressEndpoint{
		// 					{
		// 						Category: to.Ptr("azure-resource-management"),
		// 						Endpoints: []*armnetworkcloud.EndpointDependency{
		// 							{
		// 								DomainName: to.Ptr("https://storageaccountex.blob.core.windows.net"),
		// 								Port: to.Ptr[int64](443),
		// 						}},
		// 				}},
		// 				AssociatedResourceIDs: []*string{
		// 					to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/virtualMachines/virtualMachineName")},
		// 					ClusterID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/clusters/clusterName"),
		// 					DetailedStatus: to.Ptr(armnetworkcloud.CloudServicesNetworkDetailedStatusAvailable),
		// 					DetailedStatusMessage: to.Ptr("Cloud services network is up"),
		// 					EnableDefaultEgressEndpoints: to.Ptr(armnetworkcloud.CloudServicesNetworkEnableDefaultEgressEndpointsFalse),
		// 					EnabledEgressEndpoints: []*armnetworkcloud.EgressEndpoint{
		// 						{
		// 							Category: to.Ptr("azure-resource-management"),
		// 							Endpoints: []*armnetworkcloud.EndpointDependency{
		// 								{
		// 									DomainName: to.Ptr("https://storageaccountex.blob.core.windows.net"),
		// 									Port: to.Ptr[int64](443),
		// 							}},
		// 					}},
		// 					InterfaceName: to.Ptr("eth0"),
		// 					ProvisioningState: to.Ptr(armnetworkcloud.CloudServicesNetworkProvisioningStateSucceeded),
		// 				},
		// 		}},
		// 	}
	}
}
