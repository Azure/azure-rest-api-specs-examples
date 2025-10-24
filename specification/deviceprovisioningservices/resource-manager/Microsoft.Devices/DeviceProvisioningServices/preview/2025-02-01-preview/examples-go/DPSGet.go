package armdeviceprovisioningservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices"
)

// Generated from example definition: 2025-02-01-preview/DPSGet.json
func ExampleIotDpsResourceClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceprovisioningservices.NewClientFactory("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIotDpsResourceClient().Get(ctx, "myFirstProvisioningService", "myResourceGroup", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdeviceprovisioningservices.IotDpsResourceClientGetResponse{
	// 	ProvisioningServiceDescription: &armdeviceprovisioningservices.ProvisioningServiceDescription{
	// 		Name: to.Ptr("myFirstProvisioningService"),
	// 		Type: to.Ptr("Microsoft.Devices/ProvisioningServices"),
	// 		Etag: to.Ptr("AAAAAAAADGk="),
	// 		ID: to.Ptr("/subscriptions/91d12660-3dec-467a-be2a-213b5544ddc0/resourceGroups/myResourceGroup/providers/Microsoft.Devices/ProvisioningServices/myFirstProvisioningService"),
	// 		Identity: &armdeviceprovisioningservices.ManagedServiceIdentity{
	// 			Type: to.Ptr(armdeviceprovisioningservices.ManagedServiceIdentityTypeSystemAssignedUserAssigned),
	// 			PrincipalID: to.Ptr("aa80bd74-a3f0-4f14-b9da-99c5351cf9d5"),
	// 			TenantID: to.Ptr("f686d426-8d16-42db-81b7-ab578e110ccd"),
	// 			UserAssignedIdentities: map[string]*armdeviceprovisioningservices.UserAssignedIdentity{
	// 				"/subscriptions/91d12660-3dec-467a-be2a-213b5544ddc0/resourcegroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testidentity": &armdeviceprovisioningservices.UserAssignedIdentity{
	// 					ClientID: to.Ptr("c38f618d-47f6-4260-8b3d-1dd8c130f323"),
	// 					PrincipalID: to.Ptr("f1b0b133-10dc-4985-966f-a98a04675fe9"),
	// 				},
	// 			},
	// 		},
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armdeviceprovisioningservices.IotDpsPropertiesDescription{
	// 			AllocationPolicy: to.Ptr(armdeviceprovisioningservices.AllocationPolicyHashed),
	// 			AuthorizationPolicies: []*armdeviceprovisioningservices.SharedAccessSignatureAuthorizationRuleAccessRightsDescription{
	// 			},
	// 			DeviceProvisioningHostName: to.Ptr("global.azure-devices-provisioning.net"),
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
