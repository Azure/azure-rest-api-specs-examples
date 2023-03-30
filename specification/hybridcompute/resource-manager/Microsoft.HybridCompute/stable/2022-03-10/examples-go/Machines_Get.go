package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/stable/2022-03-10/examples/Machines_Get.json
func ExampleMachinesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMachinesClient().Get(ctx, "myResourceGroup", "myMachine", &armhybridcompute.MachinesClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Machine = armhybridcompute.Machine{
	// 	Name: to.Ptr("myMachine"),
	// 	Type: to.Ptr("Microsoft.HybridCompute/machines"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/machines/myMachine"),
	// 	Location: to.Ptr("eastus2euap"),
	// 	Identity: &armhybridcompute.Identity{
	// 		Type: to.Ptr("SystemAssigned"),
	// 		PrincipalID: to.Ptr("string"),
	// 		TenantID: to.Ptr("string"),
	// 	},
	// 	Properties: &armhybridcompute.MachineProperties{
	// 		AgentConfiguration: &armhybridcompute.AgentConfiguration{
	// 			ExtensionsEnabled: to.Ptr("true"),
	// 			GuestConfigurationEnabled: to.Ptr("true"),
	// 			IncomingConnectionsPorts: []*string{
	// 				to.Ptr("22"),
	// 				to.Ptr("23")},
	// 				ProxyBypass: []*string{
	// 					to.Ptr("proxy1"),
	// 					to.Ptr("proxy2")},
	// 					ProxyURL: to.Ptr("https://test.test"),
	// 				},
	// 				ClientPublicKey: to.Ptr("string"),
	// 				DetectedProperties: map[string]*string{
	// 					"cloudprovider": to.Ptr("N/A"),
	// 					"manufacturer": to.Ptr("Microsoft Corporation"),
	// 					"model": to.Ptr("Virtual Machine"),
	// 				},
	// 				LocationData: &armhybridcompute.LocationData{
	// 					Name: to.Ptr("Redmond"),
	// 					City: to.Ptr("redmond"),
	// 					CountryOrRegion: to.Ptr("usa"),
	// 				},
	// 				MssqlDiscovered: to.Ptr("false"),
	// 				OSProfile: &armhybridcompute.OSProfile{
	// 					LinuxConfiguration: &armhybridcompute.OSProfileLinuxConfiguration{
	// 						PatchSettings: &armhybridcompute.PatchSettings{
	// 						},
	// 					},
	// 					WindowsConfiguration: &armhybridcompute.OSProfileWindowsConfiguration{
	// 						PatchSettings: &armhybridcompute.PatchSettings{
	// 						},
	// 					},
	// 				},
	// 				ParentClusterResourceID: to.Ptr("{AzureStackHCIResourceId}"),
	// 				PrivateLinkScopeResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/privateLinkScopes/privateLinkScopeName"),
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 				ServiceStatuses: &armhybridcompute.ServiceStatuses{
	// 					ExtensionService: &armhybridcompute.ServiceStatus{
	// 						StartupType: to.Ptr("Automatic"),
	// 						Status: to.Ptr("Running"),
	// 					},
	// 					GuestConfigurationService: &armhybridcompute.ServiceStatus{
	// 						StartupType: to.Ptr("Automatic"),
	// 						Status: to.Ptr("Running"),
	// 					},
	// 				},
	// 				VMID: to.Ptr("b7a098cc-b0b8-46e8-a205-62f301a62a8f"),
	// 			},
	// 		}
}
