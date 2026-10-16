package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-10-01/EdgeMachineDeploymentSettings_ListByEdgeMachine.json
func ExampleEdgeMachineDeploymentSettingsClient_NewListByEdgeMachinePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEdgeMachineDeploymentSettingsClient().NewListByEdgeMachinePager("test-rg", "EdgeMachine01", nil)
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
		// page = armazurestackhci.EdgeMachineDeploymentSettingsClientListByEdgeMachineResponse{
		// 	EdgeMachineDeploymentSettingListResult: armazurestackhci.EdgeMachineDeploymentSettingListResult{
		// 		Value: []*armazurestackhci.EdgeMachineDeploymentSetting{
		// 			{
		// 				Properties: &armazurestackhci.EdgeMachineDeploymentSettingProperties{
		// 					DeploymentConfiguration: &armazurestackhci.EdgeMachineDeploymentConfiguration{
		// 						Version: to.Ptr("10.2.0.0"),
		// 						ScaleUnits: []*armazurestackhci.EdgeMachineScaleUnits{
		// 							{
		// 								DeploymentData: &armazurestackhci.EdgeMachineDeploymentData{
		// 									StorageType: to.Ptr(armazurestackhci.NodeStorageTypeSanIscsi),
		// 									IScsiConfig: &armazurestackhci.IscsiConfig{
		// 										IScsiNetworkConfig: &armazurestackhci.IscsiNetworkConfig{
		// 											AdapterProperties: &armazurestackhci.IscsiAdapterProperties{
		// 												JumboPacket: to.Ptr[int32](9014),
		// 											},
		// 											AdapterIPConfig: []*armazurestackhci.IscsiAdapterIPConfig{
		// 												{
		// 													NetworkAdapterName: to.Ptr("iSCSI-A"),
		// 													VlanID: to.Ptr[int32](401),
		// 													AddressPrefix: to.Ptr("10.10.50.0/24"),
		// 													TargetPortalAddresses: []*armazurestackhci.IscsiTargetPortalAddress{
		// 														{
		// 															IPAddress: to.Ptr("10.0.0.1"),
		// 															NextHop: to.Ptr("10.10.50.1"),
		// 														},
		// 													},
		// 												},
		// 											},
		// 										},
		// 										ChapSettings: &armazurestackhci.IscsiChapSettings{
		// 											ChapType: to.Ptr(armazurestackhci.IscsiChapTypeBidirectional),
		// 										},
		// 										InitiatorIqn: to.Ptr("iqn.2025-01.com.example:storage-target"),
		// 									},
		// 								},
		// 							},
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateSucceeded),
		// 				},
		// 				ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/test-rg/providers/Microsoft.AzureStackHCI/edgeMachines/EdgeMachine01/deploymentSettings/default"),
		// 				Name: to.Ptr("default"),
		// 				Type: to.Ptr("Microsoft.AzureStackHCI/edgeMachines/deploymentSettings"),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
