package armdeviceprovisioningservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices"
)

// Generated from example definition: 2025-02-01-preview/DPSCreate.json
func ExampleIotDpsResourceClient_BeginCreateOrUpdate_dpsCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceprovisioningservices.NewClientFactory("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIotDpsResourceClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myFirstProvisioningService", armdeviceprovisioningservices.ProvisioningServiceDescription{
		Location: to.Ptr("East US"),
		Properties: &armdeviceprovisioningservices.IotDpsPropertiesDescription{
			EnableDataResidency: to.Ptr(false),
		},
		SKU: &armdeviceprovisioningservices.IotDpsSKUInfo{
			Name:     to.Ptr(armdeviceprovisioningservices.IotDpsSKUS1),
			Capacity: to.Ptr[int64](1),
		},
		Tags: map[string]*string{},
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
	// res = armdeviceprovisioningservices.IotDpsResourceClientCreateOrUpdateResponse{
	// 	ProvisioningServiceDescription: &armdeviceprovisioningservices.ProvisioningServiceDescription{
	// 		Name: to.Ptr("myFirstProvisioningService"),
	// 		Type: to.Ptr("Microsoft.Devices/ProvisioningServices"),
	// 		Etag: to.Ptr("AAAAAAAADGk="),
	// 		ID: to.Ptr("/subscriptions/91d12660-3dec-467a-be2a-213b5544ddc0/resourceGroups//providers/Microsoft.Devices/ProvisioningServices/myFirstProvisioningService"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armdeviceprovisioningservices.IotDpsPropertiesDescription{
	// 			AllocationPolicy: to.Ptr(armdeviceprovisioningservices.AllocationPolicyHashed),
	// 			AuthorizationPolicies: []*armdeviceprovisioningservices.SharedAccessSignatureAuthorizationRuleAccessRightsDescription{
	// 			},
	// 			DeviceProvisioningHostName: to.Ptr("global.azure-devices-provisioning.net"),
	// 			EnableDataResidency: to.Ptr(false),
	// 			IDScope: to.Ptr("0ne00000012"),
	// 			PortalOperationsHostName: to.Ptr("myFirstProvisioningService.services.azure-devices-provisioning.net"),
	// 			ServiceOperationsHostName: to.Ptr("myFirstProvisioningService.azure-devices-provisioning.net"),
	// 			State: to.Ptr(armdeviceprovisioningservices.StateActive),
	// 		},
	// 		Resourcegroup: to.Ptr("myResourceGroup"),
	// 		SKU: &armdeviceprovisioningservices.IotDpsSKUInfo{
	// 			Name: to.Ptr(armdeviceprovisioningservices.IotDpsSKUS1),
	// 			Capacity: to.Ptr[int64](1),
	// 			Tier: to.Ptr("Standard"),
	// 		},
	// 		Subscriptionid: to.Ptr("91d12660-3dec-467a-be2a-213b5544ddc0"),
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
