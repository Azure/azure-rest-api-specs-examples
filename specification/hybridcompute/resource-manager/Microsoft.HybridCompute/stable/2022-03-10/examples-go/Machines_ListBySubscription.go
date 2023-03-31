package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/stable/2022-03-10/examples/Machines_ListBySubscription.json
func ExampleMachinesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMachinesClient().NewListBySubscriptionPager(nil)
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
		// page.MachineListResult = armhybridcompute.MachineListResult{
		// 	Value: []*armhybridcompute.Machine{
		// 		{
		// 			Name: to.Ptr("myMachine"),
		// 			Type: to.Ptr("Microsoft.HybridCompute/machines"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/machines/myMachine"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			Identity: &armhybridcompute.Identity{
		// 				Type: to.Ptr("SystemAssigned"),
		// 				PrincipalID: to.Ptr("string"),
		// 				TenantID: to.Ptr("string"),
		// 			},
		// 			Properties: &armhybridcompute.MachineProperties{
		// 				AgentConfiguration: &armhybridcompute.AgentConfiguration{
		// 					IncomingConnectionsPorts: []*string{
		// 						to.Ptr("22"),
		// 						to.Ptr("23")},
		// 						ProxyURL: to.Ptr("https://test.test"),
		// 					},
		// 					ClientPublicKey: to.Ptr("string"),
		// 					DetectedProperties: map[string]*string{
		// 						"cloudprovider": to.Ptr("N/A"),
		// 						"manufacturer": to.Ptr("Microsoft Corporation"),
		// 						"model": to.Ptr("Virtual Machine"),
		// 					},
		// 					LocationData: &armhybridcompute.LocationData{
		// 						Name: to.Ptr("Redmond"),
		// 					},
		// 					MssqlDiscovered: to.Ptr("false"),
		// 					OSProfile: &armhybridcompute.OSProfile{
		// 						LinuxConfiguration: &armhybridcompute.OSProfileLinuxConfiguration{
		// 							PatchSettings: &armhybridcompute.PatchSettings{
		// 							},
		// 						},
		// 						WindowsConfiguration: &armhybridcompute.OSProfileWindowsConfiguration{
		// 							PatchSettings: &armhybridcompute.PatchSettings{
		// 							},
		// 						},
		// 					},
		// 					PrivateLinkScopeResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/privateLinkScopes/privateLinkScopeName"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					VMID: to.Ptr("b7a098cc-b0b8-46e8-a205-62f301a62a8f"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("myMachine2"),
		// 				Type: to.Ptr("Microsoft.HybridCompute/machines"),
		// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup2/providers/Microsoft.HybridCompute/machines/myMachine2"),
		// 				Location: to.Ptr("westus2"),
		// 				Identity: &armhybridcompute.Identity{
		// 					Type: to.Ptr("SystemAssigned"),
		// 					PrincipalID: to.Ptr("e7a068cc-b0b8-46e8-a203-22f301a62a8f"),
		// 					TenantID: to.Ptr("c4098cc-91b8-46c2-a205-d82ab1a62a8f"),
		// 				},
		// 				Properties: &armhybridcompute.MachineProperties{
		// 					AgentConfiguration: &armhybridcompute.AgentConfiguration{
		// 						IncomingConnectionsPorts: []*string{
		// 							to.Ptr("22"),
		// 							to.Ptr("23")},
		// 							ProxyURL: to.Ptr("https://test.test"),
		// 						},
		// 						ClientPublicKey: to.Ptr("string"),
		// 						DetectedProperties: map[string]*string{
		// 							"cloudprovider": to.Ptr("N/A"),
		// 							"manufacturer": to.Ptr("Microsoft Corporation"),
		// 							"model": to.Ptr("Surfacebook"),
		// 						},
		// 						LocationData: &armhybridcompute.LocationData{
		// 							Name: to.Ptr("Redmond"),
		// 						},
		// 						MssqlDiscovered: to.Ptr("true"),
		// 						OSProfile: &armhybridcompute.OSProfile{
		// 							LinuxConfiguration: &armhybridcompute.OSProfileLinuxConfiguration{
		// 								PatchSettings: &armhybridcompute.PatchSettings{
		// 								},
		// 							},
		// 							WindowsConfiguration: &armhybridcompute.OSProfileWindowsConfiguration{
		// 								PatchSettings: &armhybridcompute.PatchSettings{
		// 								},
		// 							},
		// 						},
		// 						ParentClusterResourceID: to.Ptr("{AzureStackHCIResourceId}"),
		// 						ProvisioningState: to.Ptr("Succeeded"),
		// 						VMID: to.Ptr("a4a098cc-b0b8-46e8-a205-62f301a62a8f"),
		// 					},
		// 			}},
		// 		}
	}
}
