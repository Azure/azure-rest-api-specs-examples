package armsqlvirtualmachine_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e79a5aa63c0551c1b5af1d2853cceb495283d334/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-08-01-preview/examples/CreateOrUpdateSqlVirtualMachineStorageConfigurationNEW.json
func ExampleSQLVirtualMachinesClient_BeginCreateOrUpdate_createsOrUpdatesASqlVirtualMachineForStorageConfigurationSettingsToNewDataLogAndTempDbStoragePool() {
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
				DiskConfigurationType: to.Ptr(armsqlvirtualmachine.DiskConfigurationTypeNEW),
				SQLDataSettings: &armsqlvirtualmachine.SQLStorageSettings{
					DefaultFilePath: to.Ptr("F:\\folderpath\\"),
					Luns: []*int32{
						to.Ptr[int32](0)},
				},
				SQLLogSettings: &armsqlvirtualmachine.SQLStorageSettings{
					DefaultFilePath: to.Ptr("G:\\folderpath\\"),
					Luns: []*int32{
						to.Ptr[int32](1)},
				},
				SQLSystemDbOnDataDisk: to.Ptr(true),
				SQLTempDbSettings: &armsqlvirtualmachine.SQLTempDbSettings{
					DataFileCount:   to.Ptr[int32](8),
					DataFileSize:    to.Ptr[int32](256),
					DataGrowth:      to.Ptr[int32](512),
					DefaultFilePath: to.Ptr("D:\\TEMP"),
					LogFileSize:     to.Ptr[int32](256),
					LogGrowth:       to.Ptr[int32](512),
				},
				StorageWorkloadType: to.Ptr(armsqlvirtualmachine.StorageWorkloadTypeOLTP),
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
	// 		LeastPrivilegeMode: to.Ptr(armsqlvirtualmachine.LeastPrivilegeModeEnabled),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		SQLImageSKU: to.Ptr(armsqlvirtualmachine.SQLImageSKUEnterprise),
	// 		SQLServerLicenseType: to.Ptr(armsqlvirtualmachine.SQLServerLicenseTypePAYG),
	// 		VirtualMachineResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachines/testvm"),
	// 	},
	// }
}
