package armdeviceprovisioningservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices"
)

// Generated from example definition: 2025-02-01-preview/DPSListPrivateLinkResources.json
func ExampleIotDpsResourceClient_ListPrivateLinkResources() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceprovisioningservices.NewClientFactory("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIotDpsResourceClient().ListPrivateLinkResources(ctx, "myResourceGroup", "myFirstProvisioningService", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdeviceprovisioningservices.IotDpsResourceClientListPrivateLinkResourcesResponse{
	// 	PrivateLinkResources: &armdeviceprovisioningservices.PrivateLinkResources{
	// 		Value: []*armdeviceprovisioningservices.GroupIDInformation{
	// 			{
	// 				Name: to.Ptr("iotDps"),
	// 				Type: to.Ptr("Microsoft.Devices/ProvisioningServices/PrivateLinkResources"),
	// 				ID: to.Ptr("/subscriptions/91d12660-3dec-467a-be2a-213b5544ddc0/resourceGroups/myResourceGroup/providers/Microsoft.Devices/ProvisioningServices/myFirstProvisioningService/PrivateLinkResources/iotDps"),
	// 				Properties: &armdeviceprovisioningservices.GroupIDInformationProperties{
	// 					GroupID: to.Ptr("iotDps"),
	// 					RequiredMembers: []*string{
	// 						to.Ptr("iotDps"),
	// 					},
	// 					RequiredZoneNames: []*string{
	// 						to.Ptr("privatelink.azure-devices-provisioning.net"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
