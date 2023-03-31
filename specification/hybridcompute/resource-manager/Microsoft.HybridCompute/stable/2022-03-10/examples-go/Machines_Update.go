package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/stable/2022-03-10/examples/Machines_Update.json
func ExampleMachinesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMachinesClient().Update(ctx, "myResourceGroup", "myMachine", armhybridcompute.MachineUpdate{
		Identity: &armhybridcompute.Identity{
			Type: to.Ptr("SystemAssigned"),
		},
		Properties: &armhybridcompute.MachineUpdateProperties{
			LocationData: &armhybridcompute.LocationData{
				Name: to.Ptr("Redmond"),
			},
			OSProfile: &armhybridcompute.OSProfile{
				LinuxConfiguration: &armhybridcompute.OSProfileLinuxConfiguration{
					PatchSettings: &armhybridcompute.PatchSettings{
						AssessmentMode: to.Ptr(armhybridcompute.AssessmentModeTypesImageDefault),
						PatchMode:      to.Ptr(armhybridcompute.PatchModeTypesManual),
					},
				},
				WindowsConfiguration: &armhybridcompute.OSProfileWindowsConfiguration{
					PatchSettings: &armhybridcompute.PatchSettings{
						AssessmentMode: to.Ptr(armhybridcompute.AssessmentModeTypesImageDefault),
						PatchMode:      to.Ptr(armhybridcompute.PatchModeTypesManual),
					},
				},
			},
			ParentClusterResourceID:    to.Ptr("{AzureStackHCIResourceId}"),
			PrivateLinkScopeResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/privateLinkScopes/privateLinkScopeName"),
		},
	}, nil)
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
	// 		ClientPublicKey: to.Ptr("string"),
	// 		LocationData: &armhybridcompute.LocationData{
	// 			Name: to.Ptr("Redmond"),
	// 		},
	// 		OSProfile: &armhybridcompute.OSProfile{
	// 			LinuxConfiguration: &armhybridcompute.OSProfileLinuxConfiguration{
	// 				PatchSettings: &armhybridcompute.PatchSettings{
	// 					AssessmentMode: to.Ptr(armhybridcompute.AssessmentModeTypesImageDefault),
	// 					PatchMode: to.Ptr(armhybridcompute.PatchModeTypesManual),
	// 				},
	// 			},
	// 			WindowsConfiguration: &armhybridcompute.OSProfileWindowsConfiguration{
	// 				PatchSettings: &armhybridcompute.PatchSettings{
	// 					AssessmentMode: to.Ptr(armhybridcompute.AssessmentModeTypesImageDefault),
	// 					PatchMode: to.Ptr(armhybridcompute.PatchModeTypesManual),
	// 				},
	// 			},
	// 		},
	// 		ParentClusterResourceID: to.Ptr("{AzureStackHCIResourceId}"),
	// 		PrivateLinkScopeResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/privateLinkScopes/privateLinkScopeName"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		VMID: to.Ptr("b7a098cc-b0b8-46e8-a205-62f301a62a8f"),
	// 	},
	// }
}
