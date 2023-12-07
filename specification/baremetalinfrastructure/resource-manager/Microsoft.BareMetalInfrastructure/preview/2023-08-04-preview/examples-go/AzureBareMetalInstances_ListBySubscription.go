package armbaremetalinfrastructure_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/baremetalinfrastructure/armbaremetalinfrastructure/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/90115af9fda46f323e5c42c274f2b376108d1d47/specification/baremetalinfrastructure/resource-manager/Microsoft.BareMetalInfrastructure/preview/2023-08-04-preview/examples/AzureBareMetalInstances_ListBySubscription.json
func ExampleAzureBareMetalInstancesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbaremetalinfrastructure.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAzureBareMetalInstancesClient().NewListBySubscriptionPager(nil)
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
		// page.AzureBareMetalInstancesListResult = armbaremetalinfrastructure.AzureBareMetalInstancesListResult{
		// 	Value: []*armbaremetalinfrastructure.AzureBareMetalInstance{
		// 		{
		// 			Name: to.Ptr("myAzureBareMetalInstance1"),
		// 			Type: to.Ptr("Microsoft.BareMetalInfrastructure/bareMetalInstances"),
		// 			ID: to.Ptr("/subscriptions/57d3422f-467a-448e-b798-ebf490849542/resourceGroups/myResourceGroup/providers/Microsoft.BareMetalInfrastructure/bareMetalInstances/myAzureBareMetalInstance1"),
		// 			SystemData: &armbaremetalinfrastructure.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-21T08:01:22.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@microsoft.com"),
		// 				CreatedByType: to.Ptr(armbaremetalinfrastructure.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-13T08:01:22.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@microsoft.com"),
		// 				LastModifiedByType: to.Ptr(armbaremetalinfrastructure.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("westus2"),
		// 			Tags: map[string]*string{
		// 				"key": to.Ptr("value"),
		// 			},
		// 			Properties: &armbaremetalinfrastructure.AzureBareMetalInstanceProperties{
		// 				AzureBareMetalInstanceID: to.Ptr("23415635-4d7e-41dc-9598-8194f22c24e1"),
		// 				HardwareProfile: &armbaremetalinfrastructure.HardwareProfile{
		// 					AzureBareMetalInstanceSize: to.Ptr(armbaremetalinfrastructure.AzureBareMetalInstanceSizeNamesEnumS72),
		// 					HardwareType: to.Ptr(armbaremetalinfrastructure.AzureBareMetalHardwareTypeNamesEnumCiscoUCS),
		// 				},
		// 				HwRevision: to.Ptr("Rev 4.2"),
		// 				NetworkProfile: &armbaremetalinfrastructure.NetworkProfile{
		// 					CircuitID: to.Ptr("/subscriptions/57d3422f-467a-448e-b798-ebf490849542/resourceGroups/myResourceGroup/providers/Microsoft.Network/expressRouteCircuits/myCircuitId"),
		// 					NetworkInterfaces: []*armbaremetalinfrastructure.NetworkInterface{
		// 						{
		// 							IPAddress: to.Ptr("123.123.123.123"),
		// 					}},
		// 				},
		// 				OSProfile: &armbaremetalinfrastructure.OSProfile{
		// 					ComputerName: to.Ptr("myComputerName"),
		// 					OSType: to.Ptr("SLES 12 SP2"),
		// 					Version: to.Ptr("12 SP2"),
		// 				},
		// 				PowerState: to.Ptr(armbaremetalinfrastructure.AzureBareMetalInstancePowerStateEnumRestarting),
		// 				ProvisioningState: to.Ptr(armbaremetalinfrastructure.AzureBareMetalProvisioningStatesEnumSucceeded),
		// 				StorageProfile: &armbaremetalinfrastructure.StorageProfile{
		// 					NfsIPAddress: to.Ptr("123.123.119.123"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myAzureBareMetalInstance2"),
		// 			Type: to.Ptr("Microsoft.BareMetalInfrastructure/bareMetalInstances"),
		// 			ID: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.BareMetalInfrastructure/bareMetalInstances/myAzureBareMetalInstance2"),
		// 			SystemData: &armbaremetalinfrastructure.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-21T08:01:22.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@microsoft.com"),
		// 				CreatedByType: to.Ptr(armbaremetalinfrastructure.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-13T08:01:22.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@microsoft.com"),
		// 				LastModifiedByType: to.Ptr(armbaremetalinfrastructure.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("westus2"),
		// 			Tags: map[string]*string{
		// 				"key": to.Ptr("value"),
		// 			},
		// 			Properties: &armbaremetalinfrastructure.AzureBareMetalInstanceProperties{
		// 				AzureBareMetalInstanceID: to.Ptr("589bce49-9fe6-4dc8-82df-cf6ae25e0cb9"),
		// 				HardwareProfile: &armbaremetalinfrastructure.HardwareProfile{
		// 					AzureBareMetalInstanceSize: to.Ptr(armbaremetalinfrastructure.AzureBareMetalInstanceSizeNamesEnumS72),
		// 					HardwareType: to.Ptr(armbaremetalinfrastructure.AzureBareMetalHardwareTypeNamesEnumCiscoUCS),
		// 				},
		// 				HwRevision: to.Ptr("Rev 4.2"),
		// 				NetworkProfile: &armbaremetalinfrastructure.NetworkProfile{
		// 					CircuitID: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.Network/expressRouteCircuits/myCircuitId"),
		// 					NetworkInterfaces: []*armbaremetalinfrastructure.NetworkInterface{
		// 						{
		// 							IPAddress: to.Ptr("123.123.123.123"),
		// 					}},
		// 				},
		// 				OSProfile: &armbaremetalinfrastructure.OSProfile{
		// 					ComputerName: to.Ptr("myComputerName2"),
		// 					OSType: to.Ptr("SLES 12 SP2"),
		// 					Version: to.Ptr("12 SP2"),
		// 				},
		// 				PowerState: to.Ptr(armbaremetalinfrastructure.AzureBareMetalInstancePowerStateEnumRestarting),
		// 				ProvisioningState: to.Ptr(armbaremetalinfrastructure.AzureBareMetalProvisioningStatesEnumSucceeded),
		// 				StorageProfile: &armbaremetalinfrastructure.StorageProfile{
		// 					NfsIPAddress: to.Ptr("123.123.119.123"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
