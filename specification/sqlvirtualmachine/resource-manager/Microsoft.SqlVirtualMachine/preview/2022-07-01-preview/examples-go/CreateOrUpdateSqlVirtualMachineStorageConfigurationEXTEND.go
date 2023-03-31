package armsqlvirtualmachine_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-07-01-preview/examples/CreateOrUpdateSqlVirtualMachineStorageConfigurationEXTEND.json
func ExampleSQLVirtualMachinesClient_BeginCreateOrUpdate_createsOrUpdatesASqlVirtualMachineForStorageConfigurationSettingsToExtendDataLogOrTempDbStoragePool() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsqlvirtualmachine.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSQLVirtualMachinesClient().BeginCreateOrUpdate(ctx, "testrg", "testvm", armsqlvirtualmachine.SQLVirtualMachine{
		Location: to.Ptr("northeurope"),
		Properties: &armsqlvirtualmachine.Properties{
			StorageConfigurationSettings: &armsqlvirtualmachine.StorageConfigurationSettings{
				DiskConfigurationType: to.Ptr(armsqlvirtualmachine.DiskConfigurationTypeEXTEND),
				SQLDataSettings: &armsqlvirtualmachine.SQLStorageSettings{
					Luns: []*int32{
						to.Ptr[int32](2)},
				},
			},
			VirtualMachineResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachines/testvm"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SQLVirtualMachine = armsqlvirtualmachine.SQLVirtualMachine{
	// 	Name: to.Ptr("testvm"),
	// 	Type: to.Ptr("Microsoft.SqlVirtualMachine/sqlVirtualMachines"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm"),
	// 	Location: to.Ptr("northeurope"),
	// 	Properties: &armsqlvirtualmachine.Properties{
	// 		EnableAutomaticUpgrade: to.Ptr(false),
	// 		LeastPrivilegeMode: to.Ptr(armsqlvirtualmachine.LeastPrivilegeMode("NotSet")),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		SQLImageSKU: to.Ptr(armsqlvirtualmachine.SQLImageSKUEnterprise),
	// 		SQLServerLicenseType: to.Ptr(armsqlvirtualmachine.SQLServerLicenseTypePAYG),
	// 		VirtualMachineResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachines/testvm"),
	// 	},
	// }
}
