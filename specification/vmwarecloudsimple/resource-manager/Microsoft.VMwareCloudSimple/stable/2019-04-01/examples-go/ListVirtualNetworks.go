package armvmwarecloudsimple_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/ListVirtualNetworks.json
func ExampleVirtualNetworksClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvmwarecloudsimple.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualNetworksClient().NewListPager("westus2", "myPrivateCloud", "/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/resourcePools/resgroup-26", nil)
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
		// page.VirtualNetworkListResponse = armvmwarecloudsimple.VirtualNetworkListResponse{
		// 	Value: []*armvmwarecloudsimple.VirtualNetwork{
		// 		{
		// 			Name: to.Ptr("Datacenter/CS-Management"),
		// 			Type: to.Ptr("Microsoft.VMwareCloudSimple/virtualNetworks"),
		// 			Assignable: to.Ptr(false),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualNetworks/dvportgroup-19"),
		// 			Location: to.Ptr("westus2"),
		// 			Properties: &armvmwarecloudsimple.VirtualNetworkProperties{
		// 				PrivateCloudID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Datacenter/CS-Rescue"),
		// 			Type: to.Ptr("Microsoft.VMwareCloudSimple/virtualNetworks"),
		// 			Assignable: to.Ptr(true),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualNetworks/dvportgroup-20"),
		// 			Location: to.Ptr("westus2"),
		// 			Properties: &armvmwarecloudsimple.VirtualNetworkProperties{
		// 				PrivateCloudID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Datacenter/CS-VSAN"),
		// 			Type: to.Ptr("Microsoft.VMwareCloudSimple/virtualNetworks"),
		// 			Assignable: to.Ptr(false),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualNetworks/dvportgroup-21"),
		// 			Location: to.Ptr("westus2"),
		// 			Properties: &armvmwarecloudsimple.VirtualNetworkProperties{
		// 				PrivateCloudID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Datacenter/CS-VMotion"),
		// 			Type: to.Ptr("Microsoft.VMwareCloudSimple/virtualNetworks"),
		// 			Assignable: to.Ptr(false),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualNetworks/dvportgroup-22"),
		// 			Location: to.Ptr("westus2"),
		// 			Properties: &armvmwarecloudsimple.VirtualNetworkProperties{
		// 				PrivateCloudID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Datacenter/net-01"),
		// 			Type: to.Ptr("Microsoft.VMwareCloudSimple/virtualNetworks"),
		// 			Assignable: to.Ptr(true),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualNetworks/dvportgroup-35"),
		// 			Location: to.Ptr("westus2"),
		// 			Properties: &armvmwarecloudsimple.VirtualNetworkProperties{
		// 				PrivateCloudID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud"),
		// 			},
		// 	}},
		// }
	}
}
