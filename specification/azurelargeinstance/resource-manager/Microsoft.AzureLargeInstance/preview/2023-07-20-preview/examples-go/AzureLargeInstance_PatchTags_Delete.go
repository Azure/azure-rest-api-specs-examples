package armlargeinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/largeinstance/armlargeinstance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6051d2b126f5b1e4b623cde8edfa3e25cf730685/specification/azurelargeinstance/resource-manager/Microsoft.AzureLargeInstance/preview/2023-07-20-preview/examples/AzureLargeInstance_PatchTags_Delete.json
func ExampleAzureLargeInstanceClient_Update_azureLargeInstanceDeleteTag() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlargeinstance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAzureLargeInstanceClient().Update(ctx, "myResourceGroup", "myALInstance", armlargeinstance.AzureLargeInstanceTagsUpdate{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AzureLargeInstance = armlargeinstance.AzureLargeInstance{
	// 	Name: to.Ptr("myALInstance"),
	// 	Type: to.Ptr("Microsoft.AzureLargeInstance/AzureLargeInstances"),
	// 	ID: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.AzureLargeInstance/AzureLargeInstances/myALInstance"),
	// 	SystemData: &armlargeinstance.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-20T23:10:22.682Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@microsoft.com"),
	// 		CreatedByType: to.Ptr(armlargeinstance.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-21T08:01:22.000Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@microsoft.com"),
	// 		LastModifiedByType: to.Ptr(armlargeinstance.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armlargeinstance.AzureLargeInstanceProperties{
	// 		AzureLargeInstanceID: to.Ptr("23415635-4d7e-41dc-9598-8194f22c24e1"),
	// 		HardwareProfile: &armlargeinstance.HardwareProfile{
	// 			AzureLargeInstanceSize: to.Ptr(armlargeinstance.AzureLargeInstanceSizeNamesEnumS72),
	// 			HardwareType: to.Ptr(armlargeinstance.AzureLargeInstanceHardwareTypeNamesEnumCiscoUCS),
	// 		},
	// 		HwRevision: to.Ptr("Rev 3"),
	// 		NetworkProfile: &armlargeinstance.NetworkProfile{
	// 			CircuitID: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.Network/expressRouteCircuit"),
	// 			NetworkInterfaces: []*armlargeinstance.IPAddress{
	// 				{
	// 					IPAddress: to.Ptr("100.100.100.100"),
	// 			}},
	// 		},
	// 		OSProfile: &armlargeinstance.OsProfile{
	// 			ComputerName: to.Ptr("myComputerName"),
	// 			OSType: to.Ptr("SUSE"),
	// 			SSHPublicKey: to.Ptr("{ssh-rsa public key}"),
	// 			Version: to.Ptr("12 SP1"),
	// 		},
	// 		PartnerNodeID: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.AzureLargeInstance/AzureLargeInstances/myALInstance2"),
	// 		PowerState: to.Ptr(armlargeinstance.AzureLargeInstancePowerStateEnumStarted),
	// 		ProvisioningState: to.Ptr(armlargeinstance.AzureLargeInstanceProvisioningStatesEnumSucceeded),
	// 		ProximityPlacementGroup: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.Compute/proximityPlacementGroups/myplacementgroup"),
	// 		StorageProfile: &armlargeinstance.StorageProfile{
	// 			NfsIPAddress: to.Ptr("200.200.200.200"),
	// 		},
	// 	},
	// }
}
