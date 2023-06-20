package armdeviceprovisioningservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0d41e635294dce73dfa99b07f3da4b68a9c9e29c/specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/preview/2023-03-01-preview/examples/DPSCreate.json
func ExampleIotDpsResourceClient_BeginCreateOrUpdate_dpsCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceprovisioningservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIotDpsResourceClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myFirstProvisioningService", armdeviceprovisioningservices.ProvisioningServiceDescription{
		Location: to.Ptr("East US"),
		Tags:     map[string]*string{},
		Properties: &armdeviceprovisioningservices.IotDpsPropertiesDescription{
			EnableDataResidency: to.Ptr(false),
		},
		SKU: &armdeviceprovisioningservices.IotDpsSKUInfo{
			Name:     to.Ptr(armdeviceprovisioningservices.IotDpsSKUS1),
			Capacity: to.Ptr[int64](1),
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
	// res.ProvisioningServiceDescription = armdeviceprovisioningservices.ProvisioningServiceDescription{
	// 	Name: to.Ptr("myFirstProvisioningService"),
	// 	Type: to.Ptr("Microsoft.Devices/ProvisioningServices"),
	// 	ID: to.Ptr("/subscriptions/91d12660-3dec-467a-be2a-213b5544ddc0/resourceGroups//providers/Microsoft.Devices/ProvisioningServices/myFirstProvisioningService"),
	// 	Location: to.Ptr("eastus"),
	// 	Resourcegroup: to.Ptr("myResourceGroup"),
	// 	Subscriptionid: to.Ptr("91d12660-3dec-467a-be2a-213b5544ddc0"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Etag: to.Ptr("AAAAAAAADGk="),
	// 	Properties: &armdeviceprovisioningservices.IotDpsPropertiesDescription{
	// 		AllocationPolicy: to.Ptr(armdeviceprovisioningservices.AllocationPolicyHashed),
	// 		AuthorizationPolicies: []*armdeviceprovisioningservices.SharedAccessSignatureAuthorizationRuleAccessRightsDescription{
	// 		},
	// 		DeviceProvisioningHostName: to.Ptr("global.azure-devices-provisioning.net"),
	// 		EnableDataResidency: to.Ptr(false),
	// 		IDScope: to.Ptr("0ne00000012"),
	// 		PortalOperationsHostName: to.Ptr("myFirstProvisioningService.services.azure-devices-provisioning.net"),
	// 		ServiceOperationsHostName: to.Ptr("myFirstProvisioningService.azure-devices-provisioning.net"),
	// 		State: to.Ptr(armdeviceprovisioningservices.StateActive),
	// 	},
	// 	SKU: &armdeviceprovisioningservices.IotDpsSKUInfo{
	// 		Name: to.Ptr(armdeviceprovisioningservices.IotDpsSKUS1),
	// 		Capacity: to.Ptr[int64](1),
	// 		Tier: to.Ptr("Standard"),
	// 	},
	// }
}
