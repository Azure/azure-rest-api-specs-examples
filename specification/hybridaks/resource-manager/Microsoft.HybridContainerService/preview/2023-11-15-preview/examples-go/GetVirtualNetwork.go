package armhybridcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4bb583bcb67c2bf448712f2bd1593a64a7a8f139/specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2023-11-15-preview/examples/GetVirtualNetwork.json
func ExampleVirtualNetworksClient_Retrieve() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualNetworksClient().Retrieve(ctx, "test-arcappliance-resgrp", "test-vnet-static", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualNetwork = armhybridcontainerservice.VirtualNetwork{
	// 	Name: to.Ptr("test-vnet-static"),
	// 	Type: to.Ptr("microsoft.hybridcontainerservice/virtualnetworks"),
	// 	ID: to.Ptr("/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.HybridContainerService/virtualNetworks/test-vnet-static"),
	// 	Location: to.Ptr("westus"),
	// 	ExtendedLocation: &armhybridcontainerservice.VirtualNetworkExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourcegroups/test-arcappliance-resgrp/providers/microsoft.extendedlocation/customlocations/testcustomlocation"),
	// 		Type: to.Ptr("CustomLocation"),
	// 	},
	// 	Properties: &armhybridcontainerservice.VirtualNetworkProperties{
	// 		DNSServers: []*string{
	// 			to.Ptr("192.168.0.1")},
	// 			Gateway: to.Ptr("192.168.0.1"),
	// 			InfraVnetProfile: &armhybridcontainerservice.VirtualNetworkPropertiesInfraVnetProfile{
	// 				Vmware: &armhybridcontainerservice.VirtualNetworkPropertiesInfraVnetProfileVmware{
	// 					SegmentName: to.Ptr("test-network"),
	// 				},
	// 			},
	// 			IPAddressPrefix: to.Ptr("192.168.0.0/16"),
	// 			ProvisioningState: to.Ptr(armhybridcontainerservice.ProvisioningStateSucceeded),
	// 			VipPool: []*armhybridcontainerservice.VirtualNetworkPropertiesVipPoolItem{
	// 				{
	// 					EndIP: to.Ptr("192.168.0.50"),
	// 					StartIP: to.Ptr("192.168.0.10"),
	// 			}},
	// 			VlanID: to.Ptr[int32](10),
	// 			VmipPool: []*armhybridcontainerservice.VirtualNetworkPropertiesVmipPoolItem{
	// 				{
	// 					EndIP: to.Ptr("192.168.0.130"),
	// 					StartIP: to.Ptr("192.168.0.110"),
	// 			}},
	// 		},
	// 	}
}
