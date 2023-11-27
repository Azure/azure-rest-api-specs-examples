package armbaremetalinfrastructure_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/baremetalinfrastructure/armbaremetalinfrastructure"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/baremetalinfrastructure/resource-manager/Microsoft.BareMetalInfrastructure/stable/2021-08-09/examples/AzureBareMetalInstances_Get.json
func ExampleAzureBareMetalInstancesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbaremetalinfrastructure.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAzureBareMetalInstancesClient().Get(ctx, "myResourceGroup", "myAzureBareMetalInstance", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AzureBareMetalInstance = armbaremetalinfrastructure.AzureBareMetalInstance{
	// 	Name: to.Ptr("myAzureBareMetalInstance"),
	// 	Type: to.Ptr("Microsoft.BareMetalInfrastructure/bareMetalInstances"),
	// 	ID: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.BareMetalInfrastructure/bareMetalInstances/myAzureBareMetalInstance"),
	// 	Location: to.Ptr("westus2"),
	// 	Tags: map[string]*string{
	// 		"key": to.Ptr("value"),
	// 	},
	// 	Properties: &armbaremetalinfrastructure.AzureBareMetalInstanceProperties{
	// 		AzureBareMetalInstanceID: to.Ptr("23415635-4d7e-41dc-9598-8194f22c24e1"),
	// 		HardwareProfile: &armbaremetalinfrastructure.HardwareProfile{
	// 			AzureBareMetalInstanceSize: to.Ptr(armbaremetalinfrastructure.AzureBareMetalInstanceSizeNamesEnumS72),
	// 			HardwareType: to.Ptr(armbaremetalinfrastructure.AzureBareMetalHardwareTypeNamesEnumCiscoUCS),
	// 		},
	// 		HwRevision: to.Ptr("Rev 4.2"),
	// 		NetworkProfile: &armbaremetalinfrastructure.NetworkProfile{
	// 			CircuitID: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.Network/expressRouteCircuits/myCircuitId"),
	// 			NetworkInterfaces: []*armbaremetalinfrastructure.IPAddress{
	// 				{
	// 					IPAddress: to.Ptr("123.123.123.123"),
	// 			}},
	// 		},
	// 		OSProfile: &armbaremetalinfrastructure.OSProfile{
	// 			ComputerName: to.Ptr("myComputerName"),
	// 			OSType: to.Ptr("SLES 12 SP2"),
	// 			Version: to.Ptr("12 SP2"),
	// 		},
	// 		PowerState: to.Ptr(armbaremetalinfrastructure.AzureBareMetalInstancePowerStateEnumRestarting),
	// 		ProvisioningState: to.Ptr(armbaremetalinfrastructure.AzureBareMetalProvisioningStatesEnumSucceeded),
	// 		StorageProfile: &armbaremetalinfrastructure.StorageProfile{
	// 			NfsIPAddress: to.Ptr("123.123.119.123"),
	// 		},
	// 	},
	// 	SystemData: &armbaremetalinfrastructure.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-20T23:10:22.682Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@microsoft.com"),
	// 		CreatedByType: to.Ptr(armbaremetalinfrastructure.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-21T08:01:22.000Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@microsoft.com"),
	// 		LastModifiedByType: to.Ptr(armbaremetalinfrastructure.CreatedByTypeUser),
	// 	},
	// }
}
