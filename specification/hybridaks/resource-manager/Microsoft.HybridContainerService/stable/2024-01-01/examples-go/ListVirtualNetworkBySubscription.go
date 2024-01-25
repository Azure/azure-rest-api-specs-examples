package armhybridcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/41e4538ed7bb3ceac3c1322c9455a0812ed110ac/specification/hybridaks/resource-manager/Microsoft.HybridContainerService/stable/2024-01-01/examples/ListVirtualNetworkBySubscription.json
func ExampleVirtualNetworksClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualNetworksClient().NewListBySubscriptionPager(nil)
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
		// page.VirtualNetworksListResult = armhybridcontainerservice.VirtualNetworksListResult{
		// 	Value: []*armhybridcontainerservice.VirtualNetwork{
		// 		{
		// 			Name: to.Ptr("test-vnet-static"),
		// 			Type: to.Ptr("microsoft.hybridcontainerservice/virtualnetworks"),
		// 			ID: to.Ptr("/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.HybridContainerService/virtualNetworks/test-vnet-static"),
		// 			Location: to.Ptr("westus"),
		// 			ExtendedLocation: &armhybridcontainerservice.VirtualNetworkExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourcegroups/test-arcappliance-resgrp/providers/microsoft.extendedlocation/customlocations/testcustomlocation"),
		// 				Type: to.Ptr(armhybridcontainerservice.ExtendedLocationTypesCustomLocation),
		// 			},
		// 			Properties: &armhybridcontainerservice.VirtualNetworkProperties{
		// 				DNSServers: []*string{
		// 					to.Ptr("192.168.0.1")},
		// 					Gateway: to.Ptr("192.168.0.1"),
		// 					InfraVnetProfile: &armhybridcontainerservice.VirtualNetworkPropertiesInfraVnetProfile{
		// 						Hci: &armhybridcontainerservice.VirtualNetworkPropertiesInfraVnetProfileHci{
		// 							MocGroup: to.Ptr("target-group"),
		// 							MocLocation: to.Ptr("MocLocation"),
		// 							MocVnetName: to.Ptr("vnet1"),
		// 						},
		// 					},
		// 					IPAddressPrefix: to.Ptr("192.168.0.0/16"),
		// 					ProvisioningState: to.Ptr(armhybridcontainerservice.ProvisioningStateSucceeded),
		// 					VipPool: []*armhybridcontainerservice.VirtualNetworkPropertiesVipPoolItem{
		// 						{
		// 							EndIP: to.Ptr("192.168.0.50"),
		// 							StartIP: to.Ptr("192.168.0.10"),
		// 					}},
		// 					VlanID: to.Ptr[int32](10),
		// 					VmipPool: []*armhybridcontainerservice.VirtualNetworkPropertiesVmipPoolItem{
		// 						{
		// 							EndIP: to.Ptr("192.168.0.130"),
		// 							StartIP: to.Ptr("192.168.0.110"),
		// 					}},
		// 				},
		// 		}},
		// 	}
	}
}
