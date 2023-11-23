package armhybridcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4bb583bcb67c2bf448712f2bd1593a64a7a8f139/specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2023-11-15-preview/examples/GetVmSkus.json
func ExampleClient_GetVMSKUs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().GetVMSKUs(ctx, "subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VMSKUProfile = armhybridcontainerservice.VMSKUProfile{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("microsoft.hybridcontainerservice/skus"),
	// 	ID: to.Ptr("/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation/providers/Microsoft.HybridContainerService/skus/default"),
	// 	ExtendedLocation: &armhybridcontainerservice.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourcegroups/test-arcappliance-resgrp/providers/microsoft.extendedlocation/customlocations/testcustomlocation"),
	// 		Type: to.Ptr(armhybridcontainerservice.ExtendedLocationTypesCustomLocation),
	// 	},
	// 	Properties: &armhybridcontainerservice.VMSKUProfileProperties{
	// 		ProvisioningState: to.Ptr(armhybridcontainerservice.ResourceProvisioningStateSucceeded),
	// 		Values: []*armhybridcontainerservice.VMSKUProperties{
	// 			{
	// 				Name: to.Ptr("Standard_A0"),
	// 				Capabilities: []*armhybridcontainerservice.VMSKUCapabilities{
	// 					{
	// 						Name: to.Ptr("vCpu"),
	// 						Value: to.Ptr("2"),
	// 					},
	// 					{
	// 						Name: to.Ptr("MemoryMb"),
	// 						Value: to.Ptr("2345"),
	// 					},
	// 					{
	// 						Name: to.Ptr("DiskSizeGb"),
	// 						Value: to.Ptr("128"),
	// 					},
	// 					{
	// 						Name: to.Ptr("GpuCount"),
	// 						Value: to.Ptr("1"),
	// 					},
	// 					{
	// 						Name: to.Ptr("GpuNameType"),
	// 						Value: to.Ptr("NVIDIA Tesla T4"),
	// 					},
	// 					{
	// 						Name: to.Ptr("GpuAssignMode"),
	// 						Value: to.Ptr("1"),
	// 					},
	// 					{
	// 						Name: to.Ptr("Provider"),
	// 						Value: to.Ptr("HCI"),
	// 				}},
	// 				ResourceType: to.Ptr("VirtualMachines"),
	// 				Size: to.Ptr("A0"),
	// 				Tier: to.Ptr("Standard"),
	// 		}},
	// 	},
	// }
}
