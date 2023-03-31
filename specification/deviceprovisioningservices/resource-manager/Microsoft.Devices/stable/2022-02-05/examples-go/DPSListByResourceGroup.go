package armdeviceprovisioningservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2022-02-05/examples/DPSListByResourceGroup.json
func ExampleIotDpsResourceClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceprovisioningservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIotDpsResourceClient().NewListByResourceGroupPager("myResourceGroup", nil)
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
		// page.ProvisioningServiceDescriptionListResult = armdeviceprovisioningservices.ProvisioningServiceDescriptionListResult{
		// 	Value: []*armdeviceprovisioningservices.ProvisioningServiceDescription{
		// 		{
		// 			Name: to.Ptr("myFirstProvisioningService"),
		// 			Type: to.Ptr("Microsoft.Devices/ProvisioningServices"),
		// 			ID: to.Ptr("/subscriptions/91d12660-3dec-467a-be2a-213b5544ddc0/resourceGroups/myResourceGroup/providers/Microsoft.Devices/ProvisioningServices/myFirstProvisioningService"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("AAAAAAAADGk="),
		// 			Properties: &armdeviceprovisioningservices.IotDpsPropertiesDescription{
		// 				AllocationPolicy: to.Ptr(armdeviceprovisioningservices.AllocationPolicyHashed),
		// 				DeviceProvisioningHostName: to.Ptr("global.azure-devices-provisioning.net"),
		// 				IDScope: to.Ptr("0ne00000012"),
		// 				ServiceOperationsHostName: to.Ptr("myFirstProvisioningService.azure-devices-provisioning.net"),
		// 				State: to.Ptr(armdeviceprovisioningservices.StateActive),
		// 			},
		// 			SKU: &armdeviceprovisioningservices.IotDpsSKUInfo{
		// 				Name: to.Ptr(armdeviceprovisioningservices.IotDpsSKUS1),
		// 				Capacity: to.Ptr[int64](1),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("mySecondProvisioningService"),
		// 			Type: to.Ptr("Microsoft.Devices/ProvisioningServices"),
		// 			ID: to.Ptr("/subscriptions/91d12660-3dec-467a-be2a-213b5544ddc0/resourceGroups/myResourceGroup/providers/Microsoft.Devices/ProvisioningServices/mySecondProvisioningService"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("AAAAAAAADGk="),
		// 			Properties: &armdeviceprovisioningservices.IotDpsPropertiesDescription{
		// 				AllocationPolicy: to.Ptr(armdeviceprovisioningservices.AllocationPolicyHashed),
		// 				DeviceProvisioningHostName: to.Ptr("global.azure-devices-provisioning.net"),
		// 				IDScope: to.Ptr("0ne00000012"),
		// 				ServiceOperationsHostName: to.Ptr("mySecondProvisioningService.azure-devices-provisioning.net"),
		// 				State: to.Ptr(armdeviceprovisioningservices.StateActive),
		// 			},
		// 			SKU: &armdeviceprovisioningservices.IotDpsSKUInfo{
		// 				Name: to.Ptr(armdeviceprovisioningservices.IotDpsSKUS1),
		// 				Capacity: to.Ptr[int64](1),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 	}},
		// }
	}
}
