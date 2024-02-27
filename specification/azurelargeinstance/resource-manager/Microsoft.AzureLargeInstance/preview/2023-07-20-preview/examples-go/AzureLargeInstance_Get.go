package armlargeinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/largeinstance/armlargeinstance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6051d2b126f5b1e4b623cde8edfa3e25cf730685/specification/azurelargeinstance/resource-manager/Microsoft.AzureLargeInstance/preview/2023-07-20-preview/examples/AzureLargeInstance_Get.json
func ExampleAzureLargeInstanceClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlargeinstance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAzureLargeInstanceClient().Get(ctx, "myResourceGroup", "myAzureLargeInstance", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AzureLargeInstance = armlargeinstance.AzureLargeInstance{
	// 	Name: to.Ptr("myAzureLargeInstance"),
	// 	Type: to.Ptr("Microsoft.AzureLargeInstance/AzureLargeInstances"),
	// 	ID: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.AzureLargeInstance/AzureLargeInstances/myAzureLargeInstance"),
	// 	SystemData: &armlargeinstance.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-20T23:10:22.682Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@microsoft.com"),
	// 		CreatedByType: to.Ptr(armlargeinstance.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-21T08:01:22.000Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@microsoft.com"),
	// 		LastModifiedByType: to.Ptr(armlargeinstance.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("westus2"),
	// 	Tags: map[string]*string{
	// 		"key": to.Ptr("value"),
	// 	},
	// 	Properties: &armlargeinstance.AzureLargeInstanceProperties{
	// 		AzureLargeInstanceID: to.Ptr("23415635-4d7e-41dc-9598-8194f22c24e1"),
	// 		HardwareProfile: &armlargeinstance.HardwareProfile{
	// 			AzureLargeInstanceSize: to.Ptr(armlargeinstance.AzureLargeInstanceSizeNamesEnumS72),
	// 			HardwareType: to.Ptr(armlargeinstance.AzureLargeInstanceHardwareTypeNamesEnumCiscoUCS),
	// 		},
	// 		HwRevision: to.Ptr("Rev 4.2"),
	// 		NetworkProfile: &armlargeinstance.NetworkProfile{
	// 			CircuitID: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.Network/expressRouteCircuits/myCircuitId"),
	// 			NetworkInterfaces: []*armlargeinstance.IPAddress{
	// 				{
	// 					IPAddress: to.Ptr("123.123.123.123"),
	// 			}},
	// 		},
	// 		OSProfile: &armlargeinstance.OsProfile{
	// 			ComputerName: to.Ptr("myComputerName"),
	// 			OSType: to.Ptr("SLES 12 SP2"),
	// 			Version: to.Ptr("12 SP2"),
	// 		},
	// 		PowerState: to.Ptr(armlargeinstance.AzureLargeInstancePowerStateEnumRestarting),
	// 		ProvisioningState: to.Ptr(armlargeinstance.AzureLargeInstanceProvisioningStatesEnumSucceeded),
	// 		StorageProfile: &armlargeinstance.StorageProfile{
	// 			NfsIPAddress: to.Ptr("123.123.119.123"),
	// 		},
	// 	},
	// }
}
