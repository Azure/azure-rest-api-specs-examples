package armsqlvirtualmachine_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine"
)

// Generated from example definition: 2023-10-01/CreateOrUpdateSqlVirtualMachineVmIdentitySettings.json
func ExampleSQLVirtualMachinesClient_BeginCreateOrUpdate_createsOrUpdatesASqlVirtualMachineToEnableTheUsageOfVirtualMachineManagedIdentity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsqlvirtualmachine.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSQLVirtualMachinesClient().BeginCreateOrUpdate(ctx, "testrg", "testvm", armsqlvirtualmachine.SQLVirtualMachine{
		Location: to.Ptr("northeurope"),
		Properties: &armsqlvirtualmachine.Properties{
			VirtualMachineIdentitySettings: &armsqlvirtualmachine.VirtualMachineIdentity{
				Type:       to.Ptr(armsqlvirtualmachine.VMIdentityTypeUserAssigned),
				ResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourcegroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testvmidentity"),
			},
			VirtualMachineResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachines/testvm"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsqlvirtualmachine.SQLVirtualMachinesClientCreateOrUpdateResponse{
	// 	SQLVirtualMachine: armsqlvirtualmachine.SQLVirtualMachine{
	// 		Name: to.Ptr("testvm"),
	// 		Type: to.Ptr("Microsoft.SqlVirtualMachine/sqlVirtualMachines"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm"),
	// 		Location: to.Ptr("northeurope"),
	// 		Properties: &armsqlvirtualmachine.Properties{
	// 			AdditionalVMPatch: to.Ptr(armsqlvirtualmachine.AdditionalOsPatchWU),
	// 			EnableAutomaticUpgrade: to.Ptr(false),
	// 			LeastPrivilegeMode: to.Ptr(armsqlvirtualmachine.LeastPrivilegeModeEnabled),
	// 			OSType: to.Ptr(armsqlvirtualmachine.OsTypeWindows),
	// 			ProvisioningState: to.Ptr("Updating"),
	// 			SQLImageOffer: to.Ptr("SQL2022-WS2022"),
	// 			SQLImageSKU: to.Ptr(armsqlvirtualmachine.SQLImageSKUEnterprise),
	// 			SQLManagement: to.Ptr(armsqlvirtualmachine.SQLManagementModeFull),
	// 			SQLServerLicenseType: to.Ptr(armsqlvirtualmachine.SQLServerLicenseTypePAYG),
	// 			VirtualMachineIdentitySettings: &armsqlvirtualmachine.VirtualMachineIdentity{
	// 				Type: to.Ptr(armsqlvirtualmachine.VMIdentityTypeUserAssigned),
	// 				ResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourcegroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testvmidentity"),
	// 			},
	// 			VirtualMachineResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachines/testvm"),
	// 		},
	// 	},
	// }
}
