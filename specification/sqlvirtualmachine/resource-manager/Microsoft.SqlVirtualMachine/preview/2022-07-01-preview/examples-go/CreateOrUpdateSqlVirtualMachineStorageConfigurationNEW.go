package armsqlvirtualmachine_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-07-01-preview/examples/CreateOrUpdateSqlVirtualMachineStorageConfigurationNEW.json
func ExampleSQLVirtualMachinesClient_BeginCreateOrUpdate_createsOrUpdatesASqlVirtualMachineForStorageConfigurationSettingsToNewDataLogAndTempDbStoragePool() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsqlvirtualmachine.NewSQLVirtualMachinesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "testrg", "testvm", armsqlvirtualmachine.SQLVirtualMachine{
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
	// TODO: use response item
	_ = res
}
