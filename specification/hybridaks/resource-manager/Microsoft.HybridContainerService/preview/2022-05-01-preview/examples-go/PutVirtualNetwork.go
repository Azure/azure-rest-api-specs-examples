package armhybridcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-05-01-preview/examples/PutVirtualNetwork.json
func ExampleVirtualNetworksClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armhybridcontainerservice.NewVirtualNetworksClient("a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "test-arcappliance-resgrp", "test-vnet-static", armhybridcontainerservice.VirtualNetworks{
		Location: to.Ptr("westus"),
		ExtendedLocation: &armhybridcontainerservice.VirtualNetworksExtendedLocation{
			Name: to.Ptr("/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourcegroups/test-arcappliance-resgrp/providers/microsoft.extendedlocation/customlocations/testcustomlocation"),
			Type: to.Ptr("CustomLocation"),
		},
		Properties: &armhybridcontainerservice.VirtualNetworksProperties{
			InfraVnetProfile: &armhybridcontainerservice.VirtualNetworksPropertiesInfraVnetProfile{
				Hci: &armhybridcontainerservice.VirtualNetworksPropertiesInfraVnetProfileHci{
					MocGroup:    to.Ptr("target-group"),
					MocLocation: to.Ptr("MocLocation"),
					MocVnetName: to.Ptr("test-vnet"),
				},
			},
			VipPool: []*armhybridcontainerservice.VirtualNetworksPropertiesVipPoolItem{
				{
					EndIP:   to.Ptr("192.168.0.50"),
					StartIP: to.Ptr("192.168.0.10"),
				}},
			VmipPool: []*armhybridcontainerservice.VirtualNetworksPropertiesVmipPoolItem{
				{
					EndIP:   to.Ptr("192.168.0.130"),
					StartIP: to.Ptr("192.168.0.110"),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
