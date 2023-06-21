package armdeviceprovisioningservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0d41e635294dce73dfa99b07f3da4b68a9c9e29c/specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/preview/2023-03-01-preview/examples/DPSListBySubscription.json
func ExampleIotDpsResourceClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceprovisioningservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIotDpsResourceClient().NewListBySubscriptionPager(nil)
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
		// 			Resourcegroup: to.Ptr("myResourceGroup"),
		// 			Subscriptionid: to.Ptr("91d12660-3dec-467a-be2a-213b5544ddc0"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("AAAAAAAADGk="),
		// 			Identity: &armdeviceprovisioningservices.ManagedServiceIdentity{
		// 				Type: to.Ptr(armdeviceprovisioningservices.ManagedServiceIdentityTypeSystemAssignedUserAssigned),
		// 				PrincipalID: to.Ptr("aa80bd74-a3f0-4f14-b9da-99c5351cf9d5"),
		// 				TenantID: to.Ptr("f686d426-8d16-42db-81b7-ab578e110ccd"),
		// 				UserAssignedIdentities: map[string]*armdeviceprovisioningservices.UserAssignedIdentity{
		// 					"/subscriptions/91d12660-3dec-467a-be2a-213b5544ddc0/resourcegroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testidentity": &armdeviceprovisioningservices.UserAssignedIdentity{
		// 						ClientID: to.Ptr("c38f618d-47f6-4260-8b3d-1dd8c130f323"),
		// 						PrincipalID: to.Ptr("f1b0b133-10dc-4985-966f-a98a04675fe9"),
		// 					},
		// 				},
		// 			},
		// 			Properties: &armdeviceprovisioningservices.IotDpsPropertiesDescription{
		// 				AllocationPolicy: to.Ptr(armdeviceprovisioningservices.AllocationPolicyHashed),
		// 				DeviceProvisioningHostName: to.Ptr("global.azure-devices-provisioning.net"),
		// 				IDScope: to.Ptr("0ne00000012"),
		// 				PortalOperationsHostName: to.Ptr("myFirstProvisioningService.services.azure-devices-provisioning.net"),
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
		// 			Resourcegroup: to.Ptr("myResourceGroup"),
		// 			Subscriptionid: to.Ptr("91d12660-3dec-467a-be2a-213b5544ddc0"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("AAAAAAAADGk="),
		// 			Identity: &armdeviceprovisioningservices.ManagedServiceIdentity{
		// 				Type: to.Ptr(armdeviceprovisioningservices.ManagedServiceIdentityTypeSystemAssignedUserAssigned),
		// 				PrincipalID: to.Ptr("01341f2b-d497-4117-b5c1-1f1d50b25444"),
		// 				TenantID: to.Ptr("f686d426-8d16-42db-81b7-ab578e110ccd"),
		// 				UserAssignedIdentities: map[string]*armdeviceprovisioningservices.UserAssignedIdentity{
		// 					"/subscriptions/91d12660-3dec-467a-be2a-213b5544ddc0/resourcegroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testidentity2": &armdeviceprovisioningservices.UserAssignedIdentity{
		// 						ClientID: to.Ptr("8cd6d250-17dd-4c1b-9847-225237b94c55"),
		// 						PrincipalID: to.Ptr("8785a11f-848a-4d5d-a55b-e381b9e15512"),
		// 					},
		// 				},
		// 			},
		// 			Properties: &armdeviceprovisioningservices.IotDpsPropertiesDescription{
		// 				AllocationPolicy: to.Ptr(armdeviceprovisioningservices.AllocationPolicyHashed),
		// 				DeviceProvisioningHostName: to.Ptr("global.azure-devices-provisioning.net"),
		// 				IDScope: to.Ptr("0ne00000012"),
		// 				PortalOperationsHostName: to.Ptr("myFirstProvisioningService.services.azure-devices-provisioning.net"),
		// 				ServiceOperationsHostName: to.Ptr("myFirstProvisioningService.azure-devices-provisioning.net"),
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
